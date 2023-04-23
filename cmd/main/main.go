package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/RobboF/prometheus-actions-exporter/pkg/config"
	"github.com/RobboF/prometheus-actions-exporter/pkg/githubApi"
	"github.com/RobboF/prometheus-actions-exporter/pkg/metrics"
	"github.com/google/go-github/v51/github"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Intiate Metrics and Configuration
	config.InitConfig()
	metrics.InitMetrics()

	// Create a goroutine to loop
	go interval()

	fmt.Printf("server started on port: " + strconv.Itoa(config.Port) + "\n")
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+strconv.Itoa(config.Port), nil)
}

func interval() {
	ticker := time.NewTicker(time.Duration(config.PollRate) * time.Second)

	for ; true; <-ticker.C {

		var repos []*github.Repository
		var owner string
		if config.GithubOrgSet {
			repos = githubApi.GetReposForOrg(config.GithubOrg)
			owner = config.GithubOrg

		} else if config.GithubUserSet {
			repos = githubApi.GetReposForUser(config.GithubUser)
			owner = config.GithubUser
		}
		for _, repo := range repos {
			fmt.Printf("\n%+v\n", *repo.Name)
			workflows := githubApi.GetWorkflowRunsByRepo(*repo.Name, owner)
			metrics.SetWorkflowDuration(workflows)
			metrics.SetWorkflowTotals(workflows)

		}

	}
}
