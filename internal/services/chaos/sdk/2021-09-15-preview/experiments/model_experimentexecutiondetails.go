package experiments

type ExperimentExecutionDetails struct {
	Id         *string                               `json:"id,omitempty"`
	Name       *string                               `json:"name,omitempty"`
	Properties *ExperimentExecutionDetailsProperties `json:"properties,omitempty"`
	Type       *string                               `json:"type,omitempty"`
}
