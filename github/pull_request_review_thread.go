package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "pull_request_review_thread"

type PullRequestReviewThreadContext = probot.ProbotContext[probot.GithubClient, github.PullRequestReviewThreadEvent]

func PullRequestReviewThreadHandler(fn func(ctx PullRequestReviewThreadContext)) probot.EventHandlerFunc[probot.GithubClient, github.PullRequestReviewThreadEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.PullRequestReviewThreadEvent]) {
		fn(ctx)
	}
}

func pullRequestReviewThread() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request_review_thread", "")
}

func pullRequestReviewThread_resolved() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request_review_thread", "resolved")
}

func pullRequestReviewThread_unresolved() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request_review_thread", "unresolved")
}
