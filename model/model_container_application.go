package nsxt

// Container application within a project.
type ContainerApplication struct {
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
	// Status of the container application.
	Status string `json:"status,omitempty"`
	// Network status of container application.
	NetworkStatus string `json:"network_status,omitempty"`
	// Identifier of the container cluster this container application belongs to.
	ContainerClusterId string `json:"container_cluster_id,omitempty"`
	// Array of additional specific properties of container application in key-value format. 
	OriginProperties []KeyValuePair `json:"origin_properties,omitempty"`
	// Identifier of the container application on container cluster e.g. PCF app id, k8s service id. 
	ExternalId string `json:"external_id"`
	// Identifier of the project which this container application belongs to.
	ContainerProjectId string `json:"container_project_id,omitempty"`
	// List of network errors related to container application.
	NetworkErrors []NetworkError `json:"network_errors,omitempty"`
}
