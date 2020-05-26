package nsxt

// Virtual portgroup on a virtual switch
type VirtualPortgroup struct {
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
	// External id of the virtual portgroup
	ExternalId string `json:"external_id,omitempty"`
	// Portgroup type like DistributedVirtualPortgroup
	OriginType string `json:"origin_type,omitempty"`
	// Id of the portgroup, eg. a mo-ref from VC.
	CmLocalId string `json:"cm_local_id,omitempty"`
}
