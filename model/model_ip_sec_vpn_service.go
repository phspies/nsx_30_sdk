package nsxt

// Create and manage IPSec VPN service for given logical router.
type IpSecVpnService struct {
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
	// Log level for internet key exchange (IKE).
	IkeLogLevel string `json:"ike_log_level,omitempty"`
	// Logical router id.
	LogicalRouterId string `json:"logical_router_id"`
	// Enable/disable IPSec HA state sync. IPSec HA state sync can be disabled in case there are performance issues with the state sync messages. Default is to enable HA Sync. 
	IpsecHaSync bool `json:"ipsec_ha_sync,omitempty"`
	// Bypass policy rules are configured using VPN service. Bypass rules always have higher priority over protect rules and they affect all policy based vpn sessions associated with the IPSec VPN service. Protect rules are defined per policy based vpn session. 
	BypassRules []IpSecVpnPolicyRule `json:"bypass_rules,omitempty"`
	// If true, enable VPN services for given logical router.
	Enabled bool `json:"enabled,omitempty"`
}
