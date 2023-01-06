package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "issues"

type IssuesContext = probot.ProbotContext[probot.GithubClient, github.IssuesEvent]

func IssuesHandler(fn func(ctx IssuesContext)) probot.EventHandlerFunc[probot.GithubClient, github.IssuesEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.IssuesEvent]) {
		fn(ctx)
	}
}

func issues() probot.WebhookEvent {
	return probot.MakeWebhookEvent("issues", "")
}

func issues_assigned() probot.WebhookEvent {
	return probot.MakeWebhookEvent("issues", "assigned")
}

func issues_closed() probot.WebhookEvent {
	return probot.MakeWebhookEvent("issues", "closed")
}

func issues_deleted() probot.WebhookEvent {
	return probot.MakeWebhookEvent("issues", "deleted")
}

func issues_demilestoned() probot.WebhookEvent {
	return probot.MakeWebhookEvent("issues", "demilestoned")
}

func issues_edited() probot.WebhookEvent {
	return probot.MakeWebhookEvent("issues", "edited")
}

func issues_labeled() probot.WebhookEvent {
	return probot.MakeWebhookEvent("issues", "labeled")
}

func issues_locked() probot.WebhookEvent {
	return probot.MakeWebhookEvent("issues", "locked")
}

func issues_milestoned() probot.WebhookEvent {
	return probot.MakeWebhookEvent("issues", "milestoned")
}

func issues_opened() probot.WebhookEvent {
	return probot.MakeWebhookEvent("issues", "opened")
}

func issues_pinned() probot.WebhookEvent {
	return probot.MakeWebhookEvent("issues", "pinned")
}

func issues_reopened() probot.WebhookEvent {
	return probot.MakeWebhookEvent("issues", "reopened")
}

func issues_transferred() probot.WebhookEvent {
	return probot.MakeWebhookEvent("issues", "transferred")
}

func issues_unassigned() probot.WebhookEvent {
	return probot.MakeWebhookEvent("issues", "unassigned")
}

func issues_unlabeled() probot.WebhookEvent {
	return probot.MakeWebhookEvent("issues", "unlabeled")
}

func issues_unlocked() probot.WebhookEvent {
	return probot.MakeWebhookEvent("issues", "unlocked")
}

func issues_unpinned() probot.WebhookEvent {
	return probot.MakeWebhookEvent("issues", "unpinned")
}
