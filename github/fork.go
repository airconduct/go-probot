package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "fork"

type ForkContext = probot.ProbotContext[probot.GithubClient, github.ForkEvent]

func ForkHandler(fn func(ctx ForkContext)) probot.EventHandlerFunc[probot.GithubClient, github.ForkEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.ForkEvent]) {
		fn(ctx)
	}
}

func fork() probot.WebhookEvent {
	return probot.MakeWebhookEvent("fork", "")
}
