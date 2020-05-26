package nsxt

// Details of container cluster node i.e. container host.
type ContainerClusterNode struct {
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
	// Network status of container cluster node.
	NetworkStatus string `json:"network_status,omitempty"`
	// External identifier of the container cluster.
	ContainerClusterId string `json:"container_cluster_id,omitempty"`
	// List of IP addresses of container cluster node.
	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Array of additional specific properties of container cluster node in key-value format. 
	OriginProperties []KeyValuePair `json:"origin_properties,omitempty"`
	// External identifier of the container cluster node in K8S/PAS. 
	ExternalId string `json:"external_id"`
	// List of network errors related to container cluster node.
	NetworkErrors []NetworkError `json:"network_errors,omitempty"`
}
