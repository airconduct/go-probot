package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

type CheckSuiteContext = probot.ProbotContext[probot.GithubClient, github.CheckSuiteEvent]

func CheckSuiteHandler(fn func(ctx CheckSuiteContext)) probot.EventHandlerFunc[probot.GithubClient, github.CheckSuiteEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.CheckSuiteEvent]) {
		fn(ctx)
	}
}

func checkSuite() probot.WebhookEvent {
	return probot.MakeWebhookEvent("check_suite", "")
}

func checkSuite_completed() probot.WebhookEvent {
	return probot.MakeWebhookEvent("check_suite", "completed")
}

func checkSuite_requested() probot.WebhookEvent {
	return probot.MakeWebhookEvent("check_suite", "requested")
}

func checkSuite_rerequested() probot.WebhookEvent {
	return probot.MakeWebhookEvent("check_suite", "rerequested")
}
