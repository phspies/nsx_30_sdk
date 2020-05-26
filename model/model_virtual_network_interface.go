package nsxt

type VirtualNetworkInterface struct {
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
	// MAC address of the virtual network interface.
	MacAddress string `json:"mac_address"`
	// Owner virtual machine type; Edge, Service VM or other.
	OwnerVmType string `json:"owner_vm_type,omitempty"`
	// Device key of the virtual network interface.
	DeviceKey string `json:"device_key"`
	// Id of the host on which the vm exists.
	HostId string `json:"host_id"`
	// Id of the vm to which this virtual network interface belongs.
	OwnerVmId string `json:"owner_vm_id"`
	// Id of the vm unique within the host.
	VmLocalIdOnHost string `json:"vm_local_id_on_host"`
	// External Id of the virtual network inferface.
	ExternalId string `json:"external_id"`
	// LPort Attachment Id of the virtual network interface.
	LportAttachmentId string `json:"lport_attachment_id,omitempty"`
	// IP Addresses of the the virtual network interface, from various sources.
	IpAddressInfo []IpAddressInfo `json:"ip_address_info,omitempty"`
	// Device name of the virtual network interface.
	DeviceName string `json:"device_name,omitempty"`
}
