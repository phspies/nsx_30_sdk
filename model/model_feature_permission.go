package nsxt

// Feature Permission
type FeaturePermission struct {
	// Is execute recommended
	IsExecuteRecommended bool `json:"is_execute_recommended,omitempty"`
	// Feature Name
	FeatureName string `json:"feature_name,omitempty"`
	// Permission
	Permission string `json:"permission"`
	// Is internal
	IsInternal bool `json:"is_internal,omitempty"`
	// Feature Id
	Feature string `json:"feature"`
	// Feature Description
	FeatureDescription string `json:"feature_description,omitempty"`
}
