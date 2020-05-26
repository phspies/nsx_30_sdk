package nsxt

// Details of Container Ingress Policy.
type ContainerIngressPolicy struct {
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
	// Network status of container ingress.
	NetworkStatus string `json:"network_status,omitempty"`
	// Identifier of the container cluster this ingress policy belongs to.
	ContainerClusterId string `json:"container_cluster_id,omitempty"`
	// List of identifiers of the container application , on which ingress policy is applied. e.g. IDs of all services on which the ingress is applied in kubernetes. 
	ContainerApplicationIds []string `json:"container_application_ids,omitempty"`
	// Array of additional specific properties of container ingress in key-value format. 
	OriginProperties []KeyValuePair `json:"origin_properties,omitempty"`
	// Identifier of the container ingress policy.
	ExternalId string `json:"external_id"`
	// Identifier of the project which this container ingress belongs to.
	ContainerProjectId string `json:"container_project_id,omitempty"`
	// List of network errors related to container ingress.
	NetworkErrors []NetworkError `json:"network_errors,omitempty"`
	// Container ingress policy specification.
	Spec string `json:"spec,omitempty"`
}
