package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "deploy_key"

type DeployKeyContext = probot.ProbotContext[probot.GithubClient, github.DeployKeyEvent]

func DeployKeyHandler(fn func(ctx DeployKeyContext)) probot.EventHandlerFunc[probot.GithubClient, github.DeployKeyEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.DeployKeyEvent]) {
		fn(ctx)
	}
}

func deployKey() probot.WebhookEvent {
	return probot.MakeWebhookEvent("deploy_key", "")
}

func deployKeyCreated() probot.WebhookEvent {
	return probot.MakeWebhookEvent("deploy_key", "created")
}

func deployKey_deleted() probot.WebhookEvent {
	return probot.MakeWebhookEvent("deploy_key", "deleted")
}
