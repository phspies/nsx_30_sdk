package nsxt

// Transport Node
type TransportNode struct {
	// Indicates system owned resource
	SystemOwned bool `json:"_system_owned,omitempty"`
	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`
	// Description of this resource
	Description string `json:"description,omitempty"`
	// Opaque identifiers meaningful to the API user
	Tags []Tag `json:"tags,omitempty"`
	// ID of the user who created this resource
	CreateUser string `json:"_create_user,omitempty"`
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed             to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed                 to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super                    user and can modify it, but only when providing                    the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this           entity. 
	Protection string `json:"_protection,omitempty"`
	// Timestamp of resource creation
	CreateTime int64 `json:"_create_time,omitempty"`
	// Timestamp of last modification
	LastModifiedTime int64 `json:"_last_modified_time,omitempty"`
	// ID of the user who last modified this resource
	LastModifiedUser string `json:"_last_modified_user,omitempty"`
	// Unique identifier of this resource
	Id string `json:"id,omitempty"`
	// The type of this resource.
	ResourceType string `json:"resource_type,omitempty"`
	HostSwitchSpec *HostSwitchSpec `json:"host_switch_spec,omitempty"`
	// Unique Id of the fabric node
	NodeId string `json:"node_id,omitempty"`
	NodeDeploymentInfo *Node `json:"node_deployment_info,omitempty"`
	// The property is read-only, used for querying result. User could update transport node maintenance mode by UpdateTransportNodeMaintenanceMode call.
	MaintenanceMode string `json:"maintenance_mode,omitempty"`
	// Set failure domain of edge transport node which will help in auto placement of TIER1 logical routers, DHCP Servers and MDProxies, if failure domain based allocation is enabled in edge cluster API. It is only supported for edge transport node and not for host transport node. In case failure domain is not set by user explicitly, it will be always assigned with default system created failure domain. 
	FailureDomainId string `json:"failure_domain_id,omitempty"`
	RemoteTunnelEndpoint *TransportNodeRemoteTunnelEndpointConfig `json:"remote_tunnel_endpoint,omitempty"`
	// This flag is relevant to only those hosts which are part of a compute collection which has transport node profile (TNP) applied on it. If you change the transport node configuration and it is different than cluster level TNP then this flag will be set to true 
	IsOverridden bool `json:"is_overridden,omitempty"`
	// This is deprecated. TransportZoneEndPoints should be specified per host switch at StandardHostSwitch through Transport Node or Transport Node Profile configuration. This will ONLY include the TransportZoneEndpoints that were were specified here during the TransportNode configuration. If TransportZoneEndpoints are specified directly in {$ref: StandardHostSwitch}, such TransportZoneEndpoints will not be populated here. 
	TransportZoneEndpoints []TransportZoneEndPoint `json:"transport_zone_endpoints,omitempty"`
}
