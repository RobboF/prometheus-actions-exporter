package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/google/go-github/v51/github"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	initMetrics()
	initConfig()
	getWorkflowsByRepo("homepage", "RobboF")
	print("server started on port: " + strconv.Itoa(Port))
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+strconv.Itoa(Port), nil)
}

func getWorkflowsByRepo(repo_name string, owner string) {
	opt := &github.ListWorkflowRunsOptions{
		ListOptions: github.ListOptions{PerPage: 200},
	}
	workflows, _, _ := githubClient.Actions.ListRepositoryWorkflowRuns(context.Background(), owner, repo_name, opt)

	for _, workflow := range workflows.WorkflowRuns {
		duration := workflow.UpdatedAt.Time.Sub(workflow.RunStartedAt.Time).Seconds()
		workflowDuration.WithLabelValues(repo_name, *workflow.Name).Set(duration)
	}

}
