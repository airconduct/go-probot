package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "code_scanning_alert"

type CodeScanningAlertContext = probot.ProbotContext[probot.GithubClient, github.CodeScanningAlertEvent]

func CodeScanningAlertHandler(fn func(ctx CodeScanningAlertContext)) probot.EventHandlerFunc[probot.GithubClient, github.CodeScanningAlertEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.CodeScanningAlertEvent]) {
		fn(ctx)
	}
}

func codeScanningAlert() probot.WebhookEvent {
	return probot.MakeWebhookEvent("code_scanning_alert", "")
}

func codeScanningAlert_appeared_in_branch() probot.WebhookEvent {
	return probot.MakeWebhookEvent("code_scanning_alert", "appeared_in_branch")
}

func codeScanningAlert_closed_by_user() probot.WebhookEvent {
	return probot.MakeWebhookEvent("code_scanning_alert", "closed_by_user")
}

func codeScanningAlert_created() probot.WebhookEvent {
	return probot.MakeWebhookEvent("code_scanning_alert", "created")
}

func codeScanningAlert_fixed() probot.WebhookEvent {
	return probot.MakeWebhookEvent("code_scanning_alert", "fixed")
}

func codeScanningAlert_reopened() probot.WebhookEvent {
	return probot.MakeWebhookEvent("code_scanning_alert", "reopened")
}

func codeScanningAlert_reopened_by_user() probot.WebhookEvent {
	return probot.MakeWebhookEvent("code_scanning_alert", "reopened_by_user")
}
