package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "push"

type PushContext = probot.ProbotContext[probot.GithubClient, github.PushEvent]

func PushHandler(fn func(ctx PushContext)) probot.EventHandlerFunc[probot.GithubClient, github.PushEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.PushEvent]) {
		fn(ctx)
	}
}

func push() probot.WebhookEvent {
	return probot.MakeWebhookEvent("push", "")
}
