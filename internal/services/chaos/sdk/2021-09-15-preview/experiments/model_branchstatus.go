package experiments

type BranchStatus struct {
	Actions *[]ActionStatus `json:"actions,omitempty"`
	Id      *string         `json:"id,omitempty"`
	Name    *string         `json:"name,omitempty"`
	Status  *string         `json:"status,omitempty"`
}
