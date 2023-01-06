package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "pull_request"

type PullRequestContext = probot.ProbotContext[probot.GithubClient, github.PullRequestEvent]

func PullRequestHandler(fn func(ctx PullRequestContext)) probot.EventHandlerFunc[probot.GithubClient, github.PullRequestEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.PullRequestEvent]) {
		fn(ctx)
	}
}

func pullRequest() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request", "")
}

func pullRequest_assigned() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request", "assigned")
}

func pullRequest_auto_merge_disabled() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request", "auto_merge_disabled")
}

func pullRequest_auto_merge_enabled() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request", "auto_merge_enabled")
}

func pullRequest_closed() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request", "closed")
}

func pullRequest_converted_to_draft() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request", "converted_to_draft")
}

func pullRequest_demilestoned() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request", "demilestoned")
}

func pullRequest_dequeued() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request", "dequeued")
}

func pullRequest_edited() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request", "edited")
}

func pullRequest_labeled() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request", "labeled")
}

func pullRequest_locked() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request", "locked")
}

func pullRequest_milestoned() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request", "milestoned")
}

func pullRequest_opened() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request", "opened")
}

func pullRequest_ready_for_review() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request", "ready_for_review")
}

func pullRequest_reopened() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request", "reopened")
}

func pullRequest_review_request_removed() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request", "review_request_removed")
}

func pullRequest_review_requested() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request", "review_requested")
}

func pullRequest_synchronize() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request", "synchronize")
}

func pullRequest_unassigned() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request", "unassigned")
}

func pullRequest_unlabeled() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request", "unlabeled")
}

func pullRequest_unlocked() probot.WebhookEvent {
	return probot.MakeWebhookEvent("pull_request", "unlocked")
}
