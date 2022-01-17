package capabilities

type CapabilityType struct {
	Id         *string                   `json:"id,omitempty"`
	Location   *string                   `json:"location,omitempty"`
	Name       *string                   `json:"name,omitempty"`
	Properties *CapabilityTypeProperties `json:"properties,omitempty"`
	SystemData *SystemData               `json:"systemData,omitempty"`
	Type       *string                   `json:"type,omitempty"`
}
