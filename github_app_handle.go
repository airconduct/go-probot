// Code generated by codegen. DO NOT EDIT.

package probot

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/v48/github"
)

func (app *githubApp) handelEvent(
	ctx context.Context,
	w http.ResponseWriter, r *http.Request,
	eventType string, rawPayload []byte,
) (handlerKey string) {
	switch string(eventType) {
	case GitHub.BranchProtectionRule.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.BranchProtectionRuleEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.BranchProtectionRuleEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.BranchProtectionRuleEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.CheckRun.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.CheckRunEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.CheckRunEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.CheckRunEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.CheckSuite.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.CheckSuiteEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.CheckSuiteEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.CheckSuiteEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.CommitComment.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.CommitCommentEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.CommitCommentEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.CommitCommentEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.Create.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.CreateEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.CreateEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.CreateEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.Delete.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.DeleteEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.DeleteEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.DeleteEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.DeployKey.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.DeployKeyEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.DeployKeyEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.DeployKeyEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.DeploymentStatus.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.DeploymentStatusEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.DeploymentStatusEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.DeploymentStatusEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.Deployment.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.DeploymentEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.DeploymentEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.DeploymentEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.Discussion.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.DiscussionEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.DiscussionEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.DiscussionEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.Fork.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.ForkEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.ForkEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.ForkEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.GitHubAppAuthorization.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.GitHubAppAuthorizationEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.GitHubAppAuthorizationEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.GitHubAppAuthorizationEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.Gollum.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.GollumEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.GollumEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.GollumEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.InstallationRepositories.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.InstallationRepositoriesEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.InstallationRepositoriesEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.InstallationRepositoriesEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.Installation.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.InstallationEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.InstallationEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.InstallationEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.IssueComment.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.IssueCommentEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.IssueCommentEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.IssueCommentEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.Issues.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.IssuesEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.IssuesEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.IssuesEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.Label.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.LabelEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.LabelEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.LabelEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.MarketplacePurchase.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.MarketplacePurchaseEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.MarketplacePurchaseEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.MarketplacePurchaseEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.Member.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.MemberEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.MemberEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.MemberEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.Membership.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.MembershipEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.MembershipEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.MembershipEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.MergeGroup.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.MergeGroupEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.MergeGroupEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.MergeGroupEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.Meta.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.MetaEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.MetaEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.MetaEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.Milestone.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.MilestoneEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.MilestoneEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.MilestoneEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.OrgBlock.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.OrgBlockEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.OrgBlockEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.OrgBlockEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.Organization.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.OrganizationEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.OrganizationEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.OrganizationEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.Package.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.PackageEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.PackageEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.PackageEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.PageBuild.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.PageBuildEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.PageBuildEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.PageBuildEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.Ping.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.PingEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.PingEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.PingEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.ProjectCard.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.ProjectCardEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.ProjectCardEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.ProjectCardEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.Project.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.ProjectEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.ProjectEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.ProjectEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.ProjectColumn.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.ProjectColumnEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.ProjectColumnEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.ProjectColumnEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.Public.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.PublicEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.PublicEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.PublicEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.PullRequestReviewComment.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.PullRequestReviewCommentEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.PullRequestReviewCommentEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.PullRequestReviewCommentEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.PullRequestReviewThread.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.PullRequestReviewThreadEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.PullRequestReviewThreadEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.PullRequestReviewThreadEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.PullRequestReview.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.PullRequestReviewEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.PullRequestReviewEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.PullRequestReviewEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.PullRequest.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.PullRequestEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.PullRequestEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.PullRequestEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.Push.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.PushEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.PushEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.PushEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.Release.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.ReleaseEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.ReleaseEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.ReleaseEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.Repository.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.RepositoryEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.RepositoryEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.RepositoryEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.RepositoryDispatch.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.RepositoryDispatchEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.RepositoryDispatchEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.RepositoryDispatchEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.RepositoryVulnerabilityAlert.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.RepositoryVulnerabilityAlertEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.RepositoryVulnerabilityAlertEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.RepositoryVulnerabilityAlertEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.SecretScanningAlert.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.SecretScanningAlertEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.SecretScanningAlertEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.SecretScanningAlertEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.Star.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.StarEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.StarEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.StarEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.TeamAdd.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.TeamAddEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.TeamAddEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.TeamAddEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.Team.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.TeamEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.TeamEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.TeamEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.Watch.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.WatchEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.WatchEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.WatchEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.Status.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.StatusEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.StatusEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.StatusEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.WorkflowDispatch.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.WorkflowDispatchEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.WorkflowDispatchEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.WorkflowDispatchEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.WorkflowJob.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.WorkflowJobEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.WorkflowJobEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.WorkflowJobEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}
	case GitHub.WorkflowRun.Type():
		if err := genericHandleFunc(
			ctx, app.logger, string(eventType), func(e string, p *github.WorkflowRunEvent) string {
				return getHandlerKey(e, p)
			},
			func(p *github.WorkflowRunEvent) error {
				return parseWebHook(string(eventType), rawPayload, p)
			},
			func(payload *github.WorkflowRunEvent) (*github.Client, GitGraphQLClient, error) {
				handlerKey = getHandlerKey(string(eventType), payload)
				return app.getClient(payload.GetInstallation().GetID())
			},
			app.handlers,
		); err != nil {
			app.handleError(w, err, http.StatusBadRequest)
			return
		}

	default:
		app.handleError(w, fmt.Errorf("event %s not found", eventType), http.StatusNotFound)
	}
	return
}
