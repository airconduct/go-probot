package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "gollum"

type GollumContext = probot.ProbotContext[probot.GithubClient, github.GollumEvent]

func GollumHandler(fn func(ctx GollumContext)) probot.EventHandlerFunc[probot.GithubClient, github.GollumEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.GollumEvent]) {
		fn(ctx)
	}
}

func gollum() probot.WebhookEvent {
	return probot.MakeWebhookEvent("gollum", "")
}
