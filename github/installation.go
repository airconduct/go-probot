package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "installation"

type InstallationContext = probot.ProbotContext[probot.GithubClient, github.InstallationEvent]

func InstallationHandler(fn func(ctx InstallationContext)) probot.EventHandlerFunc[probot.GithubClient, github.InstallationEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.InstallationEvent]) {
		fn(ctx)
	}
}

func installation() probot.WebhookEvent {
	return probot.MakeWebhookEvent("installation", "")
}

func installation_created() probot.WebhookEvent {
	return probot.MakeWebhookEvent("installation", "created")
}

func installation_deleted() probot.WebhookEvent {
	return probot.MakeWebhookEvent("installation", "deleted")
}

func installation_new_permissions_accepted() probot.WebhookEvent {
	return probot.MakeWebhookEvent("installation", "new_permissions_accepted")
}

func installation_suspend() probot.WebhookEvent {
	return probot.MakeWebhookEvent("installation", "suspend")
}

func installation_unsuspend() probot.WebhookEvent {
	return probot.MakeWebhookEvent("installation", "unsuspend")
}
