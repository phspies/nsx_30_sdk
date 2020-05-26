package nsxt

// Virtual switch on a compute manager
type VirtualSwitch struct {
	// Timestamp of last modification
	LastSyncTime int64 `json:"_last_sync_time,omitempty"`
	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`
	// Description of this resource
	Description string `json:"description,omitempty"`
	// The type of this resource.
	ResourceType string `json:"resource_type"`
	// Opaque identifiers meaningful to the API user
	Tags []Tag `json:"tags,omitempty"`
	// ID of the virtual switch in compute manager
	CmLocalId string `json:"cm_local_id,omitempty"`
	// External id of the virtual switch
	ExternalId string `json:"external_id,omitempty"`
	// Switch type like VmwareDistributedVirtualSwitch
	OriginType string `json:"origin_type,omitempty"`
	// ID of the compute manager where this virtual switch is discovered. 
	OriginId string `json:"origin_id,omitempty"`
}
