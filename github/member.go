package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "member"

type MemberContext = probot.ProbotContext[probot.GithubClient, github.MemberEvent]

func MemberHandler(fn func(ctx MemberContext)) probot.EventHandlerFunc[probot.GithubClient, github.MemberEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.MemberEvent]) {
		fn(ctx)
	}
}

func member() probot.WebhookEvent {
	return probot.MakeWebhookEvent("member", "")
}

func member_added() probot.WebhookEvent {
	return probot.MakeWebhookEvent("member", "added")
}

func member_edited() probot.WebhookEvent {
	return probot.MakeWebhookEvent("member", "edited")
}

func member_removed() probot.WebhookEvent {
	return probot.MakeWebhookEvent("member", "removed")
}
