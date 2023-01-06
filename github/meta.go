package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "meta"

type MetaContext = probot.ProbotContext[probot.GithubClient, github.MetaEvent]

func MetaHandler(fn func(ctx MetaContext)) probot.EventHandlerFunc[probot.GithubClient, github.MetaEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.MetaEvent]) {
		fn(ctx)
	}
}

func meta() probot.WebhookEvent {
	return probot.MakeWebhookEvent("meta", "")
}

func metaDeleted() probot.WebhookEvent {
	return probot.MakeWebhookEvent("meta", "deleted")
}
