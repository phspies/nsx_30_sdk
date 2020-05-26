package nsxt

type VmToolsInfo struct {
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
	// Type of VM - Edge, Service or other.
	VmType string `json:"vm_type,omitempty"`
	// Version of network agent on the VM of a third party partner solution.
	NetworkAgentVersion string `json:"network_agent_version,omitempty"`
	// Id of the VM which is assigned locally by the host. It is the VM-moref on ESXi hosts, in other environments it is VM UUID.
	HostLocalId string `json:"host_local_id,omitempty"`
	// Current external id of this virtual machine in the system.
	ExternalId string `json:"external_id,omitempty"`
	// Version of VMTools installed on the VM.
	ToolsVersion string `json:"tools_version,omitempty"`
	// Version of file agent on the VM of a third party partner solution.
	FileAgentVersion string `json:"file_agent_version,omitempty"`
}
