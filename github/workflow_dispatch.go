package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "workflow_dispatch"

type WorkflowDispatchContext = probot.ProbotContext[probot.GithubClient, github.WorkflowDispatchEvent]

func WorkflowDispatchHandler(fn func(ctx WorkflowDispatchContext)) probot.EventHandlerFunc[probot.GithubClient, github.WorkflowDispatchEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.WorkflowDispatchEvent]) {
		fn(ctx)
	}
}

func workflowDispatch() probot.WebhookEvent {
	return probot.MakeWebhookEvent("workflow_dispatch", "")
}
