package nsxt

type LogicalRouter struct {
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
	// For stateful services, the logical router should be associated with edge cluster. For TIER 1 logical router, for manual placement of service router within the cluster, edge cluster member indices needs to be provided else same will be auto-allocated. You can provide maximum two indices for HA ACTIVE_STANDBY. For TIER0 logical router this property is no use and placement is derived from logical router uplink or loopback port. 
	EdgeClusterMemberIndices []int64 `json:"edge_cluster_member_indices,omitempty"`
	Ipv6Profiles *IPv6Profiles `json:"ipv6_profiles,omitempty"`
	AllocationProfile *EdgeClusterMemberAllocationProfile `json:"allocation_profile,omitempty"`
	// List of Firewall sections related to Logical Router.
	FirewallSections []ResourceReference `json:"firewall_sections,omitempty"`
	// Determines the behavior when a logical router instance restarts after a failure. If set to PREEMPTIVE, the preferred node will take over, even if it causes another failure. If set to NON_PREEMPTIVE, then the instance that restarted will remain secondary. This property must not be populated unless the high_availability_mode property is set to ACTIVE_STANDBY. If high_availability_mode property is set to ACTIVE_STANDBY and this property is not specified then default will be NON_PREEMPTIVE. 
	FailoverMode string `json:"failover_mode,omitempty"`
	AdvancedConfig *LogicalRouterConfig `json:"advanced_config,omitempty"`
	// TIER0 for external connectivity. TIER1 for two tier topology with TIER0 on top. VRF for isolation of routing table on TIER0. 
	RouterType string `json:"router_type"`
	// Preferred edge cluster member index which is required for PREEMPTIVE failover mode. Used for Tier0 routers only. 
	PreferredEdgeClusterMemberIndex int64 `json:"preferred_edge_cluster_member_index,omitempty"`
	// High availability mode
	HighAvailabilityMode string `json:"high_availability_mode,omitempty"`
	// Used for tier0 routers
	EdgeClusterId string `json:"edge_cluster_id,omitempty"`
}
