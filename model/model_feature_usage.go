package nsxt

// feature usage result item
type FeatureUsage struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Capacity Usage List
	CapacityUsage []CapacityUsage `json:"capacity_usage,omitempty"`
	// name of the feature
	Feature string `json:"feature,omitempty"`
}
