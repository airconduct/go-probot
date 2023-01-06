package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "status"

type StatusContext = probot.ProbotContext[probot.GithubClient, github.StatusEvent]

func StatusHandler(fn func(ctx StatusContext)) probot.EventHandlerFunc[probot.GithubClient, github.StatusEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.StatusEvent]) {
		fn(ctx)
	}
}

func status() probot.WebhookEvent {
	return probot.MakeWebhookEvent("status", "")
}
