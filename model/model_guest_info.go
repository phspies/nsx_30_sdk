package nsxt

// Guest virtual machine details include OS name and computer name of guest VM. 
type GuestInfo struct {
	// OS name of guest virtual machine. Currently this is supported for guests on ESXi that have VMware Tools installed. 
	OsName string `json:"os_name,omitempty"`
	// Computer name of guest virtual machine, which is set inside guest OS. Currently this is supported for guests on ESXi that have VMware Tools installed. 
	ComputerName string `json:"computer_name,omitempty"`
}
