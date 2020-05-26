package nsxt

// Defines the tunnel local and peer addresses along with the multiple tansport tunnels for redundancy. L2VpnSession belongs on to only one L2VpnService.
type L2VpnSession struct {
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
	// L2VPN service id
	L2vpnServiceId string `json:"l2vpn_service_id"`
	TunnelEncapsulation *L2VpnTunnelEncapsulation `json:"tunnel_encapsulation,omitempty"`
	// Enable to extend all the associated logical switches.
	Enabled bool `json:"enabled,omitempty"`
	// List of transport tunnels for redundancy.
	TransportTunnels []ResourceReference `json:"transport_tunnels"`
}