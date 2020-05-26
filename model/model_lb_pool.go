package nsxt

type LbPool struct {
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
	MemberGroup *PoolMemberGroup `json:"member_group,omitempty"`
	SnatTranslation *LbSnatTranslation `json:"snat_translation,omitempty"`
	// Load balancing algorithm, configurable per pool controls how the incoming connections are distributed among the members. 
	Algorithm string `json:"algorithm,omitempty"`
	// Server pool consists of one or more pool members. Each pool member is identified, typically, by an IP address and a port. 
	Members []PoolMember `json:"members,omitempty"`
	// Passive healthchecks are disabled by default and can be enabled by attaching a passive health monitor to a server pool. Each time a client connection to a pool member fails, its failed count is incremented. For pools bound to L7 virtual servers, a connection is considered to be failed and failed count is incremented if any TCP connection errors (e.g. TCP RST or failure to send data) or SSL handshake failures occur. For pools bound to L4 virtual servers, if no response is received to a TCP SYN sent to the pool member or if a TCP RST is received in response to a TCP SYN, then the pool member is considered to have failed and the failed count is incremented. 
	PassiveMonitorId string `json:"passive_monitor_id,omitempty"`
	// The maximum number of TCP connections per pool that are idly kept alive for sending future client requests. 
	TcpMultiplexingNumber int64 `json:"tcp_multiplexing_number,omitempty"`
	// In case of active healthchecks, load balancer itself initiates new connections (or sends ICMP ping) to the servers periodically to check their health, completely independent of any data traffic. Active healthchecks are disabled by default and can be enabled for a server pool by binding a health monitor to the pool. If multiple active monitors are configured, the pool member status is UP only when the health check status for all the monitors are UP. 
	ActiveMonitorIds []string `json:"active_monitor_ids,omitempty"`
	// TCP multiplexing allows the same TCP connection between load balancer and the backend server to be used for sending multiple client requests from different client TCP connections. 
	TcpMultiplexingEnabled bool `json:"tcp_multiplexing_enabled,omitempty"`
	// A pool is considered active if there are at least certain minimum number of members. 
	MinActiveMembers int64 `json:"min_active_members,omitempty"`
}
