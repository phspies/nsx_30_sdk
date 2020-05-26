package nsxt

type BgpNeighbor struct {
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
	// BGP Graceful Restart mode. If specified, then it will take precedence over global Graceful Restart mode configured in logical router BgpConfig otherwise BgpConfig level Graceful Restart mode will be applicable for peer. 
	GracefulRestartMode string `json:"graceful_restart_mode,omitempty"`
	// This is a deprecated property, Please use 'remote_as_num' instead.
	RemoteAs int64 `json:"remote_as,omitempty"`
	// This is a deprecated property, Please use 'address_family' instead.
	FilterOutIpprefixlistId string `json:"filter_out_ipprefixlist_id,omitempty"`
	// Wait period (seconds) before declaring peer dead
	HoldDownTimer int64 `json:"hold_down_timer,omitempty"`
	// BGP neighborship will be formed from all these source addresses to this neighbour.
	SourceAddresses []string `json:"source_addresses,omitempty"`
	// This value is set on TTL(time to live) of BGP header. When router receives the BGP packet, it decrements the TTL. The default value of TTL is one when BPG request is initiated.So in the case of a BGP peer multiple hops away and and value of TTL is one, then  next router in the path will decrement the TTL to 0, realize it cant forward the packet and will drop it. If the hop count value to reach neighbor is equal to or less than the maximum_hop_limit value then intermediate router decrements the TTL count by one and forwards the request to BGP neighour. If the hop count value is greater than the maximum_hop_limit value then intermediate router discards the request when TTL becomes 0. 
	MaximumHopLimit int32 `json:"maximum_hop_limit,omitempty"`
	// Flag to enable this BGP Neighbor
	Enabled bool `json:"enabled,omitempty"`
	// 4 Byte ASN of the neighbor in ASPLAIN/ASDOT Format
	RemoteAsNum string `json:"remote_as_num,omitempty"`
	// User can enable the neighbor for the specific address families and also define filters per address family. When the neighbor is created, it is default enabled for IPV4_UNICAST address family for backward compatibility reasons. User can change that if required, by defining the address family configuration. 
	AddressFamilies []BgpNeighborAddressFamily `json:"address_families,omitempty"`
	BfdConfig *BfdConfigParameters `json:"bfd_config,omitempty"`
	// Logical router id
	LogicalRouterId string `json:"logical_router_id,omitempty"`
	// This is a deprecated property, Please  use 'address_family' instead.
	FilterInIpprefixlistId string `json:"filter_in_ipprefixlist_id,omitempty"`
	// This is a deprecated property, Please use 'address_family' instead.
	FilterOutRoutemapId string `json:"filter_out_routemap_id,omitempty"`
	// This is a deprecated property, Please use 'address_family' instead.
	FilterInRoutemapId string `json:"filter_in_routemap_id,omitempty"`
	// Frequency (seconds) with which keep alive messages are sent to peers
	KeepAliveTimer int64 `json:"keep_alive_timer,omitempty"`
	// User can create (POST) the neighbor with or without the password. The view (GET) on the neighbor, would never reveal if the password is set or not. The password can be set later using edit neighbor workFlow (PUT) On the edit neighbor (PUT), if the user does not specify the password property, the older value is retained. Maximum length of this field is 20 characters. 
	Password string `json:"password,omitempty"`
	// Deprecated - do not provide a value for this field. Use source_addresses instead.
	SourceAddress string `json:"source_address,omitempty"`
	// Flag to enable allowas_in option for BGP neighbor
	AllowAsIn bool `json:"allow_as_in,omitempty"`
	// Flag to enable BFD for this BGP Neighbor. Enable this if the neighbor supports BFD as this will lead to faster convergence.
	EnableBfd bool `json:"enable_bfd,omitempty"`
	// Neighbor IP Address
	NeighborAddress string `json:"neighbor_address"`
}
