package config

import (
	"context"
	"os"

	"github.com/google/go-github/v51/github"
)

var (
	Port          int
	GithubClient  *github.Client
	PollRate      int
	GithubUser    string
	GithubUserSet bool
	GithubOrgSet  bool
	GithubOrg     string
)

func InitConfig() {
	PollRate = 60 // Poll rate of Github in seconds
	Port = 2112
	GithubClient = github.NewTokenClient(context.Background(), os.Getenv("GITHUB_PAT_TOKEN"))
	GithubUser, GithubUserSet = os.LookupEnv("GITHUB_USER")
	GithubOrg, GithubOrgSet = os.LookupEnv("GITHUB_ORG")

	if GithubOrgSet && GithubUserSet {
		print("Both GITHUB_USER and GITHUB_ORG set, set one but not both\n")
		os.Exit(1)
	} else if !GithubOrgSet && !GithubUserSet {
		print("Please set either GITHUB_USER or GITHUB_ORG\n")
		os.Exit(1)
	}

}
