package capabilities

type Capability struct {
	Id         *string               `json:"id,omitempty"`
	Name       *string               `json:"name,omitempty"`
	Properties *CapabilityProperties `json:"properties,omitempty"`
	SystemData *SystemData           `json:"systemData,omitempty"`
	Type       *string               `json:"type,omitempty"`
}
