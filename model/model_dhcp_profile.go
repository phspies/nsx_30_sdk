package nsxt

// DHCP profile to specify edge cluster and members on which the dhcp server will run. A DhcpProfile can be referenced by different logical DHCP servers. 
type DhcpProfile struct {
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
	// The Edge nodes on which the DHCP servers run. If none is provided, the NSX will auto-select two edge-nodes from the given edge cluster. If only one edge node is provided, the DHCP servers will run without HA support. 
	EdgeClusterMemberIndexes []int64 `json:"edge_cluster_member_indexes,omitempty"`
	// Flag to enable the auto-relocation of standby DHCP Service in case of edge node failure. Only tier 1 and auto placed DHCP servers are considered for the relocation. 
	EnableStandbyRelocation bool `json:"enable_standby_relocation,omitempty"`
	// Edge cluster uuid on which the referencing logical DHCP server runs. 
	EdgeClusterId string `json:"edge_cluster_id"`
}
