package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "deployment_status"

type DeploymentStatusContext = probot.ProbotContext[probot.GithubClient, github.DeploymentStatusEvent]

func DeploymentStatusHandler(fn func(ctx DeploymentStatusContext)) probot.EventHandlerFunc[probot.GithubClient, github.DeploymentStatusEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.DeploymentStatusEvent]) {
		fn(ctx)
	}
}

func deploymentStatus() probot.WebhookEvent {
	return probot.MakeWebhookEvent("deployment_status", "")
}

func deploymentStatus_created() probot.WebhookEvent {
	return probot.MakeWebhookEvent("deployment_status", "created")
}
