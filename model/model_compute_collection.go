package nsxt

type ComputeCollection struct {
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
	// Id of the compute manager from where this Compute Collection was discovered
	OriginId string `json:"origin_id,omitempty"`
	// Key-Value map of additional specific properties of compute collection in the Compute Manager 
	OriginProperties []KeyValuePair `json:"origin_properties,omitempty"`
	// External ID of the ComputeCollection in the source Compute manager, e.g. mo-ref in VC 
	ExternalId string `json:"external_id,omitempty"`
	// Id of the owner of compute collection in the Compute Manager
	OwnerId string `json:"owner_id,omitempty"`
	// ComputeCollection type like VC_Cluster. Here the Compute Manager type prefix would help in differentiating similar named Compute Collection types from different Compute Managers 
	OriginType string `json:"origin_type,omitempty"`
	// Local Id of the compute collection in the Compute Manager
	CmLocalId string `json:"cm_local_id,omitempty"`
}
