package nsxt

// List of VMs on which a particular signature was detected with the count. 
type IdsVmStats struct {
	// Number of unique VMs on which a particular signature was detected.
	Count int64 `json:"count,omitempty"`
	// List of VM names  on which intrusions of that particular signature type were detected.
	VmList []string `json:"vm_list,omitempty"`
}
