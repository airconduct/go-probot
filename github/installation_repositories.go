package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "installation_repositories"

type InstallationRepositoriesContext = probot.ProbotContext[probot.GithubClient, github.InstallationRepositoriesEvent]

func InstallationRepositoriesHandler(fn func(ctx InstallationRepositoriesContext)) probot.EventHandlerFunc[probot.GithubClient, github.InstallationRepositoriesEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.InstallationRepositoriesEvent]) {
		fn(ctx)
	}
}

func installationRepositories() probot.WebhookEvent {
	return probot.MakeWebhookEvent("installation_repositories", "")
}

func installationRepositories_added() probot.WebhookEvent {
	return probot.MakeWebhookEvent("installation_repositories", "added")
}

func installationRepositories_removed() probot.WebhookEvent {
	return probot.MakeWebhookEvent("installation_repositories", "removed")
}
