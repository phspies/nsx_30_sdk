package nsxt

// Network policy applied to container.
type ContainerNetworkPolicy struct {
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
	// Network status of container network policy.
	NetworkStatus string `json:"network_status,omitempty"`
	// Identifier of the container cluster this network policy belongs to.
	ContainerClusterId string `json:"container_cluster_id,omitempty"`
	// Type e.g. Network Policy, ASG.
	PolicyType string `json:"policy_type,omitempty"`
	// Array of additional specific properties of container network policy in key-value format. 
	OriginProperties []KeyValuePair `json:"origin_properties,omitempty"`
	// Identifier of the container network policy.
	ExternalId string `json:"external_id"`
	// Identifier of the project which this network policy belongs to.
	ContainerProjectId string `json:"container_project_id,omitempty"`
	// List of network errors related to container network policy.
	NetworkErrors []NetworkError `json:"network_errors,omitempty"`
	// Container network policy specification.
	Spec string `json:"spec,omitempty"`
}
