package nsxt

// Container application instance within a project.
type ContainerApplicationInstance struct {
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
	// Status of the container application instance.
	Status string `json:"status,omitempty"`
	// Network status of container application instance.
	NetworkStatus string `json:"network_status,omitempty"`
	// Identifier of the container cluster this application instance belongs to.
	ContainerClusterId string `json:"container_cluster_id,omitempty"`
	// Cluster node id where application instance is running.
	ClusterNodeId string `json:"cluster_node_id,omitempty"`
	// Identifier of the container application instance on container cluster.
	ExternalId string `json:"external_id"`
	// Array of additional specific properties of container application instance in key-value format. 
	OriginProperties []KeyValuePair `json:"origin_properties,omitempty"`
	// List of identifiers of the container application.
	ContainerApplicationIds []string `json:"container_application_ids,omitempty"`
	// Identifier of the container project which this container application instance belongs to. 
	ContainerProjectId string `json:"container_project_id,omitempty"`
	// List of network errors related to container application instance.
	NetworkErrors []NetworkError `json:"network_errors,omitempty"`
}
