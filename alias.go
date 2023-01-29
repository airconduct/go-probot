package probot

import (
	"github.com/google/go-github/v48/github"
	"github.com/xanzy/go-gitlab"
)

//go:generate go run tools/codegen/main.go -t github_events.go.tmpl -v github_events.yaml -o github_events.go
//go:generate go run tools/codegen/main.go -t github_events_test.go.tmpl -v github_events.yaml -o github_events_test.go
var Github = &githubEvent{}

type GitClientType interface {
	GithubClient | gitlab.Client
}

type GithubClient = github.Client
type GitlabClient = gitlab.Client
