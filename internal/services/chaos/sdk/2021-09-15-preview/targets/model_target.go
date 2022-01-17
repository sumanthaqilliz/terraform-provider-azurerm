package targets

type Target struct {
	Id         *string     `json:"id,omitempty"`
	Location   *string     `json:"location,omitempty"`
	Name       *string     `json:"name,omitempty"`
	Properties interface{} `json:"properties"`
	SystemData *SystemData `json:"systemData,omitempty"`
	Type       *string     `json:"type,omitempty"`
}
