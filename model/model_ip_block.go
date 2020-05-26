package nsxt

// A block of IPv4/IPv6 addresses defined by a start address and a mask/prefix (network CIDR). An IP block is typically large & allocated to a tenant for automated consumption. An IP block is always a contiguous address space, for example 192.0.0.1/8. An IP block can be further subdivided into subnets called IP block subnets. These IP block subnets can later be added to IP pools and used for IP allocation. An IP pool is typically a collection of subnets that are often not a contiguous address space. Clients are allocated IP addresses only from IP pools. Sample Structure Diagram IpBlock_VMware 192.0.0.1/8 ======================================================================= /                          ___________________________________________/________ /  IpBlockSubnet_Finance  ( IpBlockSubnet_Eng1   IpBlockSubnet_Eng2   /        ) /  192.168.0.1/16         (   192.170.1.1/16      192.180.1.1/24      /        ) IpPool_Eng /                         (___________________________________________/________) /                                                                     / =======================================================================
type IpBlock struct {
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
	// Represents network address and the prefix length which will be associated with a layer-2 broadcast domain
	Cidr string `json:"cidr"`
}
