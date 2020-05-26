package nsxt

// Advertisement config for different types of routes which need to be advertised from TIER1 logical router to the linked TIER0 logical router 
type AdvertisementConfig struct {
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
	// Flag to advertise all connected routes
	AdvertiseNsxConnectedRoutes bool `json:"advertise_nsx_connected_routes,omitempty"`
	// Flag to advertise lb vip ips
	AdvertiseLbVip bool `json:"advertise_lb_vip,omitempty"`
	// Flag to advertise all static routes
	AdvertiseStaticRoutes bool `json:"advertise_static_routes,omitempty"`
	// TIER1 logical router id on which to enable this configuration
	LogicalRouterId string `json:"logical_router_id,omitempty"`
	// Flag to advertise all routes of dns forwarder listener ips and source ips
	AdvertiseDnsForwarder bool `json:"advertise_dns_forwarder,omitempty"`
	// Flag to advertise all routes of nat
	AdvertiseNatRoutes bool `json:"advertise_nat_routes,omitempty"`
	// Flag to advertise all IPSec VPN local endpoint ips to linked TIER0 logical router
	AdvertiseIpsecLocalIp bool `json:"advertise_ipsec_local_ip,omitempty"`
	// Flag to enable this configuration
	Enabled bool `json:"enabled,omitempty"`
	// Flag to advertise all lb SNAT ips
	AdvertiseLbSnatIp bool `json:"advertise_lb_snat_ip,omitempty"`
}
