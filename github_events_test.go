// Code generated by codegen. DO NOT EDIT.

package probot_test

import (
	"fmt"
	"testing"

	"github.com/airconduct/go-probot"
	"github.com/google/go-github/v48/github"
)

func TestBranchProtectionRule(t *testing.T) {
	checkEvent(t, probot.GitHub.BranchProtectionRule, "branch_protection_rule", "")
	checkEvent(t, probot.GitHub.BranchProtectionRule.Deleted, "branch_protection_rule", "deleted")
	checkEvent(t, probot.GitHub.BranchProtectionRule.Created, "branch_protection_rule", "created")
	checkEvent(t, probot.GitHub.BranchProtectionRule.Edited, "branch_protection_rule", "edited")

	h := probot.GitHub.BranchProtectionRule.Handler(func(ctx probot.GitHubBranchProtectionRuleContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.BranchProtectionRuleEvent])(nil)
}

func TestCheckRun(t *testing.T) {
	checkEvent(t, probot.GitHub.CheckRun, "check_run", "")
	checkEvent(t, probot.GitHub.CheckRun.Completed, "check_run", "completed")
	checkEvent(t, probot.GitHub.CheckRun.Created, "check_run", "created")
	checkEvent(t, probot.GitHub.CheckRun.RequestedAction, "check_run", "requested_action")
	checkEvent(t, probot.GitHub.CheckRun.Rerequested, "check_run", "rerequested")

	h := probot.GitHub.CheckRun.Handler(func(ctx probot.GitHubCheckRunContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.CheckRunEvent])(nil)
}

func TestCheckSuite(t *testing.T) {
	checkEvent(t, probot.GitHub.CheckSuite, "check_suite", "")
	checkEvent(t, probot.GitHub.CheckSuite.Completed, "check_suite", "completed")
	checkEvent(t, probot.GitHub.CheckSuite.Requested, "check_suite", "requested")
	checkEvent(t, probot.GitHub.CheckSuite.Rerequested, "check_suite", "rerequested")

	h := probot.GitHub.CheckSuite.Handler(func(ctx probot.GitHubCheckSuiteContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.CheckSuiteEvent])(nil)
}

func TestCommitComment(t *testing.T) {
	checkEvent(t, probot.GitHub.CommitComment, "commit_comment", "")
	checkEvent(t, probot.GitHub.CommitComment.Created, "commit_comment", "created")

	h := probot.GitHub.CommitComment.Handler(func(ctx probot.GitHubCommitCommentContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.CommitCommentEvent])(nil)
}

func TestCreate(t *testing.T) {
	checkEvent(t, probot.GitHub.Create, "create", "")

	h := probot.GitHub.Create.Handler(func(ctx probot.GitHubCreateContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.CreateEvent])(nil)
}

func TestDelete(t *testing.T) {
	checkEvent(t, probot.GitHub.Delete, "delete", "")

	h := probot.GitHub.Delete.Handler(func(ctx probot.GitHubDeleteContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.DeleteEvent])(nil)
}

func TestDeployKey(t *testing.T) {
	checkEvent(t, probot.GitHub.DeployKey, "deploy_key", "")
	checkEvent(t, probot.GitHub.DeployKey.Created, "deploy_key", "created")
	checkEvent(t, probot.GitHub.DeployKey.Deleted, "deploy_key", "deleted")

	h := probot.GitHub.DeployKey.Handler(func(ctx probot.GitHubDeployKeyContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.DeployKeyEvent])(nil)
}

func TestDeploymentStatus(t *testing.T) {
	checkEvent(t, probot.GitHub.DeploymentStatus, "deployment_status", "")
	checkEvent(t, probot.GitHub.DeploymentStatus.Created, "deployment_status", "created")

	h := probot.GitHub.DeploymentStatus.Handler(func(ctx probot.GitHubDeploymentStatusContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.DeploymentStatusEvent])(nil)
}

func TestDeployment(t *testing.T) {
	checkEvent(t, probot.GitHub.Deployment, "deployment", "")
	checkEvent(t, probot.GitHub.Deployment.Created, "deployment", "created")

	h := probot.GitHub.Deployment.Handler(func(ctx probot.GitHubDeploymentContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.DeploymentEvent])(nil)
}

func TestDiscussion(t *testing.T) {
	checkEvent(t, probot.GitHub.Discussion, "discussion", "")
	checkEvent(t, probot.GitHub.Discussion.Answered, "discussion", "answered")
	checkEvent(t, probot.GitHub.Discussion.CategoryChanged, "discussion", "category_changed")
	checkEvent(t, probot.GitHub.Discussion.Created, "discussion", "created")
	checkEvent(t, probot.GitHub.Discussion.Deleted, "discussion", "deleted")
	checkEvent(t, probot.GitHub.Discussion.Edited, "discussion", "edited")
	checkEvent(t, probot.GitHub.Discussion.Labeled, "discussion", "labeled")
	checkEvent(t, probot.GitHub.Discussion.Locked, "discussion", "locked")
	checkEvent(t, probot.GitHub.Discussion.Pinned, "discussion", "pinned")
	checkEvent(t, probot.GitHub.Discussion.Transferred, "discussion", "transferred")
	checkEvent(t, probot.GitHub.Discussion.Unanswered, "discussion", "unanswered")
	checkEvent(t, probot.GitHub.Discussion.Unlabeled, "discussion", "unlabeled")
	checkEvent(t, probot.GitHub.Discussion.Unlocked, "discussion", "unlocked")
	checkEvent(t, probot.GitHub.Discussion.Unpinned, "discussion", "unpinned")

	h := probot.GitHub.Discussion.Handler(func(ctx probot.GitHubDiscussionContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.DiscussionEvent])(nil)
}

func TestFork(t *testing.T) {
	checkEvent(t, probot.GitHub.Fork, "fork", "")

	h := probot.GitHub.Fork.Handler(func(ctx probot.GitHubForkContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.ForkEvent])(nil)
}

func TestGitHubAppAuthorization(t *testing.T) {
	checkEvent(t, probot.GitHub.GitHubAppAuthorization, "github_app_authorization", "")

	h := probot.GitHub.GitHubAppAuthorization.Handler(func(ctx probot.GitHubGitHubAppAuthorizationContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.GitHubAppAuthorizationEvent])(nil)
}

func TestGollum(t *testing.T) {
	checkEvent(t, probot.GitHub.Gollum, "gollum", "")

	h := probot.GitHub.Gollum.Handler(func(ctx probot.GitHubGollumContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.GollumEvent])(nil)
}

func TestInstallationRepositories(t *testing.T) {
	checkEvent(t, probot.GitHub.InstallationRepositories, "installation_repositories", "")
	checkEvent(t, probot.GitHub.InstallationRepositories.Added, "installation_repositories", "added")
	checkEvent(t, probot.GitHub.InstallationRepositories.Removed, "installation_repositories", "removed")

	h := probot.GitHub.InstallationRepositories.Handler(func(ctx probot.GitHubInstallationRepositoriesContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.InstallationRepositoriesEvent])(nil)
}

func TestInstallation(t *testing.T) {
	checkEvent(t, probot.GitHub.Installation, "installation", "")
	checkEvent(t, probot.GitHub.Installation.Created, "installation", "created")
	checkEvent(t, probot.GitHub.Installation.Deleted, "installation", "deleted")
	checkEvent(t, probot.GitHub.Installation.NewPermissionsAccepted, "installation", "new_permissions_accepted")
	checkEvent(t, probot.GitHub.Installation.Suspend, "installation", "suspend")
	checkEvent(t, probot.GitHub.Installation.Unsuspend, "installation", "unsuspend")

	h := probot.GitHub.Installation.Handler(func(ctx probot.GitHubInstallationContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.InstallationEvent])(nil)
}

func TestIssueComment(t *testing.T) {
	checkEvent(t, probot.GitHub.IssueComment, "issue_comment", "")
	checkEvent(t, probot.GitHub.IssueComment.Created, "issue_comment", "created")
	checkEvent(t, probot.GitHub.IssueComment.Deleted, "issue_comment", "deleted")
	checkEvent(t, probot.GitHub.IssueComment.Edited, "issue_comment", "edited")

	h := probot.GitHub.IssueComment.Handler(func(ctx probot.GitHubIssueCommentContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.IssueCommentEvent])(nil)
}

func TestIssues(t *testing.T) {
	checkEvent(t, probot.GitHub.Issues, "issues", "")
	checkEvent(t, probot.GitHub.Issues.Assigned, "issues", "assigned")
	checkEvent(t, probot.GitHub.Issues.Closed, "issues", "closed")
	checkEvent(t, probot.GitHub.Issues.Deleted, "issues", "deleted")
	checkEvent(t, probot.GitHub.Issues.Demilestoned, "issues", "demilestoned")
	checkEvent(t, probot.GitHub.Issues.Edited, "issues", "edited")
	checkEvent(t, probot.GitHub.Issues.Labeled, "issues", "labeled")
	checkEvent(t, probot.GitHub.Issues.Locked, "issues", "locked")
	checkEvent(t, probot.GitHub.Issues.Milestoned, "issues", "milestoned")
	checkEvent(t, probot.GitHub.Issues.Opened, "issues", "opened")
	checkEvent(t, probot.GitHub.Issues.Pinned, "issues", "pinned")
	checkEvent(t, probot.GitHub.Issues.Reopened, "issues", "reopened")
	checkEvent(t, probot.GitHub.Issues.Transferred, "issues", "transferred")
	checkEvent(t, probot.GitHub.Issues.Unassigned, "issues", "unassigned")
	checkEvent(t, probot.GitHub.Issues.Unlabeled, "issues", "unlabeled")
	checkEvent(t, probot.GitHub.Issues.Unlocked, "issues", "unlocked")
	checkEvent(t, probot.GitHub.Issues.Unpinned, "issues", "unpinned")

	h := probot.GitHub.Issues.Handler(func(ctx probot.GitHubIssuesContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.IssuesEvent])(nil)
}

func TestLabel(t *testing.T) {
	checkEvent(t, probot.GitHub.Label, "label", "")
	checkEvent(t, probot.GitHub.Label.Created, "label", "created")
	checkEvent(t, probot.GitHub.Label.Deleted, "label", "deleted")
	checkEvent(t, probot.GitHub.Label.Edited, "label", "edited")

	h := probot.GitHub.Label.Handler(func(ctx probot.GitHubLabelContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.LabelEvent])(nil)
}

func TestMarketplacePurchase(t *testing.T) {
	checkEvent(t, probot.GitHub.MarketplacePurchase, "marketplace_purchase", "")
	checkEvent(t, probot.GitHub.MarketplacePurchase.Cancelled, "marketplace_purchase", "cancelled")
	checkEvent(t, probot.GitHub.MarketplacePurchase.Changed, "marketplace_purchase", "changed")
	checkEvent(t, probot.GitHub.MarketplacePurchase.PendingChange, "marketplace_purchase", "pending_change")
	checkEvent(t, probot.GitHub.MarketplacePurchase.PendingChangeCancelled, "marketplace_purchase", "pending_change_cancelled")
	checkEvent(t, probot.GitHub.MarketplacePurchase.Purchased, "marketplace_purchase", "purchased")

	h := probot.GitHub.MarketplacePurchase.Handler(func(ctx probot.GitHubMarketplacePurchaseContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.MarketplacePurchaseEvent])(nil)
}

func TestMember(t *testing.T) {
	checkEvent(t, probot.GitHub.Member, "member", "")
	checkEvent(t, probot.GitHub.Member.Added, "member", "added")
	checkEvent(t, probot.GitHub.Member.Edited, "member", "edited")
	checkEvent(t, probot.GitHub.Member.Removed, "member", "removed")

	h := probot.GitHub.Member.Handler(func(ctx probot.GitHubMemberContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.MemberEvent])(nil)
}

func TestMembership(t *testing.T) {
	checkEvent(t, probot.GitHub.Membership, "membership", "")
	checkEvent(t, probot.GitHub.Membership.Added, "membership", "added")
	checkEvent(t, probot.GitHub.Membership.Removed, "membership", "removed")

	h := probot.GitHub.Membership.Handler(func(ctx probot.GitHubMembershipContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.MembershipEvent])(nil)
}

func TestMergeGroup(t *testing.T) {
	checkEvent(t, probot.GitHub.MergeGroup, "merge_group", "")

	h := probot.GitHub.MergeGroup.Handler(func(ctx probot.GitHubMergeGroupContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.MergeGroupEvent])(nil)
}

func TestMeta(t *testing.T) {
	checkEvent(t, probot.GitHub.Meta, "meta", "")
	checkEvent(t, probot.GitHub.Meta.Deleted, "meta", "deleted")

	h := probot.GitHub.Meta.Handler(func(ctx probot.GitHubMetaContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.MetaEvent])(nil)
}

func TestMilestone(t *testing.T) {
	checkEvent(t, probot.GitHub.Milestone, "milestone", "")
	checkEvent(t, probot.GitHub.Milestone.Closed, "milestone", "closed")
	checkEvent(t, probot.GitHub.Milestone.Created, "milestone", "created")
	checkEvent(t, probot.GitHub.Milestone.Deleted, "milestone", "deleted")
	checkEvent(t, probot.GitHub.Milestone.Edited, "milestone", "edited")
	checkEvent(t, probot.GitHub.Milestone.Opened, "milestone", "opened")

	h := probot.GitHub.Milestone.Handler(func(ctx probot.GitHubMilestoneContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.MilestoneEvent])(nil)
}

func TestOrgBlock(t *testing.T) {
	checkEvent(t, probot.GitHub.OrgBlock, "org_block", "")
	checkEvent(t, probot.GitHub.OrgBlock.Blocked, "org_block", "blocked")
	checkEvent(t, probot.GitHub.OrgBlock.Unblocked, "org_block", "unblocked")

	h := probot.GitHub.OrgBlock.Handler(func(ctx probot.GitHubOrgBlockContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.OrgBlockEvent])(nil)
}

func TestOrganization(t *testing.T) {
	checkEvent(t, probot.GitHub.Organization, "organization", "")
	checkEvent(t, probot.GitHub.Organization.Deleted, "organization", "deleted")
	checkEvent(t, probot.GitHub.Organization.MemberAdded, "organization", "member_added")
	checkEvent(t, probot.GitHub.Organization.MemberInvited, "organization", "member_invited")
	checkEvent(t, probot.GitHub.Organization.MemberRemoved, "organization", "member_removed")
	checkEvent(t, probot.GitHub.Organization.Renamed, "organization", "renamed")

	h := probot.GitHub.Organization.Handler(func(ctx probot.GitHubOrganizationContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.OrganizationEvent])(nil)
}

func TestPackage(t *testing.T) {
	checkEvent(t, probot.GitHub.Package, "package", "")
	checkEvent(t, probot.GitHub.Package.Published, "package", "published")
	checkEvent(t, probot.GitHub.Package.Updated, "package", "updated")

	h := probot.GitHub.Package.Handler(func(ctx probot.GitHubPackageContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.PackageEvent])(nil)
}

func TestPageBuild(t *testing.T) {
	checkEvent(t, probot.GitHub.PageBuild, "page_build", "")

	h := probot.GitHub.PageBuild.Handler(func(ctx probot.GitHubPageBuildContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.PageBuildEvent])(nil)
}

func TestPing(t *testing.T) {
	checkEvent(t, probot.GitHub.Ping, "ping", "")

	h := probot.GitHub.Ping.Handler(func(ctx probot.GitHubPingContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.PingEvent])(nil)
}

func TestProjectCard(t *testing.T) {
	checkEvent(t, probot.GitHub.ProjectCard, "project_card", "")
	checkEvent(t, probot.GitHub.ProjectCard.Converted, "project_card", "converted")
	checkEvent(t, probot.GitHub.ProjectCard.Created, "project_card", "created")
	checkEvent(t, probot.GitHub.ProjectCard.Deleted, "project_card", "deleted")
	checkEvent(t, probot.GitHub.ProjectCard.Edited, "project_card", "edited")
	checkEvent(t, probot.GitHub.ProjectCard.Moved, "project_card", "moved")

	h := probot.GitHub.ProjectCard.Handler(func(ctx probot.GitHubProjectCardContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.ProjectCardEvent])(nil)
}

func TestProject(t *testing.T) {
	checkEvent(t, probot.GitHub.Project, "project", "")
	checkEvent(t, probot.GitHub.Project.Closed, "project", "closed")
	checkEvent(t, probot.GitHub.Project.Created, "project", "created")
	checkEvent(t, probot.GitHub.Project.Deleted, "project", "deleted")
	checkEvent(t, probot.GitHub.Project.Edited, "project", "edited")
	checkEvent(t, probot.GitHub.Project.Reopened, "project", "reopened")

	h := probot.GitHub.Project.Handler(func(ctx probot.GitHubProjectContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.ProjectEvent])(nil)
}

func TestProjectColumn(t *testing.T) {
	checkEvent(t, probot.GitHub.ProjectColumn, "project_column", "")
	checkEvent(t, probot.GitHub.ProjectColumn.Created, "project_column", "created")
	checkEvent(t, probot.GitHub.ProjectColumn.Deleted, "project_column", "deleted")
	checkEvent(t, probot.GitHub.ProjectColumn.Edited, "project_column", "edited")
	checkEvent(t, probot.GitHub.ProjectColumn.Moved, "project_column", "moved")

	h := probot.GitHub.ProjectColumn.Handler(func(ctx probot.GitHubProjectColumnContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.ProjectColumnEvent])(nil)
}

func TestPublic(t *testing.T) {
	checkEvent(t, probot.GitHub.Public, "public", "")

	h := probot.GitHub.Public.Handler(func(ctx probot.GitHubPublicContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.PublicEvent])(nil)
}

func TestPullRequestReviewComment(t *testing.T) {
	checkEvent(t, probot.GitHub.PullRequestReviewComment, "pull_request_review_comment", "")
	checkEvent(t, probot.GitHub.PullRequestReviewComment.Created, "pull_request_review_comment", "created")
	checkEvent(t, probot.GitHub.PullRequestReviewComment.Deleted, "pull_request_review_comment", "deleted")
	checkEvent(t, probot.GitHub.PullRequestReviewComment.Edited, "pull_request_review_comment", "edited")

	h := probot.GitHub.PullRequestReviewComment.Handler(func(ctx probot.GitHubPullRequestReviewCommentContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.PullRequestReviewCommentEvent])(nil)
}

func TestPullRequestReviewThread(t *testing.T) {
	checkEvent(t, probot.GitHub.PullRequestReviewThread, "pull_request_review_thread", "")
	checkEvent(t, probot.GitHub.PullRequestReviewThread.Resolved, "pull_request_review_thread", "resolved")
	checkEvent(t, probot.GitHub.PullRequestReviewThread.Unresolved, "pull_request_review_thread", "unresolved")

	h := probot.GitHub.PullRequestReviewThread.Handler(func(ctx probot.GitHubPullRequestReviewThreadContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.PullRequestReviewThreadEvent])(nil)
}

func TestPullRequestReview(t *testing.T) {
	checkEvent(t, probot.GitHub.PullRequestReview, "pull_request_review", "")
	checkEvent(t, probot.GitHub.PullRequestReview.Dismissed, "pull_request_review", "dismissed")
	checkEvent(t, probot.GitHub.PullRequestReview.Edited, "pull_request_review", "edited")
	checkEvent(t, probot.GitHub.PullRequestReview.Submitted, "pull_request_review", "submitted")

	h := probot.GitHub.PullRequestReview.Handler(func(ctx probot.GitHubPullRequestReviewContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.PullRequestReviewEvent])(nil)
}

func TestPullRequest(t *testing.T) {
	checkEvent(t, probot.GitHub.PullRequest, "pull_request", "")
	checkEvent(t, probot.GitHub.PullRequest.Assigned, "pull_request", "assigned")
	checkEvent(t, probot.GitHub.PullRequest.AutoMergeDisabled, "pull_request", "auto_merge_disabled")
	checkEvent(t, probot.GitHub.PullRequest.AutoMergeEnabled, "pull_request", "auto_merge_enabled")
	checkEvent(t, probot.GitHub.PullRequest.Closed, "pull_request", "closed")
	checkEvent(t, probot.GitHub.PullRequest.ConvertedToDraft, "pull_request", "converted_to_draft")
	checkEvent(t, probot.GitHub.PullRequest.Demilestoned, "pull_request", "demilestoned")
	checkEvent(t, probot.GitHub.PullRequest.Dequeued, "pull_request", "dequeued")
	checkEvent(t, probot.GitHub.PullRequest.Edited, "pull_request", "edited")
	checkEvent(t, probot.GitHub.PullRequest.Labeled, "pull_request", "labeled")
	checkEvent(t, probot.GitHub.PullRequest.Locked, "pull_request", "locked")
	checkEvent(t, probot.GitHub.PullRequest.Milestoned, "pull_request", "milestoned")
	checkEvent(t, probot.GitHub.PullRequest.Opened, "pull_request", "opened")
	checkEvent(t, probot.GitHub.PullRequest.ReadyForReview, "pull_request", "ready_for_review")
	checkEvent(t, probot.GitHub.PullRequest.Reopened, "pull_request", "reopened")
	checkEvent(t, probot.GitHub.PullRequest.ReviewRequestRemoved, "pull_request", "review_request_removed")
	checkEvent(t, probot.GitHub.PullRequest.ReviewRequested, "pull_request", "review_requested")
	checkEvent(t, probot.GitHub.PullRequest.Synchronize, "pull_request", "synchronize")
	checkEvent(t, probot.GitHub.PullRequest.Unassigned, "pull_request", "unassigned")
	checkEvent(t, probot.GitHub.PullRequest.Unlabeled, "pull_request", "unlabeled")
	checkEvent(t, probot.GitHub.PullRequest.Unlocked, "pull_request", "unlocked")

	h := probot.GitHub.PullRequest.Handler(func(ctx probot.GitHubPullRequestContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.PullRequestEvent])(nil)
}

func TestPush(t *testing.T) {
	checkEvent(t, probot.GitHub.Push, "push", "")

	h := probot.GitHub.Push.Handler(func(ctx probot.GitHubPushContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.PushEvent])(nil)
}

func TestRelease(t *testing.T) {
	checkEvent(t, probot.GitHub.Release, "release", "")
	checkEvent(t, probot.GitHub.Release.Created, "release", "created")
	checkEvent(t, probot.GitHub.Release.Deleted, "release", "deleted")
	checkEvent(t, probot.GitHub.Release.Edited, "release", "edited")
	checkEvent(t, probot.GitHub.Release.Prereleased, "release", "prereleased")
	checkEvent(t, probot.GitHub.Release.Published, "release", "published")
	checkEvent(t, probot.GitHub.Release.Released, "release", "released")
	checkEvent(t, probot.GitHub.Release.Unpublished, "release", "unpublished")

	h := probot.GitHub.Release.Handler(func(ctx probot.GitHubReleaseContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.ReleaseEvent])(nil)
}

func TestRepository(t *testing.T) {
	checkEvent(t, probot.GitHub.Repository, "repository", "")
	checkEvent(t, probot.GitHub.Repository.Archived, "repository", "archived")
	checkEvent(t, probot.GitHub.Repository.Created, "repository", "created")
	checkEvent(t, probot.GitHub.Repository.Deleted, "repository", "deleted")
	checkEvent(t, probot.GitHub.Repository.Edited, "repository", "edited")
	checkEvent(t, probot.GitHub.Repository.Privatized, "repository", "privatized")
	checkEvent(t, probot.GitHub.Repository.Publicized, "repository", "publicized")
	checkEvent(t, probot.GitHub.Repository.Renamed, "repository", "renamed")
	checkEvent(t, probot.GitHub.Repository.Transferred, "repository", "transferred")
	checkEvent(t, probot.GitHub.Repository.Unarchived, "repository", "unarchived")

	h := probot.GitHub.Repository.Handler(func(ctx probot.GitHubRepositoryContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.RepositoryEvent])(nil)
}

func TestRepositoryDispatch(t *testing.T) {
	checkEvent(t, probot.GitHub.RepositoryDispatch, "repository_dispatch", "")

	h := probot.GitHub.RepositoryDispatch.Handler(func(ctx probot.GitHubRepositoryDispatchContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.RepositoryDispatchEvent])(nil)
}

func TestRepositoryVulnerabilityAlert(t *testing.T) {
	checkEvent(t, probot.GitHub.RepositoryVulnerabilityAlert, "repository_vulnerability_alert", "")
	checkEvent(t, probot.GitHub.RepositoryVulnerabilityAlert.Create, "repository_vulnerability_alert", "create")
	checkEvent(t, probot.GitHub.RepositoryVulnerabilityAlert.Dismiss, "repository_vulnerability_alert", "dismiss")
	checkEvent(t, probot.GitHub.RepositoryVulnerabilityAlert.Reopen, "repository_vulnerability_alert", "reopen")
	checkEvent(t, probot.GitHub.RepositoryVulnerabilityAlert.Resolve, "repository_vulnerability_alert", "resolve")

	h := probot.GitHub.RepositoryVulnerabilityAlert.Handler(func(ctx probot.GitHubRepositoryVulnerabilityAlertContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.RepositoryVulnerabilityAlertEvent])(nil)
}

func TestSecretScanningAlert(t *testing.T) {
	checkEvent(t, probot.GitHub.SecretScanningAlert, "secret_scanning_alert", "")
	checkEvent(t, probot.GitHub.SecretScanningAlert.Created, "secret_scanning_alert", "created")
	checkEvent(t, probot.GitHub.SecretScanningAlert.Reopened, "secret_scanning_alert", "reopened")
	checkEvent(t, probot.GitHub.SecretScanningAlert.Resolved, "secret_scanning_alert", "resolved")
	checkEvent(t, probot.GitHub.SecretScanningAlert.Revoked, "secret_scanning_alert", "revoked")

	h := probot.GitHub.SecretScanningAlert.Handler(func(ctx probot.GitHubSecretScanningAlertContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.SecretScanningAlertEvent])(nil)
}

func TestStar(t *testing.T) {
	checkEvent(t, probot.GitHub.Star, "star", "")
	checkEvent(t, probot.GitHub.Star.Created, "star", "created")
	checkEvent(t, probot.GitHub.Star.Deleted, "star", "deleted")

	h := probot.GitHub.Star.Handler(func(ctx probot.GitHubStarContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.StarEvent])(nil)
}

func TestTeamAdd(t *testing.T) {
	checkEvent(t, probot.GitHub.TeamAdd, "team_add", "")

	h := probot.GitHub.TeamAdd.Handler(func(ctx probot.GitHubTeamAddContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.TeamAddEvent])(nil)
}

func TestTeam(t *testing.T) {
	checkEvent(t, probot.GitHub.Team, "team", "")
	checkEvent(t, probot.GitHub.Team.AddedToRepository, "team", "added_to_repository")
	checkEvent(t, probot.GitHub.Team.Created, "team", "created")
	checkEvent(t, probot.GitHub.Team.Deleted, "team", "deleted")
	checkEvent(t, probot.GitHub.Team.Edited, "team", "edited")
	checkEvent(t, probot.GitHub.Team.RemovedFromRepository, "team", "removed_from_repository")

	h := probot.GitHub.Team.Handler(func(ctx probot.GitHubTeamContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.TeamEvent])(nil)
}

func TestWatch(t *testing.T) {
	checkEvent(t, probot.GitHub.Watch, "watch", "")

	h := probot.GitHub.Watch.Handler(func(ctx probot.GitHubWatchContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.WatchEvent])(nil)
}

func TestStatus(t *testing.T) {
	checkEvent(t, probot.GitHub.Status, "status", "")

	h := probot.GitHub.Status.Handler(func(ctx probot.GitHubStatusContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.StatusEvent])(nil)
}

func TestWorkflowDispatch(t *testing.T) {
	checkEvent(t, probot.GitHub.WorkflowDispatch, "workflow_dispatch", "")

	h := probot.GitHub.WorkflowDispatch.Handler(func(ctx probot.GitHubWorkflowDispatchContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.WorkflowDispatchEvent])(nil)
}

func TestWorkflowJob(t *testing.T) {
	checkEvent(t, probot.GitHub.WorkflowJob, "workflow_job", "")
	checkEvent(t, probot.GitHub.WorkflowJob.Completed, "workflow_job", "completed")
	checkEvent(t, probot.GitHub.WorkflowJob.InProgress, "workflow_job", "in_progress")
	checkEvent(t, probot.GitHub.WorkflowJob.Queued, "workflow_job", "queued")
	checkEvent(t, probot.GitHub.WorkflowJob.Waiting, "workflow_job", "waiting")

	h := probot.GitHub.WorkflowJob.Handler(func(ctx probot.GitHubWorkflowJobContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.WorkflowJobEvent])(nil)
}

func TestWorkflowRun(t *testing.T) {
	checkEvent(t, probot.GitHub.WorkflowRun, "workflow_run", "")
	checkEvent(t, probot.GitHub.WorkflowRun.Completed, "workflow_run", "completed")
	checkEvent(t, probot.GitHub.WorkflowRun.InProgress, "workflow_run", "in_progress")
	checkEvent(t, probot.GitHub.WorkflowRun.Requested, "workflow_run", "requested")

	h := probot.GitHub.WorkflowRun.Handler(func(ctx probot.GitHubWorkflowRunContext) {
		fmt.Println(ctx)
	})
	h.(probot.EventHandlerFunc[probot.GitHubClient, github.WorkflowRunEvent])(nil)
}

func checkEvent(t *testing.T, e probot.WebhookEvent, tp, action string) {
	if e.Type() != tp {
		t.Fatal("Failed to check event type", "expect", tp, "actual", e.Type())
	}
	if e.Action() != action {
		t.Fatal("Failed to check event action", "expect", action, "actual", e.Action())
	}
}
