package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "commit_comment"

type CommitCommentContext = probot.ProbotContext[probot.GithubClient, github.CommitCommentEvent]

func CommitCommentHandler(fn func(ctx CommitCommentContext)) probot.EventHandlerFunc[probot.GithubClient, github.CommitCommentEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.CommitCommentEvent]) {
		fn(ctx)
	}
}

func commitComment() probot.WebhookEvent {
	return probot.MakeWebhookEvent("commit_comment", "")
}

func commitComment_created() probot.WebhookEvent {
	return probot.MakeWebhookEvent("commit_comment", "created")
}
