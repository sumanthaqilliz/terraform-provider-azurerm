package experiments

type Branch struct {
	Actions []Action `json:"actions"`
	Name    string   `json:"name"`
}
