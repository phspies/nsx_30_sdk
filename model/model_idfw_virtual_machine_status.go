package nsxt

// ID and status of the VM on Identity Firewall compute collection. 
type IdfwVirtualMachineStatus struct {
	// Status of the Identity Firewall compute collection's Virtual Machine. 
	VmStatus []IdfwVirtualMachineCondition `json:"vm_status"`
	// VM ID of the VM on Identity Firewall compute collection. 
	VmId string `json:"vm_id"`
}
