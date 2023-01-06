package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "pull_request_review"

type PullRequestReviewContext = probot.ProbotContext[probot.GithubClient, github.PullRequestReviewEvent]

func PullRequestReviewHandler(fn func(ctx PullRequestReviewContext)) probot.EventHandlerFunc[probot.GithubClient, github.PullRequestReviewEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.PullRequestReviewEvent]) {
		fn(ctx)
	}
}

func pullRequestReview() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request_review", "")
}

func pullRequestReview_dismissed() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request_review", "dismissed")
}

func pullRequestReview_edited() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request_review", "edited")
}

func pullRequestReview_submitted() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request_review", "submitted")
}
