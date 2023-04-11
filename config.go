package main

import "github.com/google/go-github/v51/github"

var (
	Port         int
	githubClient *github.Client
)

func initConfig() {
	Port = 2112
	githubClient = github.NewClient(nil)

}
