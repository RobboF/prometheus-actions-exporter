package githubApi

import (
	"context"
	"fmt"

	"github.com/RobboF/prometheus-actions-exporter/pkg/config"
	"github.com/google/go-github/v51/github"
)

func GetWorkflowRunsByRepo(repo_name string, owner string, workflows chan<- *github.WorkflowRuns) {
	fmt.Printf("Collecting Repo workflows for %s\n", repo_name)
	opt := &github.ListWorkflowRunsOptions{
		ListOptions: github.ListOptions{PerPage: 200},
	}
	githubWorkflows, _, _ := config.GithubClient.Actions.ListRepositoryWorkflowRuns(context.Background(), owner, repo_name, opt)
	workflows <- githubWorkflows

}
