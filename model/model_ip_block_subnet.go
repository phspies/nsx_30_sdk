package nsxt

// A set of IPv4/IPv6 addresses defined by a start address and a mask/prefix which will typically be associated with a layer-2 broadcast domain.
type IpBlockSubnet struct {
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
	// For internal system use Only. Represents start ip address of the subnet from IP block. Subnet ip adddress will start from this ip address.
	StartIp string `json:"start_ip,omitempty"`
	// Represents network address and the prefix length which will be associated with a layer-2 broadcast domain
	Cidr string `json:"cidr,omitempty"`
	// A collection of IPv4/IPv6 IP ranges used for IP allocation.
	AllocationRanges []IpPoolRange `json:"allocation_ranges,omitempty"`
	// Block id for which the subnet is created.
	BlockId string `json:"block_id"`
	// Represents the size or number of ip addresses in the subnet
	Size int64 `json:"size"`
}