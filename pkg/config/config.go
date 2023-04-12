package config

import (
	"context"
	"os"

	"github.com/google/go-github/v51/github"
)

var (
	Port         int
	GithubClient *github.Client
	PollRate     int
)

func InitConfig() {
	PollRate = 60 // Poll rate of Github in seconds
	Port = 2112
	GithubClient = github.NewTokenClient(context.Background(), os.Getenv("GITHUB_PAT_TOKEN"))

}
