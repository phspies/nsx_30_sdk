package nsxt

type LogicalSwitch struct {
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
	// This field indicates purpose of a LogicalSwitch. It is set by manager internally or user can provide this field. If not set, DEFAULT type is assigned. NSX components can use this field to create LogicalSwitch that provides component specific functionality. DEFAULT type LogicalSwitches are created for basic L2 connectivity by API users. SERVICE_PLANE type LogicalSwitches are system created service plane LogicalSwitches for Service Insertion service. User can not create SERVICE_PLANE type of LogicalSwitch. DHCP_RELAY type LogicalSwitches are created by external user like Policy with special permissions or by system and will be treated as internal LogicalSwitches. Such LogicalSwitch will not be exposed to vSphere user. GLOBAL type LogicalSwitches are created to span multiple NSX domains to connect multiple remote sites. 
	SwitchType string `json:"switch_type,omitempty"`
	// Replication mode of the Logical Switch
	ReplicationMode string `json:"replication_mode,omitempty"`
	// This property could be used for vendor specific configuration in key value string pairs, the setting in extra_configs will be automatically inheritted by logical ports in the logical switch. 
	ExtraConfigs []ExtraConfig `json:"extra_configs,omitempty"`
	// This name has to be one of the switching uplink teaming policy names listed inside the logical switch's TransportZone. If this field is not specified, the logical switch will not have a teaming policy associated with it and the host switch's default teaming policy will be used.
	UplinkTeamingPolicyName string `json:"uplink_teaming_policy_name,omitempty"`
	// Address bindings for the Logical switch
	AddressBindings []PacketAddressClassifier `json:"address_bindings,omitempty"`
	// IP pool id that associated with a LogicalSwitch.
	IpPoolId string `json:"ip_pool_id,omitempty"`
	// This property is dedicated to VLAN based network, to set VLAN of logical network. It is mutually exclusive with 'vlan_trunk_spec'. 
	Vlan int64 `json:"vlan,omitempty"`
	// If this flag is set to true, then all the logical switch ports attached to this logical switch will behave in a hybrid fashion. The hybrid logical switch port indicates to NSX that the VM intends to operate in underlay mode, but retains the ability to forward egress traffic to the NSX overlay network. This flag can be enabled only for the logical switches in the overlay type transport zone which has host switch mode as STANDARD and also has either CrossCloud or CloudScope tag scopes. Only the NSX public cloud gateway (PCG) uses this flag, other host agents like ESX, KVM and Edge will ignore it. This property cannot be modified once the logical switch is created. 
	Hybrid bool `json:"hybrid,omitempty"`
	// Mac pool id that associated with a LogicalSwitch.
	MacPoolId string `json:"mac_pool_id,omitempty"`
	// Only for OVERLAY network. A VNI will be auto-allocated from the default VNI pool if not given; otherwise the given VNI has to be inside the default pool and not used by any other LogicalSwitch. 
	Vni int32 `json:"vni,omitempty"`
	VlanTrunkSpec *VlanTrunkSpec `json:"vlan_trunk_spec,omitempty"`
	// Represents Desired state of the Logical Switch
	AdminState string `json:"admin_state"`
	// Id of the TransportZone to which this LogicalSwitch is associated
	TransportZoneId string `json:"transport_zone_id"`
	// Each manager ID represents the NSX Local Manager the logical switch connects. This will be populated by the manager.
	Span []string `json:"span,omitempty"`
	SwitchingProfileIds []SwitchingProfileTypeIdEntry `json:"switching_profile_ids,omitempty"`
	// The VNI is used for intersite traffic and the global logical switch ID. The global VNI pool is agnostic of the local VNI pool, and there is no need to have an exclusive VNI range. For example, VNI x can be the global VNI for logical switch B and the local VNI for logical switch A.
	GlobalVni int32 `json:"global_vni,omitempty"`
}
