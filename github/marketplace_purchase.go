package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

// "marketplace_purchase"

type MarketplacePurchaseContext = probot.ProbotContext[probot.GithubClient, github.MarketplacePurchaseEvent]

func MarketplacePurchaseHandler(fn func(ctx MarketplacePurchaseContext)) probot.EventHandlerFunc[probot.GithubClient, github.MarketplacePurchaseEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.MarketplacePurchaseEvent]) {
		fn(ctx)
	}
}

func marketplacePurchase() probot.WebhookEvent {
	return probot.MakeWebhookEvent("marketplace_purchase", "")
}

func marketplacePurchase_cancelled() probot.WebhookEvent {
	return probot.MakeWebhookEvent("marketplace_purchase", "cancelled")
}

func marketplacePurchase_changed() probot.WebhookEvent {
	return probot.MakeWebhookEvent("marketplace_purchase", "changed")
}

func marketplacePurchase_pending_change() probot.WebhookEvent {
	return probot.MakeWebhookEvent("marketplace_purchase", "pending_change")
}

func marketplacePurchase_pending_change_cancelled() probot.WebhookEvent {
	return probot.MakeWebhookEvent("marketplace_purchase", "pending_change_cancelled")
}

func marketplacePurchase_purchased() probot.WebhookEvent {
	return probot.MakeWebhookEvent("marketplace_purchase", "purchased")
}
