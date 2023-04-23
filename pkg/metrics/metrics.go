package metrics

import (
	"github.com/google/go-github/v51/github"
)

func SetWorkflowDuration(workflowRuns *github.WorkflowRuns) {
	for _, workflow := range workflowRuns.WorkflowRuns {
		// fmt.Printf("\n%+v\n", workflow.Actor)
		duration := workflow.UpdatedAt.Time.Sub(workflow.RunStartedAt.Time).Seconds()
		workflowDuration.WithLabelValues(*workflow.Repository.Name, *workflow.Name).Set(duration)
	}
}

func SetWorkflowTotals(workflowRuns *github.WorkflowRuns) {
	for _, workflow := range workflowRuns.WorkflowRuns {
		workflowTotals.WithLabelValues(*workflow.Repository.Name, *workflow.Name, *workflow.Status).Inc()

	}
}
