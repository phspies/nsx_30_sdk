package nsxt

// Stores the information about cloud native service instance.
type CloudNativeServiceInstance struct {
	// Timestamp of last modification
	LastSyncTime int64 `json:"_last_sync_time,omitempty"`
	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`
	// Description of this resource
	Description string `json:"description,omitempty"`
	// The type of this resource.
	ResourceType string `json:"resource_type"`
	// Opaque identifiers meaningful to the API user
	Tags []Tag `json:"tags,omitempty"`
	// Type of cloud native service.
	ServiceType string `json:"service_type,omitempty"`
	Source *ResourceReference `json:"source,omitempty"`
	// Id of service instance fetched from public cloud. 
	ExternalId string `json:"external_id,omitempty"`
}
