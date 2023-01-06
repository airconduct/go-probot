package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "deployment"

type DeploymentContext = probot.ProbotContext[probot.GithubClient, github.DeploymentEvent]

func DeploymentHandler(fn func(ctx DeploymentContext)) probot.EventHandlerFunc[probot.GithubClient, github.DeploymentEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.DeploymentEvent]) {
		fn(ctx)
	}
}

func deployment() probot.WebhookEvent {
	return probot.MakeWebhookEvent("deployment", "")
}

func deployment_created() probot.WebhookEvent {
	return probot.MakeWebhookEvent("deployment", "created")
}
