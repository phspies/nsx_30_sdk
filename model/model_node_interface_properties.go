package nsxt

// Node network interface properties
type NodeInterfaceProperties struct {
	// Indicates whether interface is managed by the host
	HostManaged bool `json:"host_managed,omitempty"`
	// Type of switch associated with the interface.
	ConnectedSwitchType string `json:"connected_switch_type,omitempty"`
	// Interface administration status
	LinkStatus string `json:"link_status,omitempty"`
	// Indicates whether interface is enabled for Enhanced Networking Stack
	EnsEnabled bool `json:"ens_enabled,omitempty"`
	// Interface Type
	InterfaceType string `json:"interface_type,omitempty"`
	// Connected switch
	ConnectedSwitch string `json:"connected_switch,omitempty"`
	// Interface ID
	InterfaceId string `json:"interface_id,omitempty"`
	// Source of status data
	Source string `json:"source,omitempty"`
	// Interface administration status
	AdminStatus string `json:"admin_status,omitempty"`
	// PCI device.
	Pci string `json:"pci,omitempty"`
	// Indicates whether backing of VIRTUAL network interface is managed by NSX
	BackingNsxManaged bool `json:"backing_nsx_managed,omitempty"`
	// Device key.
	Key string `json:"key,omitempty"`
	// IP Alias
	InterfaceAlias []NodeInterfaceAlias `json:"interface_alias,omitempty"`
	// Device name.
	Device string `json:"device,omitempty"`
	// LPort Attachment Id assigned to VIRTUAL network interface of a node
	LportAttachmentId string `json:"lport_attachment_id,omitempty"`
	// Interface MTU
	Mtu int64 `json:"mtu,omitempty"`
	// Driver name.
	Driver string `json:"driver,omitempty"`
	// This boolean property describes if network interface is capable for Enhanced Networking Stack interrupt
	EnsInterruptCapable bool `json:"ens_interrupt_capable,omitempty"`
	// Interface capability for Enhanced Networking Stack
	EnsCapable bool `json:"ens_capable,omitempty"`
	// This boolean property describes if network interface is enabled for Enhanced Networking Stack interrupt
	EnsInterruptEnabled bool `json:"ens_interrupt_enabled,omitempty"`
}
