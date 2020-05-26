package nsxt

// Feature stack collection configuration list result
type FeatureStackCollectionConfigurationList struct {
	// The complete set of feature stack data collection configurations
	Results []FeatureStackCollectionConfiguration `json:"results,omitempty"`
}
