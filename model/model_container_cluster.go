package nsxt

// Details of container cluster.
type ContainerCluster struct {
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
	// Network status of container cluster.
	NetworkStatus string `json:"network_status,omitempty"`
	Infrastructure *ContainerInfrastructureInfo `json:"infrastructure,omitempty"`
	// Type of the container cluster. In case of creating container cluster first time, it is expected to pass the valid cluster-type. In case of update, if there is no change in cluster-type, then this field can be omitted in the request. 
	ClusterType string `json:"cluster_type,omitempty"`
	// Array of additional specific properties of container cluster in key-value format. 
	OriginProperties []KeyValuePair `json:"origin_properties,omitempty"`
	// External identifier of the container cluster.
	ExternalId string `json:"external_id,omitempty"`
	// List of network errors related to container cluster.
	NetworkErrors []NetworkError `json:"network_errors,omitempty"`
}
