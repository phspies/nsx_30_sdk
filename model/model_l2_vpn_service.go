package nsxt

// Defines if service running as server or client Also defines all the common properties for the multiple L2VpnSessions associated with this service.
type L2VpnService struct {
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
	// Full mesh topology auto disables traffic replication between connected peers. However, this property is deprecated. Please refer enable_hub property instead to control client to client forwarding via the server. The value of enable_full_mesh will not be used anymore. If enable_hub is not provided explicitly, the default value of it will be used. 
	EnableFullMesh bool `json:"enable_full_mesh,omitempty"`
	// This property only applies in SERVER mode. If set to true, traffic from any client will be replicated to all other clients. If set to false, traffic received from clients is only replicated to the local VPN endpoint. 
	EnableHub bool `json:"enable_hub,omitempty"`
	// Logical router id
	LogicalRouterId string `json:"logical_router_id"`
	// Specify an L2VPN service mode as SERVER or CLIENT. L2VPN service in SERVER mode requires user to configure L2VPN session explicitly. L2VPN service in CLIENT mode can use peercode generated from SERVER to configure L2VPN session. 
	Mode string `json:"mode,omitempty"`
	// IP Pool to allocate local and peer endpoint IPs for L2VpnSession logical Tap.
	LogicalTapIpPool []string `json:"logical_tap_ip_pool,omitempty"`
}
