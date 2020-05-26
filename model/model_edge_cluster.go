package nsxt

type EdgeCluster struct {
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
	// Edge cluster is homogenous collection of transport nodes. Hence all transport nodes of the cluster must be of same type. This readonly field shows the type of transport nodes. 
	MemberNodeType string `json:"member_node_type,omitempty"`
	// EdgeCluster only supports homogeneous members. These member should be backed by either EdgeNode or PublicCloudGatewayNode. TransportNode type of these nodes should be the same. DeploymentType (VIRTUAL_MACHINE|PHYSICAL_MACHINE) of these EdgeNodes is recommended to be the same. EdgeCluster supports members of different deployment types. 
	Members []EdgeClusterMember `json:"members,omitempty"`
	// List of remote tunnel endpoint ipaddress configured on edge cluster for each transport node.
	NodeRtepIps []NodeRtepIpsConfig `json:"node_rtep_ips,omitempty"`
	// Edge cluster profile bindings
	ClusterProfileBindings []ClusterProfileTypeIdEntry `json:"cluster_profile_bindings,omitempty"`
	// Flag should be only use in federation for inter site l2 and l3 forwarding. Before enabling this flag, all the edge cluster members must have remote tunnel endpoint configured. TIER0/TIER1 logical routers managed by GM must be associated with edge cluster which has inter-site forwarding enabled. 
	EnableInterSiteForwarding bool `json:"enable_inter_site_forwarding,omitempty"`
	// Set of allocation rules and respected action for auto placement of logical router, DHCP and MDProxy on edge cluster members. 
	AllocationRules []AllocationRule `json:"allocation_rules,omitempty"`
	// This field is a readonly field which shows the deployment_type of members. It returns UNKNOWN if there are no members, and returns VIRTUAL_MACHINE| PHYSICAL_MACHINE if all edge members are VIRTUAL_MACHINE|PHYSICAL_MACHINE. It returns HYBRID if the cluster contains edge members of both types VIRTUAL_MACHINE and PHYSICAL_MACHINE. 
	DeploymentType string `json:"deployment_type,omitempty"`
}
