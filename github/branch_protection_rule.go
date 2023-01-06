package github

import (
	"github.com/google/go-github/v48/github"

	"github.com/airconduct/go-probot"
)

type BranchProtectionRuleContext = probot.ProbotContext[probot.GithubClient, github.BranchProtectionRuleEvent]

func BranchProtectionRuleHandler(fn func(ctx BranchProtectionRuleContext)) probot.EventHandlerFunc[probot.GithubClient, github.BranchProtectionRuleEvent] {
	return func(ctx probot.ProbotContext[probot.GithubClient, github.BranchProtectionRuleEvent]) {
		fn(ctx)
	}
}

func branchProtectionRule() probot.WebhookEvent {
	return probot.MakeWebhookEvent("branch_protection_rule", "")
}

func branchProtectionRule_deleted() probot.WebhookEvent {
	return probot.MakeWebhookEvent("branch_protection_rule", "deleted")
}

func branchProtectionRule_created() probot.WebhookEvent {
	return probot.MakeWebhookEvent("branch_protection_rule", "created")
}

func branchProtectionRule_edited() probot.WebhookEvent {
	return probot.MakeWebhookEvent("branch_protection_rule", "edited")
}
