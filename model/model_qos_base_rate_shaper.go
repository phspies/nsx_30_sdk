package nsxt

// A shaper configuration entry that specifies type and metrics
type QosBaseRateShaper struct {
	Enabled bool `json:"enabled"`
	ResourceType string `json:"resource_type"`
}
