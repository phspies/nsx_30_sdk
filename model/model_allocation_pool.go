package nsxt

type AllocationPool struct {
	// Represents the number of standby services running on the edge node. 
	StandbyServiceCount int32 `json:"standby_service_count,omitempty"`
	// Represents the number of acitve services running on the edge node. 
	ActiveServiceCount int32 `json:"active_service_count,omitempty"`
	// Allocation details of sub-pools configured on edge node.
	SubPools []SubPool `json:"sub_pools,omitempty"`
}