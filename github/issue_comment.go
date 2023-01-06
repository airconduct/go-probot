package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "issue_comment"

type IssueCommentContext = probot.ProbotContext[probot.GithubClient, github.IssueCommentEvent]

func IssueCommentHandler(fn func(ctx IssueCommentContext)) probot.EventHandlerFunc[probot.GithubClient, github.IssueCommentEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.IssueCommentEvent]) {
		fn(ctx)
	}
}

func issueComment() probot.WebhookEvent {
	return probot.MakeWebhookEvent("issue_comment", "")
}

func issueComment_created() probot.WebhookEvent {
	return probot.MakeWebhookEvent("issue_comment", "created")
}

func issueComment_deleted() probot.WebhookEvent {
	return probot.MakeWebhookEvent("issue_comment", "deleted")
}

func issueComment_edited() probot.WebhookEvent {
	return probot.MakeWebhookEvent("issue_comment", "edited")
}
