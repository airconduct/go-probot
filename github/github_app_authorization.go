package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "github_app_authorization"

type GithubAppAuthorizationContext = probot.ProbotContext[probot.GithubClient, github.GitHubAppAuthorizationEvent]

func GithubAppAuthorizationHandler(fn func(ctx GithubAppAuthorizationContext)) probot.EventHandlerFunc[probot.GithubClient, github.GitHubAppAuthorizationEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.GitHubAppAuthorizationEvent]) {
		fn(ctx)
	}
}

func githubAppAuthorization() probot.WebhookEvent {
	return probot.MakeWebhookEvent("github_app_authorization", "")
}
