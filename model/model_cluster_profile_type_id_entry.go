package nsxt

type ClusterProfileTypeIdEntry struct {
	// key value
	ProfileId string `json:"profile_id"`
	// Supported cluster profiles.
	ResourceType string `json:"resource_type,omitempty"`
}
