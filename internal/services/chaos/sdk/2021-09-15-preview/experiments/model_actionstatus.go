package experiments

type ActionStatus struct {
	Id      *string                                             `json:"id,omitempty"`
	Name    *string                                             `json:"name,omitempty"`
	Status  *string                                             `json:"status,omitempty"`
	Targets *[]ExperimentExecutionActionTargetDetailsProperties `json:"targets,omitempty"`
}
