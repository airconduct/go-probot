package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "membership"

type MembershipContext = probot.ProbotContext[probot.GithubClient, github.MembershipEvent]

func MembershipHandler(fn func(ctx MembershipContext)) probot.EventHandlerFunc[probot.GithubClient, github.MembershipEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.MembershipEvent]) {
		fn(ctx)
	}
}

func membership() probot.WebhookEvent {
	return probot.MakeWebhookEvent("membership", "")
}

func membership_added() probot.WebhookEvent {
	return probot.MakeWebhookEvent("membership", "added")
}

func membership_removed() probot.WebhookEvent {
	return probot.MakeWebhookEvent("membership", "removed")
}
