package nsxt

type Node struct {
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
	// Fabric node type, for example 'HostNode', 'EdgeNode' or 'PublicCloudGatewayNode'
	ResourceType string `json:"resource_type"`
	// Discovered IP Addresses of the fabric node, version 4 or 6
	DiscoveredIpAddresses []string `json:"discovered_ip_addresses,omitempty"`
	// IP Addresses of the Node, version 4 or 6. This property is mandatory for all nodes except for automatic deployment of edge virtual machine node. For automatic deployment, the ip address from management_port_subnets property will be considered. 
	IpAddresses []string `json:"ip_addresses,omitempty"`
	// ID of the Node maintained on the Node and used to recognize the Node
	ExternalId string `json:"external_id,omitempty"`
	// Fully qualified domain name of the fabric node
	Fqdn string `json:"fqdn,omitempty"`
}