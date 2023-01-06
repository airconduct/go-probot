package probot

import "github.com/google/go-github/v48/github"

type GitClientType interface {
	GithubClient
}

type GithubClient = github.Client
