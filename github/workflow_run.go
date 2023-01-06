package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "workflow_run"

type WorkflowRunContext = probot.ProbotContext[probot.GithubClient, github.WorkflowRunEvent]

func WorkflowRunHandler(fn func(ctx WorkflowRunContext)) probot.EventHandlerFunc[probot.GithubClient, github.WorkflowRunEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.WorkflowRunEvent]) {
		fn(ctx)
	}
}

func workflowRun() probot.WebhookEvent {
	return probot.MakeWebhookEvent("workflow_run", "")
}

func workflowRun_completed() probot.WebhookEvent {
	return probot.MakeWebhookEvent("workflow_run", "completed")
}

func workflowRun_in_progress() probot.WebhookEvent {
	return probot.MakeWebhookEvent("workflow_run", "in_progress")
}

func workflowRun_requested() probot.WebhookEvent {
	return probot.MakeWebhookEvent("workflow_run", "requested")
}
