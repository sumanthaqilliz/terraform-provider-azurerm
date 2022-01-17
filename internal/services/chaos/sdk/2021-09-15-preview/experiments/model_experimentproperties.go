package experiments

type ExperimentProperties struct {
	Selectors       []Selector `json:"selectors"`
	StartOnCreation *bool      `json:"startOnCreation,omitempty"`
	Steps           []Step     `json:"steps"`
}
