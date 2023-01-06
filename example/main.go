package main

import (
	"context"

	"github.com/google/go-github/v48/github"
	"github.com/spf13/pflag"

	"github.com/airconduct/go-probot"
	probotgh "github.com/airconduct/go-probot/github"
)

func main() {
	app := probot.NewGithubAPP()
	app.AddFlags(pflag.CommandLine)
	pflag.Parse()

	// Add a handler for events "issue_comment.created"
	app.On(probotgh.Event.IssueComment_created).WithHandler(probotgh.IssueCommentHandler(func(ctx probotgh.IssueCommentContext) {
		payload := ctx.Payload()
		ctx.Logger().Info("Get IssueComment event", "payload", payload)
		owner := payload.Repo.Owner.GetLogin()
		repo := payload.Repo.GetName()
		issueNumber := *payload.Issue.Number
		// If any error happen, the error message will be logged and sent as response
		ctx.Must(ctx.Client().Issues.CreateComment(ctx, owner, repo, issueNumber, &github.IssueComment{
			Body: github.String("Reply to this comment."),
		}))
	}))

	// Add a handler for multiple events
	app.On(
		probotgh.Event.PullRequest_opened,      // pull_request.opened
		probotgh.Event.PullRequest_edited,      // pull_request.edited
		probotgh.Event.PullRequest_synchronize, // pull_request.synchronize
		probotgh.Event.PullRequest_labeled,     // pull_request.labeled
		probotgh.Event.PullRequest_assigned,    // pull_request.assigned
	).WithHandler(probotgh.PullRequestHandler(func(ctx probotgh.PullRequestContext) {
		payload := ctx.Payload()
		ctx.Logger().Info("Do something", "action", payload.GetAction(), "PullRequest labels", payload.PullRequest.Labels)
	}))

	if err := app.Run(context.Background()); err != nil {
		panic(err)
	}
}
