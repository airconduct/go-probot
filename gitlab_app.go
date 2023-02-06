package probot

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/pflag"
	"github.com/xanzy/go-gitlab"
)

func NewGitLabAPP() App[GitLabClient] {
	return &gitlabApp{
		handlers: make(map[string]Handler),
		// graphqlClients: make(map[int64]*githubv4.Client),
		serveMux: http.NewServeMux(),
		metrics:  new(eventMetrics),
	}
}

type gitlabApp struct {
	handlers map[string]Handler
	client   *gitlab.Client

	baseURL    string
	token      []byte
	tokenFile  string
	secret     []byte
	secretFile string

	serverOpts ServerOptions
	serveMux   *http.ServeMux

	metrics *eventMetrics

	loggerOptions
}

func (app *gitlabApp) AddFlags(flags *pflag.FlagSet) {
	app.serverOpts.AddFlags(flags)
	app.loggerOptions.AddFlags(flags)
	flags.StringVar(&app.tokenFile, "gitlab.token-file", "", "GitLab token file")
	flags.StringVar(&app.secretFile, "gitlab.secret-file", "", "GitLab token file")
	flags.StringVar(&app.baseURL, "gitlab.base-url", "https://gitlab.com", "GitLab base URL")
}

func (app *gitlabApp) On(events ...WebhookEvent) handlerLoader {
	return handlerLoadFunc(func(h Handler) error {
		for _, event := range events {
			key := event.Type()
			if action := event.Action(); action != "" {
				key = key + "." + action
			}
			if _, ok := app.handlers[key]; ok {
				return fmt.Errorf("event type %s already exists", key)
			}
			app.handlers[key] = h
			app.metrics.add(key, event.Type())
		}
		return nil
	})
}

func (app *gitlabApp) Run(ctx context.Context) error {
	// TODO: Implement
	if err := app.initialize(); err != nil {
		return err
	}

	if app.serverOpts.Path == "" {
		app.serverOpts.Path = "/"
	}
	app.serveMux.HandleFunc(app.serverOpts.Path, app.handle)
	server := &http.Server{Addr: fmt.Sprintf("%s:%d", app.serverOpts.Address, app.serverOpts.Port), Handler: app.serveMux}
	server.RegisterOnShutdown(app.shutdown)
	app.logger.Info("Kuilei hook is serving", "addr", server.Addr)
	return server.ListenAndServe()
}

//go:generate go run tools/codegen/main.go -t gitlab_app_handle.go.tmpl -v gitlab_events.yaml -o gitlab_app_handle.go
func (app *gitlabApp) handle(w http.ResponseWriter, r *http.Request) {
	var handlerKey string
	defer func() {
		if e := recover(); e != nil {
			err, ok := e.(error)
			if ok {
				app.handleError(w, err, http.StatusBadRequest)
				app.logger.Error(
					err, "Failed handle event",
					"handler_key", handlerKey,
				)
			}
		} else {
			app.logger.Info("Succeeded to handle event", "handler_key", handlerKey)
		}
	}()

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	if r.Method != http.MethodPost {
		app.handleError(w, fmt.Errorf("method %s not found", r.Method), http.StatusNotFound)
		return
	}

	eventType := gitlab.WebhookEventType(r)
	app.logger.Info("Handle event", "event", eventType)

	rawPayload, err := validateGitLabPayload(r, string(app.secret))
	if err != nil {
		app.handleError(w, fmt.Errorf("failed to validate payload, %w", err), http.StatusBadRequest)
		return
	}
	payload, err := gitlab.ParseHook(eventType, rawPayload)
	if err != nil {
		app.handleError(w, fmt.Errorf("failed to parse webhook, %w", err), http.StatusBadRequest)
		return
	}

	handlerKey = app.handelEvent(ctx, w, r, string(eventType), payload)
}

func (app *gitlabApp) shutdown() {}

func (app *gitlabApp) ServeMux() *http.ServeMux {
	return app.serveMux
}

func (app *gitlabApp) GetServerOptions() ServerOptions {
	return app.serverOpts
}

func (app *gitlabApp) GetHTTPHandler() http.Handler {
	return http.HandlerFunc(app.handle)
}

func (app *gitlabApp) GetSecretToken() []byte {
	// TODO: Implement
	return nil
}

func (app *gitlabApp) initialize() (err error) {
	app.token, err = os.ReadFile(app.tokenFile)
	if err != nil {
		return
	}
	app.client, err = gitlab.NewClient(string(app.token), gitlab.WithBaseURL(app.baseURL), gitlab.WithHTTPClient(http.DefaultClient))
	if err != nil {
		return
	}
	return nil
}

func (app *gitlabApp) handleError(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)
	fmt.Fprint(w, err.Error())
}

func validateGitLabPayload(r *http.Request, secret string) (payload []byte, err error) {
	signature := r.Header.Get("X-Gitlab-Token")
	if signature != secret {
		return nil, errors.New("token validation failed")
	}

	return io.ReadAll(r.Body)
}
