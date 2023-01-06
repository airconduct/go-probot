package main

import (
	"context"
	"time"

	gh "github.com/google/go-github/v48/github"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/pflag"

	"github.com/airconduct/go-probot"
	"github.com/airconduct/go-probot/github"
	"github.com/airconduct/go-probot/mock"
)

var _ = Describe("Test Probot Example", func() {
	It("Reply issue comments", func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		flags := pflag.NewFlagSet("", pflag.PanicOnError)
		app := probot.NewGithubAPP()
		app.AddFlags(flags)
		flags.Parse([]string{
			"--github.hmac-token-file=testdata/hmac_token",
			"--github.private-key-file=testdata/tls.key",
			"--github.appid=1",
			"--address=127.0.0.1",
			"--port=7771",
			"--path=/hook",
		})

		app.On(github.Event.IssueComment_created, github.Event.IssueComment_edited).
			WithHandler(github.IssueCommentHandler(func(ctx github.IssueCommentContext) {
				payload := ctx.Payload()
				ctx.Logger().Info("Get IssueComment event", "payload", payload)
				owner := *payload.Repo.Owner.Login
				repo := *payload.Repo.Name
				issueNumber := *payload.Issue.Number
				ctx.Must(ctx.Client().Issues.CreateComment(ctx, owner, repo, issueNumber, &gh.IssueComment{
					Body: probot.ToPointer("Thanks!"),
				}))
			}))

		go func() {
			Expect(app.Run(ctx)).Should(Succeed())
		}()

		Eventually(func(g Gomega) {
			g.Expect(mock.Send(
				app.(mock.AppMock[probot.GithubClient]),
				github.Event.IssueComment_created,
				gh.IssueCommentEvent{
					Action:       probot.ToPointer("created"),
					Installation: &gh.Installation{ID: probot.ToPointer(int64(installationID))},
					Repo: &gh.Repository{
						Name:  probot.ToPointer("bar"),
						Owner: &gh.User{Login: probot.ToPointer("foo")},
					},
					Issue: &gh.Issue{Number: probot.ToPointer(1)},
				},
			)).Should(Succeed())
		}, 5*time.Second, time.Second).Should(Succeed())
	})
})
