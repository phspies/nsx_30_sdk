package nsxt

// Status of the Identity Firewall compute collection's VM.
type IdfwVirtualMachineCondition struct {
	// VM IDFW Status.
	Status string `json:"status"`
	// IDFW compute collection's VM condition.
	StatusDetail string `json:"status_detail,omitempty"`
}
