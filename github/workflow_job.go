package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "workflow_job"

type WorkflowJobContext = probot.ProbotContext[probot.GithubClient, github.WorkflowJobEvent]

func WorkflowJobHandler(fn func(ctx WorkflowJobContext)) probot.EventHandlerFunc[probot.GithubClient, github.WorkflowJobEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.WorkflowJobEvent]) {
		fn(ctx)
	}
}

func workflowJob() probot.WebhookEvent {
	return probot.MakeWebhookEvent("workflow_job", "")
}

func workflowJob_completed() probot.WebhookEvent {
	return probot.MakeWebhookEvent("workflow_job", "completed")
}

func workflowJob_in_progress() probot.WebhookEvent {
	return probot.MakeWebhookEvent("workflow_job", "in_progress")
}

func workflowJob_queued() probot.WebhookEvent {
	return probot.MakeWebhookEvent("workflow_job", "queued")
}

func workflowJob_waiting() probot.WebhookEvent {
	return probot.MakeWebhookEvent("workflow_job", "waiting")
}
