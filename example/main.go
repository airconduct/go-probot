package main

import (
	"context"

	"github.com/google/go-github/v48/github"
	"github.com/spf13/pflag"

	"github.com/airconduct/go-probot"
)

func main() {
	app := probot.NewGithubAPP()
	app.AddFlags(pflag.CommandLine)
	pflag.Parse()

	// Add a handler for events "issue_comment.created"
	app.On(probot.Github.IssueComment.Created).WithHandler(probot.Github.IssueComment.Handler(func(ctx probot.IssueCommentContext) {
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
		probot.Github.PullRequest.Opened,      // pull_request.opened
		probot.Github.PullRequest.Edited,      // pull_request.edited
		probot.Github.PullRequest.Synchronize, // pull_request.synchronize
		probot.Github.PullRequest.Labeled,     // pull_request.labeled
		probot.Github.PullRequest.Assigned,    // pull_request.assigned
	).WithHandler(probot.Github.PullRequest.Handler(func(ctx probot.PullRequestContext) {
		payload := ctx.Payload()
		ctx.Logger().Info("Do something", "action", payload.GetAction(), "PullRequest labels", payload.PullRequest.Labels)
	}))

	// Add a handler for event "pull_request_review" with all action type
	app.On(probot.Github.PullRequestReview).WithHandler(probot.Github.PullRequestReview.Handler(func(ctx probot.PullRequestReviewContext) {
		payload := ctx.Payload()
		ctx.Logger().Info("Do something", "action", payload.GetAction(), "body", payload.Review.GetBody())
	}))

	if err := app.Run(context.Background()); err != nil {
		panic(err)
	}
}
