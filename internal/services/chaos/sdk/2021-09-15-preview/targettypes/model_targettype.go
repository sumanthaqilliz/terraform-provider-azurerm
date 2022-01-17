package targettypes

type TargetType struct {
	Id         *string              `json:"id,omitempty"`
	Location   *string              `json:"location,omitempty"`
	Name       *string              `json:"name,omitempty"`
	Properties TargetTypeProperties `json:"properties"`
	SystemData *SystemData          `json:"systemData,omitempty"`
	Type       *string              `json:"type,omitempty"`
}
