package nsxt

// list of feature usage items
type FeatureUsageList struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Feature Usage List
	FeatureUsageInfo []FeatureUsage `json:"feature_usage_info,omitempty"`
}
