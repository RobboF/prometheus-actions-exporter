package main

import "github.com/prometheus/client_golang/prometheus"

var (
	workflowDuration *prometheus.GaugeVec
)

func initMetrics() {
	workflowDuration = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "github_workflow_duration",
			Help: "Workflow duration",
		}, []string{"repo_name", "workflow_name"},
	)
}
