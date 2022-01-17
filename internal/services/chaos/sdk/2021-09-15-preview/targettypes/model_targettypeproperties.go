package targettypes

type TargetTypeProperties struct {
	Description      *string   `json:"description,omitempty"`
	DisplayName      *string   `json:"displayName,omitempty"`
	PropertiesSchema *string   `json:"propertiesSchema,omitempty"`
	ResourceTypes    *[]string `json:"resourceTypes,omitempty"`
}
