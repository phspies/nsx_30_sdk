package nsxt

type LogicalPort struct {
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
	// Id of the Logical switch that this port belongs to.
	LogicalSwitchId string `json:"logical_switch_id"`
	// Set initial state when a new logical port is created. 'UNBLOCKED_VLAN' means new port will be unblocked on traffic in creation, also VLAN will be set with corresponding logical switch setting. This port setting can only be configured at port creation (POST), and cannot be modified. 
	InitState string `json:"init_state,omitempty"`
	SwitchingProfileIds []SwitchingProfileTypeIdEntry `json:"switching_profile_ids,omitempty"`
	Attachment *LogicalPortAttachment `json:"attachment,omitempty"`
	// The internal_id of the logical port may or may not be identical to it's managed resource ID. If a VirtualMachine connected to logical port migrates from one site to another, then on the destination site, it will be connected to different logical port managed resource. However, the internal_id field will be persisted across vmotion. 
	InternalId string `json:"internal_id,omitempty"`
	// This property could be used for vendor specific configuration in key value string pairs. Logical port setting will override logical switch setting if the same key was set on both logical switch and logical port. 
	ExtraConfigs []ExtraConfig `json:"extra_configs,omitempty"`
	// Each address binding must contain both an IPElement and MAC address. VLAN ID is optional. This binding configuration can be used by features such as spoof-guard and overrides any discovered bindings. Any non unique entries are deduplicated to generate a unique set of address bindings and then stored. For IPv6 addresses, a subnet address cannot have host bits set. A maximum of 128 unique address bindings is allowed per port. 
	AddressBindings []PacketAddressClassifier `json:"address_bindings,omitempty"`
	// IP Discovery module uses various mechanisms to discover address bindings being used on each port. If a user would like to ignore any specific discovered address bindings or prevent the discovery of a particular set of discovered bindings, then those address bindings can be provided here. Currently IP range in CIDR format is not supported. 
	IgnoreAddressBindings []PacketAddressClassifier `json:"ignore_address_bindings,omitempty"`
	// Represents Desired state of the logical port
	AdminState string `json:"admin_state"`
}
