package nsxt

type VirtualMachine struct {
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
	Source *ResourceReference `json:"source,omitempty"`
	// Id of the vm unique within the host.
	LocalIdOnHost string `json:"local_id_on_host"`
	// Virtual Machine type; Edge, Service VM or other.
	Type_ string `json:"type,omitempty"`
	GuestInfo *GuestInfo `json:"guest_info,omitempty"`
	// Current power state of this virtual machine in the system.
	PowerState string `json:"power_state"`
	// List of external compute ids of the virtual machine in the format 'id-type-key:value' , list of external compute ids ['uuid:xxxx-xxxx-xxxx-xxxx', 'moIdOnHost:moref-11', 'instanceUuid:xxxx-xxxx-xxxx-xxxx']
	ComputeIds []string `json:"compute_ids"`
	// Id of the host in which this virtual machine exists.
	HostId string `json:"host_id,omitempty"`
	// Current external id of this virtual machine in the system.
	ExternalId string `json:"external_id"`
}
