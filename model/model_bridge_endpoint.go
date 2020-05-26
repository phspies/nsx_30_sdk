package nsxt

// A bridge endpoint can be created on a bridge cluster or on an edge cluster. Few of the properties of this class will not be used depending on the type of bridge endpoint being created. When creating a bridge endpoint on a bridge cluster, following propeties will be used: vlan, bridge_cluster_id and ha_enable. Similarly, for creating a bridge endpoint on an edge cluster following properties will be used: vlan, bridge_endpoint_profile_id and vlan_transport_zone_id. 
type BridgeEndpoint struct {
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
	// This field will not be used if an edge cluster is being used for the bridge endpoint 
	HaEnable bool `json:"ha_enable,omitempty"`
	// This field will not be used if an edge cluster is being used for the bridge endpoint 
	BridgeClusterId string `json:"bridge_cluster_id,omitempty"`
	// This field will not be used if a bridge cluster is being used for the bridge endpoint 
	VlanTransportZoneId string `json:"vlan_transport_zone_id,omitempty"`
	// This field will not be used if a bridge cluster is being used for the bridge endpoint 
	BridgeEndpointProfileId string `json:"bridge_endpoint_profile_id,omitempty"`
	// This name has to be one of the switching uplink teaming policy names listed inside the TransportZone. If this field is not specified, bridge will use the first pnic in host-switch config. This field will not be used if a bridge cluster is being used for the bridge endpoint
	UplinkTeamingPolicyName string `json:"uplink_teaming_policy_name,omitempty"`
	VlanTrunkSpec *VlanTrunkSpec `json:"vlan_trunk_spec,omitempty"`
	// This property is used for VLAN specification of bridge endpoint. It's mutually exclusive with 'vlan_trunk_spec', either 'vlan' or 'vlan_trunk_spec' should be specified. 
	Vlan int64 `json:"vlan,omitempty"`
}