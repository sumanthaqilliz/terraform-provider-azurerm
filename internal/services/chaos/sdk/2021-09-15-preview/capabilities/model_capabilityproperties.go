package capabilities

type CapabilityProperties struct {
	Description      *string `json:"description,omitempty"`
	ParametersSchema *string `json:"parametersSchema,omitempty"`
	Publisher        *string `json:"publisher,omitempty"`
	TargetType       *string `json:"targetType,omitempty"`
	Urn              *string `json:"urn,omitempty"`
}
