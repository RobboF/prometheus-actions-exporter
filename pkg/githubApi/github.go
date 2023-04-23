package githubApi

import (
	"context"
	"fmt"
	"log"

	"github.com/RobboF/prometheus-actions-exporter/pkg/config"
	"github.com/google/go-github/v51/github"
)

func GetWorkflowRunsByRepo(repo_name string, owner string) (workflows *github.WorkflowRuns) {
	fmt.Printf("Collecting Repo workflows for %s\n", repo_name)
	opt := &github.ListWorkflowRunsOptions{
		ListOptions: github.ListOptions{PerPage: 200},
	}
	githubWorkflows, _, err := config.GithubClient.Actions.ListRepositoryWorkflowRuns(context.Background(), owner, repo_name, opt)
	if err != nil {
		log.Fatal(err)
	}
	return githubWorkflows

}

func GetReposForUser(owner string) (repos []*github.Repository) {
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 200},
	}
	githubRepos, _, err := config.GithubClient.Repositories.List(context.Background(), owner, opt)
	if err != nil {
		log.Fatal(err)
	}
	return githubRepos
}

func GetReposForOrg(owner string) (repos []*github.Repository) {
	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 200},
	}
	githubRepos, _, err := config.GithubClient.Repositories.ListByOrg(context.Background(), owner, opt)
	if err != nil {
		log.Fatal(err)
	}
	return githubRepos
}
