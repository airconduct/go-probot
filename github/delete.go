package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "delete"

type DeleteContext = probot.ProbotContext[probot.GithubClient, github.DeleteEvent]

func DeleteHandler(fn func(ctx DeleteContext)) probot.EventHandlerFunc[probot.GithubClient, github.DeleteEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.DeleteEvent]) {
		fn(ctx)
	}
}

func delete() probot.WebhookEvent {
	return probot.MakeWebhookEvent("delete", "")
}
