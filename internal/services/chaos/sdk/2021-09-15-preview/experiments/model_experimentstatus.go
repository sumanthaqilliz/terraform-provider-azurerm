package experiments

type ExperimentStatus struct {
	Id         *string                     `json:"id,omitempty"`
	Name       *string                     `json:"name,omitempty"`
	Properties *ExperimentStatusProperties `json:"properties,omitempty"`
	Type       *string                     `json:"type,omitempty"`
}
