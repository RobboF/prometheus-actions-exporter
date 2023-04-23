package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	workflowDuration *prometheus.GaugeVec
	workflowTotals   *prometheus.GaugeVec
)

func InitMetrics() {
	workflowDuration = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "github_workflow_duration",
			Help: "Workflow duration",
		}, []string{"repo_name", "workflow_name"},
	)
	prometheus.MustRegister(workflowDuration)

	workflowTotals = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "github_total_workflowRuns",
			Help: "Number of workflow runs",
		}, []string{"repo_name", "workflow_name", "status"},
	)
	prometheus.MustRegister(workflowTotals)

}
