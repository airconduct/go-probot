package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "discussion"

type DiscussionContext = probot.ProbotContext[probot.GithubClient, github.DiscussionEvent]

func DiscussionHandler(fn func(ctx DiscussionContext)) probot.EventHandlerFunc[probot.GithubClient, github.DiscussionEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.DiscussionEvent]) {
		fn(ctx)
	}
}

func discussion() probot.WebhookEvent {
	return probot.MakeWebhookEvent("discussion", "")
}

func discussion_answered() probot.WebhookEvent {
	return probot.MakeWebhookEvent("discussion", "answered")
}

func discussion_category_changed() probot.WebhookEvent {
	return probot.MakeWebhookEvent("discussion", "category_changed")
}

func discussion_created() probot.WebhookEvent {
	return probot.MakeWebhookEvent("discussion", "created")
}

func discussion_deleted() probot.WebhookEvent {
	return probot.MakeWebhookEvent("discussion", "deleted")
}

func discussion_edited() probot.WebhookEvent {
	return probot.MakeWebhookEvent("discussion", "edited")
}

func discussion_labeled() probot.WebhookEvent {
	return probot.MakeWebhookEvent("discussion", "labeled")
}

func discussion_locked() probot.WebhookEvent {
	return probot.MakeWebhookEvent("discussion", "locked")
}

func discussion_pinned() probot.WebhookEvent {
	return probot.MakeWebhookEvent("discussion", "pinned")
}

func discussion_transferred() probot.WebhookEvent {
	return probot.MakeWebhookEvent("discussion", "transferred")
}

func discussion_unanswered() probot.WebhookEvent {
	return probot.MakeWebhookEvent("discussion", "unanswered")
}

func discussion_unlabeled() probot.WebhookEvent {
	return probot.MakeWebhookEvent("discussion", "unlabeled")
}

func discussion_unlocked() probot.WebhookEvent {
	return probot.MakeWebhookEvent("discussion", "unlocked")
}

func discussion_unpinned() probot.WebhookEvent {
	return probot.MakeWebhookEvent("discussion", "unpinned")
}
