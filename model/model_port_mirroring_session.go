package nsxt

type PortMirroringSession struct {
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
	// Port mirroring session direction
	Direction string `json:"direction"`
	// Mirror sources
	MirrorSources []MirrorSource `json:"mirror_sources"`
	// Only for Remote SPAN Port Mirror.
	EncapsulationVlanId int64 `json:"encapsulation_vlan_id,omitempty"`
	// If this property is unset, this session will be treated as LocalPortMirrorSession. 
	SessionType string `json:"session_type,omitempty"`
	// If this property is set, the packet will be truncated to the provided length. If this property is unset, entire packet will be mirrored. 
	SnapLength int64 `json:"snap_length,omitempty"`
	// An array of 5-tuples used to filter packets for the mirror session, if not provided, all the packets will be mirrored.
	PortMirroringFilters []PortMirroringFilter `json:"port_mirroring_filters,omitempty"`
	// Only for Remote SPAN Port Mirror. Whether to preserve original VLAN.
	PreserveOriginalVlan bool `json:"preserve_original_vlan,omitempty"`
	MirrorDestination *MirrorDestination `json:"mirror_destination"`
}