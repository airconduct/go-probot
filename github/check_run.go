package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

type CheckRunContext = probot.ProbotContext[probot.GithubClient, github.CheckRunEvent]

func CheckRunHandler(fn func(ctx CheckRunContext)) probot.EventHandlerFunc[probot.GithubClient, github.CheckRunEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.CheckRunEvent]) {
		fn(ctx)
	}
}

func checkRun() probot.WebhookEvent {
	return probot.MakeWebhookEvent("check_run", "")
}

func checkRun_completed() probot.WebhookEvent {
	return probot.MakeWebhookEvent("check_run", "completed")
}

func checkRun_created() probot.WebhookEvent {
	return probot.MakeWebhookEvent("check_run", "created")
}

func checkRun_requested_action() probot.WebhookEvent {
	return probot.MakeWebhookEvent("check_run", "requested_action")
}

func checkRun_rerequested() probot.WebhookEvent {
	return probot.MakeWebhookEvent("check_run", "rerequested")
}
