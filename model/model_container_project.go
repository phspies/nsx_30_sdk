package nsxt

// Details of org/namespace within a container cluster.
type ContainerProject struct {
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
	// Network status of container project.
	NetworkStatus string `json:"network_status,omitempty"`
	// Identifier of the container cluster to which this project/namespace belongs.
	ContainerClusterId string `json:"container_cluster_id,omitempty"`
	// Array of additional specific properties of container project in key-value format. 
	OriginProperties []KeyValuePair `json:"origin_properties,omitempty"`
	// External identifier of the container project.
	ExternalId string `json:"external_id"`
	// List of network errors related to container project.
	NetworkErrors []NetworkError `json:"network_errors,omitempty"`
}
