package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "pull_request_review_comment"

type PullRequestReviewCommentContext = probot.ProbotContext[probot.GithubClient, github.PullRequestReviewCommentEvent]

func PullRequestReviewCommentHandler(fn func(ctx PullRequestReviewCommentContext)) probot.EventHandlerFunc[probot.GithubClient, github.PullRequestReviewCommentEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.PullRequestReviewCommentEvent]) {
		fn(ctx)
	}
}

func pullRequestReviewComment() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request_review_comment", "")
}

func pullRequestReviewComment_created() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request_review_comment", "created")
}

func pullRequestReviewComment_deleted() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request_review_comment", "deleted")
}

func pullRequestReviewComment_edited() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request_review_comment", "edited")
}
