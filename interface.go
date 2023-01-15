package probot

import (
	"context"
	"net/http"

	"github.com/go-logr/logr"
	"github.com/spf13/pflag"
)

type App[GT GitClientType] interface {
	AddFlags(flags *pflag.FlagSet)
	On(events ...WebhookEvent) handlerLoader
	Run(ctx context.Context) error

	ServeMux() *http.ServeMux
}

type ProbotContext[GT GitClientType, PT gitEventType] interface {
	context.Context

	Payload() *PT
	Client() *GT
	GraphQL() GitGraphQLClient
	Logger() logr.Logger
	Must(...interface{})
}

type GitGraphQLClient interface {
	Query(ctx context.Context, q interface{}, variables map[string]interface{}) error
}
