package nsxt

// Base type of IPv6 ip-allocation extended by ip-pool and static-binding. 
type DhcpV6InfoBase struct {
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
	// Lease time, in seconds.
	LeaseTime int64 `json:"lease_time,omitempty"`
	// SNTP server ips.
	SntpServers []string `json:"sntp_servers,omitempty"`
	// Preferred time, in seconds. If this value is not provided, the value of lease_time*0.8 will be used. 
	PreferredTime int64 `json:"preferred_time,omitempty"`
	// Primary and secondary DNS server address to assign host. They can be overridden by ip-pool or static-binding level property. 
	DnsNameservers []string `json:"dns_nameservers,omitempty"`
	// Host name or prefix to be assigned to host. It can be overridden by ip-pool or static-binding level property. 
	DomainNames []string `json:"domain_names,omitempty"`
}
