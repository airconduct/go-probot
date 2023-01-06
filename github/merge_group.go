package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "merge_group"

type MergeGroupContext = probot.ProbotContext[probot.GithubClient, github.MergeGroupEvent]

func MergeGroupHandler(fn func(ctx MergeGroupContext)) probot.EventHandlerFunc[probot.GithubClient, github.MergeGroupEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.MergeGroupEvent]) {
		fn(ctx)
	}
}

func mergeGroup() probot.WebhookEvent {
	return probot.MakeWebhookEvent("merge_group", "")
}
