package nsxt

// A collection of one or more IPv4 or IPv6 subnets or ranges that are often not a contiguous address space. Clients are allocated IPs from an IP pool. Often used when a client that consumes addresses exhausts an initial subnet or range and needs to be expanded but the adjoining address space is not available as it has been allocated to a different client.
type IpPool struct {
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
	// Subnets can be IPv4 or IPv6 and they should not overlap. The maximum number will not exceed 5 subnets.
	Subnets []IpPoolSubnet `json:"subnets,omitempty"`
	PoolUsage *PoolUsage `json:"pool_usage,omitempty"`
	// Delay in milliseconds, while releasing allocated IP address from IP pool (Default is 2 mins).
	IpReleaseDelay int64 `json:"ip_release_delay,omitempty"`
}
