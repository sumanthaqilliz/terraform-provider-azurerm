package experiments

type StepStatus struct {
	Branches *[]BranchStatus `json:"branches,omitempty"`
	Id       *string         `json:"id,omitempty"`
	Name     *string         `json:"name,omitempty"`
	Status   *string         `json:"status,omitempty"`
}
