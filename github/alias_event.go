package github

import "github.com/airconduct/go-probot"

var Event = githubEventSets{
	BranchProtectionRule:         branchProtectionRule(),
	BranchProtectionRule_deleted: branchProtectionRule_deleted(),
	BranchProtectionRule_created: branchProtectionRule_created(),
	BranchProtectionRule_edited:  branchProtectionRule_edited(),

	CheckRun:                  checkRun(),
	CheckRun_completed:        checkRun_completed(),
	CheckRun_created:          checkRun_created(),
	CheckRun_requested_action: checkRun_requested_action(),
	CheckRun_rerequested:      checkRun_rerequested(),

	CheckSuite:             checkSuite(),
	CheckSuite_completed:   checkSuite_completed(),
	CheckSuite_requested:   checkSuite_requested(),
	CheckSuite_rerequested: checkSuite_rerequested(),

	CodeScanningAlert:                    codeScanningAlert(),
	CodeScanningAlert_appeared_in_branch: codeScanningAlert_appeared_in_branch(),
	CodeScanningAlert_closed_by_user:     codeScanningAlert_closed_by_user(),
	CodeScanningAlert_created:            codeScanningAlert_created(),
	CodeScanningAlert_fixed:              codeScanningAlert_fixed(),
	CodeScanningAlert_reopened:           codeScanningAlert_reopened(),
	CodeScanningAlert_reopened_by_user:   codeScanningAlert_reopened_by_user(),

	CommitComment:         commitComment(),
	CommitComment_created: commitComment_created(),

	Create: create(),
	Delete: delete(),

	DeployKey:         deployKey(),
	DeployKeyCreated:  deployKeyCreated(),
	DeployKey_deleted: deployKey_deleted(),

	DeploymentStatus:         deploymentStatus(),
	DeploymentStatus_created: deploymentStatus_created(),

	Deployment:         deployment(),
	Deployment_created: deployment_created(),

	Discussion:                  discussion(),
	Discussion_answered:         discussion_answered(),
	Discussion_category_changed: discussion_category_changed(),
	Discussion_created:          discussion_created(),
	Discussion_deleted:          discussion_deleted(),
	Discussion_edited:           discussion_edited(),
	Discussion_labeled:          discussion_labeled(),
	Discussion_locked:           discussion_locked(),
	Discussion_pinned:           discussion_pinned(),
	Discussion_transferred:      discussion_transferred(),
	Discussion_unanswered:       discussion_unanswered(),
	Discussion_unlabeled:        discussion_unlabeled(),
	Discussion_unlocked:         discussion_unlocked(),
	Discussion_unpinned:         discussion_unpinned(),

	Fork: fork(),

	GithubAppAuthorization: githubAppAuthorization(),

	Gollum: gollum(),

	InstallationRepositories:         installationRepositories(),
	InstallationRepositories_added:   installationRepositories_added(),
	InstallationRepositories_removed: installationRepositories_removed(),

	Installation:                          installation(),
	Installation_created:                  installation_created(),
	Installation_deleted:                  installation_deleted(),
	Installation_new_permissions_accepted: installation_new_permissions_accepted(),
	Installation_suspend:                  installation_suspend(),
	Installation_unsuspend:                installation_unsuspend(),

	IssueComment:         issueComment(),
	IssueComment_created: issueComment_created(),
	IssueComment_deleted: issueComment_deleted(),
	IssueComment_edited:  issueComment_edited(),

	Issues:              issues(),
	Issues_assigned:     issues_assigned(),
	Issues_closed:       issues_closed(),
	Issues_deleted:      issues_deleted(),
	Issues_demilestoned: issues_demilestoned(),
	Issues_edited:       issues_edited(),
	Issues_labeled:      issues_labeled(),
	Issues_locked:       issues_locked(),
	Issues_milestoned:   issues_milestoned(),
	Issues_opened:       issues_opened(),
	Issues_pinned:       issues_pinned(),
	Issues_reopened:     issues_reopened(),
	Issues_transferred:  issues_transferred(),
	Issues_unassigned:   issues_unassigned(),
	Issues_unlabeled:    issues_unlabeled(),
	Issues_unlocked:     issues_unlocked(),
	Issues_unpinned:     issues_unpinned(),

	Label:         label(),
	Lable_created: lable_created(),
	Lable_deleted: lable_deleted(),
	Lable_edited:  lable_edited(),

	MarketplacePurchase:                          marketplacePurchase(),
	MarketplacePurchase_cancelled:                marketplacePurchase_cancelled(),
	MarketplacePurchase_changed:                  marketplacePurchase_changed(),
	MarketplacePurchase_pending_change:           marketplacePurchase_pending_change(),
	MarketplacePurchase_pending_change_cancelled: marketplacePurchase_pending_change_cancelled(),
	MarketplacePurchase_purchased:                marketplacePurchase_purchased(),

	Member:         member(),
	Member_added:   member_added(),
	Member_edited:  member_edited(),
	Member_removed: member_removed(),

	Membership:         membership(),
	Membership_added:   membership_added(),
	Membership_removed: membership_removed(),

	MergeGroup: mergeGroup(),

	Meta:        meta(),
	MetaDeleted: metaDeleted(),

	PullRequestReviewComment:         pullRequestReviewComment(),
	PullRequestReviewComment_created: pullRequestReviewComment_created(),
	PullRequestReviewComment_deleted: pullRequestReviewComment_deleted(),
	PullRequestReviewComment_edited:  pullRequestReviewComment_edited(),

	PullRequestReviewThread:            pullRequestReviewThread(),
	PullRequestReviewThread_resolved:   pullRequestReviewThread_resolved(),
	PullRequestReviewThread_unresolved: pullRequestReviewThread_unresolved(),

	PullRequestReview:           pullRequestReview(),
	PullRequestReview_dismissed: pullRequestReview_dismissed(),
	PullRequestReview_edited:    pullRequestReview_edited(),
	PullRequestReview_submitted: pullRequestReview_submitted(),

	PullRequest:                        pullRequest(),
	PullRequest_assigned:               pullRequest_assigned(),
	PullRequest_auto_merge_disabled:    pullRequest_auto_merge_disabled(),
	PullRequest_auto_merge_enabled:     pullRequest_auto_merge_enabled(),
	PullRequest_closed:                 pullRequest_closed(),
	PullRequest_converted_to_draft:     pullRequest_converted_to_draft(),
	PullRequest_demilestoned:           pullRequest_demilestoned(),
	PullRequest_dequeued:               pullRequest_dequeued(),
	PullRequest_edited:                 pullRequest_edited(),
	PullRequest_labeled:                pullRequest_labeled(),
	PullRequest_locked:                 pullRequest_locked(),
	PullRequest_milestoned:             pullRequest_milestoned(),
	PullRequest_opened:                 pullRequest_opened(),
	PullRequest_ready_for_review:       pullRequest_ready_for_review(),
	PullRequest_reopened:               pullRequest_reopened(),
	PullRequest_review_request_removed: pullRequest_review_request_removed(),
	PullRequest_review_requested:       pullRequest_review_requested(),
	PullRequest_synchronize:            pullRequest_synchronize(),
	PullRequest_unassigned:             pullRequest_unassigned(),
	PullRequest_unlabeled:              pullRequest_unlabeled(),
	PullRequest_unlocked:               pullRequest_unlocked(),

	Push: push(),

	Status: status(),

	WorkflowDispatch: workflowDispatch(),

	WorkflowJob:             workflowJob(),
	WorkflowJob_completed:   workflowJob_completed(),
	WorkflowJob_in_progress: workflowJob_in_progress(),
	WorkflowJob_queued:      workflowJob_queued(),
	WorkflowJob_waiting:     workflowJob_waiting(),

	WorkflowRun:             workflowRun(),
	WorkflowRun_completed:   workflowRun_completed(),
	WorkflowRun_in_progress: workflowRun_in_progress(),
	WorkflowRun_requested:   workflowRun_requested(),
}

type githubEventSets struct {
	BranchProtectionRule         probot.WebhookEvent
	BranchProtectionRule_deleted probot.WebhookEvent
	BranchProtectionRule_created probot.WebhookEvent
	BranchProtectionRule_edited  probot.WebhookEvent

	CheckRun                  probot.WebhookEvent
	CheckRun_completed        probot.WebhookEvent
	CheckRun_created          probot.WebhookEvent
	CheckRun_requested_action probot.WebhookEvent
	CheckRun_rerequested      probot.WebhookEvent

	CheckSuite             probot.WebhookEvent
	CheckSuite_completed   probot.WebhookEvent
	CheckSuite_requested   probot.WebhookEvent
	CheckSuite_rerequested probot.WebhookEvent

	CodeScanningAlert                    probot.WebhookEvent
	CodeScanningAlert_appeared_in_branch probot.WebhookEvent
	CodeScanningAlert_closed_by_user     probot.WebhookEvent
	CodeScanningAlert_created            probot.WebhookEvent
	CodeScanningAlert_fixed              probot.WebhookEvent
	CodeScanningAlert_reopened           probot.WebhookEvent
	CodeScanningAlert_reopened_by_user   probot.WebhookEvent

	CommitComment         probot.WebhookEvent
	CommitComment_created probot.WebhookEvent

	Create probot.WebhookEvent
	Delete probot.WebhookEvent

	DeployKey         probot.WebhookEvent
	DeployKeyCreated  probot.WebhookEvent
	DeployKey_deleted probot.WebhookEvent

	DeploymentStatus         probot.WebhookEvent
	DeploymentStatus_created probot.WebhookEvent

	Deployment         probot.WebhookEvent
	Deployment_created probot.WebhookEvent

	Discussion                  probot.WebhookEvent
	Discussion_answered         probot.WebhookEvent
	Discussion_category_changed probot.WebhookEvent
	Discussion_created          probot.WebhookEvent
	Discussion_deleted          probot.WebhookEvent
	Discussion_edited           probot.WebhookEvent
	Discussion_labeled          probot.WebhookEvent
	Discussion_locked           probot.WebhookEvent
	Discussion_pinned           probot.WebhookEvent
	Discussion_transferred      probot.WebhookEvent
	Discussion_unanswered       probot.WebhookEvent
	Discussion_unlabeled        probot.WebhookEvent
	Discussion_unlocked         probot.WebhookEvent
	Discussion_unpinned         probot.WebhookEvent

	Fork probot.WebhookEvent

	GithubAppAuthorization probot.WebhookEvent

	Gollum probot.WebhookEvent

	InstallationRepositories         probot.WebhookEvent
	InstallationRepositories_added   probot.WebhookEvent
	InstallationRepositories_removed probot.WebhookEvent

	Installation                          probot.WebhookEvent
	Installation_created                  probot.WebhookEvent
	Installation_deleted                  probot.WebhookEvent
	Installation_new_permissions_accepted probot.WebhookEvent
	Installation_suspend                  probot.WebhookEvent
	Installation_unsuspend                probot.WebhookEvent

	IssueComment         probot.WebhookEvent
	IssueComment_created probot.WebhookEvent
	IssueComment_deleted probot.WebhookEvent
	IssueComment_edited  probot.WebhookEvent

	Issues              probot.WebhookEvent
	Issues_assigned     probot.WebhookEvent
	Issues_closed       probot.WebhookEvent
	Issues_deleted      probot.WebhookEvent
	Issues_demilestoned probot.WebhookEvent
	Issues_edited       probot.WebhookEvent
	Issues_labeled      probot.WebhookEvent
	Issues_locked       probot.WebhookEvent
	Issues_milestoned   probot.WebhookEvent
	Issues_opened       probot.WebhookEvent
	Issues_pinned       probot.WebhookEvent
	Issues_reopened     probot.WebhookEvent
	Issues_transferred  probot.WebhookEvent
	Issues_unassigned   probot.WebhookEvent
	Issues_unlabeled    probot.WebhookEvent
	Issues_unlocked     probot.WebhookEvent
	Issues_unpinned     probot.WebhookEvent

	Label         probot.WebhookEvent
	Lable_created probot.WebhookEvent
	Lable_deleted probot.WebhookEvent
	Lable_edited  probot.WebhookEvent

	MarketplacePurchase                          probot.WebhookEvent
	MarketplacePurchase_cancelled                probot.WebhookEvent
	MarketplacePurchase_changed                  probot.WebhookEvent
	MarketplacePurchase_pending_change           probot.WebhookEvent
	MarketplacePurchase_pending_change_cancelled probot.WebhookEvent
	MarketplacePurchase_purchased                probot.WebhookEvent

	Member         probot.WebhookEvent
	Member_added   probot.WebhookEvent
	Member_edited  probot.WebhookEvent
	Member_removed probot.WebhookEvent

	Membership         probot.WebhookEvent
	Membership_added   probot.WebhookEvent
	Membership_removed probot.WebhookEvent

	MergeGroup probot.WebhookEvent

	Meta        probot.WebhookEvent
	MetaDeleted probot.WebhookEvent

	PullRequestReviewComment         probot.WebhookEvent
	PullRequestReviewComment_created probot.WebhookEvent
	PullRequestReviewComment_deleted probot.WebhookEvent
	PullRequestReviewComment_edited  probot.WebhookEvent

	PullRequestReviewThread            probot.WebhookEvent
	PullRequestReviewThread_resolved   probot.WebhookEvent
	PullRequestReviewThread_unresolved probot.WebhookEvent

	PullRequestReview           probot.WebhookEvent
	PullRequestReview_dismissed probot.WebhookEvent
	PullRequestReview_edited    probot.WebhookEvent
	PullRequestReview_submitted probot.WebhookEvent

	PullRequest                        probot.WebhookEvent
	PullRequest_assigned               probot.WebhookEvent
	PullRequest_auto_merge_disabled    probot.WebhookEvent
	PullRequest_auto_merge_enabled     probot.WebhookEvent
	PullRequest_closed                 probot.WebhookEvent
	PullRequest_converted_to_draft     probot.WebhookEvent
	PullRequest_demilestoned           probot.WebhookEvent
	PullRequest_dequeued               probot.WebhookEvent
	PullRequest_edited                 probot.WebhookEvent
	PullRequest_labeled                probot.WebhookEvent
	PullRequest_locked                 probot.WebhookEvent
	PullRequest_milestoned             probot.WebhookEvent
	PullRequest_opened                 probot.WebhookEvent
	PullRequest_ready_for_review       probot.WebhookEvent
	PullRequest_reopened               probot.WebhookEvent
	PullRequest_review_request_removed probot.WebhookEvent
	PullRequest_review_requested       probot.WebhookEvent
	PullRequest_synchronize            probot.WebhookEvent
	PullRequest_unassigned             probot.WebhookEvent
	PullRequest_unlabeled              probot.WebhookEvent
	PullRequest_unlocked               probot.WebhookEvent

	Push probot.WebhookEvent

	Status probot.WebhookEvent

	WorkflowDispatch probot.WebhookEvent

	WorkflowJob             probot.WebhookEvent
	WorkflowJob_completed   probot.WebhookEvent
	WorkflowJob_in_progress probot.WebhookEvent
	WorkflowJob_queued      probot.WebhookEvent
	WorkflowJob_waiting     probot.WebhookEvent

	WorkflowRun             probot.WebhookEvent
	WorkflowRun_completed   probot.WebhookEvent
	WorkflowRun_in_progress probot.WebhookEvent
	WorkflowRun_requested   probot.WebhookEvent
}
