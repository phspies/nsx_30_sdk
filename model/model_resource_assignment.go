package nsxt

// Amount of memory and CPU allocated to the Edge VM. 
type ResourceAssignment struct {
	// Memory allocation in MB. 
	MemoryAllocationInMb int32 `json:"memory_allocation_in_mb,omitempty"`
	// CPU count. 
	CpuCount int32 `json:"cpu_count,omitempty"`
}
