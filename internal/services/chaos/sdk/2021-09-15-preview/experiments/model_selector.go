package experiments

type Selector struct {
	Id      string            `json:"id"`
	Targets []TargetReference `json:"targets"`
	Type    SelectorType      `json:"type"`
}
