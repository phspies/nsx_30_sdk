package nsxt

// TCP profile allows customization of TCP stack behavior for each application. As TCP connections are terminated (or initiated) on the load balancer only for L7 virtual servers, TCP profiles are applicable only to them. As the desired TCP characteristics (e.g. Rx/Tx buffers) may be different for client (typically on WAN) and server (typically on LAN) sides, two separate profiles can be bound to virtual server, one for client-side (LbVirtualServer.client_tcp_profile_id) and another for server-side (LbVirtualServer.server_tcp_profile_id). 
type LbTcpProfile struct {
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
	// Setting this option to enable or disable Nagle's algorithm.
	NagleAlgorithmEnabled bool `json:"nagle_algorithm_enabled,omitempty"`
	// If the field is not specified, the load balancer will use the default setting per load balancer service flavor. 
	ReceiveWindowSize int64 `json:"receive_window_size,omitempty"`
	// If the field is not specified, the load balancer will use the default setting per load balancer service flavor. 
	FinWait2Timeout int64 `json:"fin_wait2_timeout,omitempty"`
	// If the field is not specified, the load balancer will use the default setting per load balancer service flavor. 
	MaxSynRetransmissions int64 `json:"max_syn_retransmissions,omitempty"`
	// If the field is not specified, the load balancer will use the default setting per load balancer service flavor. 
	TransmitWindowSize int64 `json:"transmit_window_size,omitempty"`
}
