package nsxt

type TransportZoneProfileTypeIdEntry struct {
	// profile id of the resource type
	ProfileId string `json:"profile_id"`
	// Selects the type of the transport zone profile
	ResourceType string `json:"resource_type,omitempty"`
}
