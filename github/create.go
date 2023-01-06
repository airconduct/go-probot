package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "create"

type CreateContext = probot.ProbotContext[probot.GithubClient, github.CreateEvent]

func CreateHandler(fn func(ctx CreateContext)) probot.EventHandlerFunc[probot.GithubClient, github.CreateEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.CreateEvent]) {
		fn(ctx)
	}
}

func create() probot.WebhookEvent {
	return probot.MakeWebhookEvent("create", "")
}
