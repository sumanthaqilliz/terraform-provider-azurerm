package capabilitytypes

type CapabilityTypeProperties struct {
	Description      *string `json:"description,omitempty"`
	DisplayName      *string `json:"displayName,omitempty"`
	ParametersSchema *string `json:"parametersSchema,omitempty"`
	Publisher        *string `json:"publisher,omitempty"`
	TargetType       *string `json:"targetType,omitempty"`
	Urn              *string `json:"urn,omitempty"`
}
