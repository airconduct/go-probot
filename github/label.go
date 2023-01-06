package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "label"

type LabelContext = probot.ProbotContext[probot.GithubClient, github.LabelEvent]

func LabelHandler(fn func(ctx LabelContext)) probot.EventHandlerFunc[probot.GithubClient, github.LabelEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.LabelEvent]) {
		fn(ctx)
	}
}

func label() probot.WebhookEvent {
	return probot.MakeWebhookEvent("label", "")
}

func lable_created() probot.WebhookEvent {
	return probot.MakeWebhookEvent("label", "created")
}

func lable_deleted() probot.WebhookEvent {
	return probot.MakeWebhookEvent("label", "deleted")
}

func lable_edited() probot.WebhookEvent {
	return probot.MakeWebhookEvent("label", "edited")
}
