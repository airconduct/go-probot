package probot

import (
	"github.com/google/go-github/v48/github"
	"github.com/xanzy/go-gitlab"
)

//go:generate go run tools/codegen/main.go -t github_events.go.tmpl -v github_events.yaml -o github_events.go
//go:generate go run tools/codegen/main.go -t github_events_test.go.tmpl -v github_events.yaml -o github_events_test.go
var GitHub = &githubEvent{}

//go:generate go run tools/codegen/main.go -t gitlab_events.go.tmpl -v gitlab_events.yaml -o gitlab_events.go
var GitLab = &gitlabEvent{}

type GitClientType interface {
	GitHubClient | GitLabClient
}

type GitHubClient = github.Client
type GitLabClient = gitlab.Client
