package metrics

import (
	"github.com/google/go-github/v51/github"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	workflowDuration *prometheus.GaugeVec
)

func InitMetrics() {
	workflowDuration = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "github_workflow_duration",
			Help: "Workflow duration",
		}, []string{"repo_name", "workflow_name"},
	)
	prometheus.MustRegister(workflowDuration)
}
func SetWorkflowDuration(workflowRuns *github.WorkflowRuns) {
	for _, workflow := range workflowRuns.WorkflowRuns {
		// fmt.Printf("\n%+v\n", workflow.Actor)
		duration := workflow.UpdatedAt.Time.Sub(workflow.RunStartedAt.Time).Seconds()
		workflowDuration.WithLabelValues(*workflow.Repository.Name, *workflow.Name).Set(duration)
	}
}
