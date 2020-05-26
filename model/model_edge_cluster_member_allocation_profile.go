package nsxt

type EdgeClusterMemberAllocationProfile struct {
	AllocationPool *EdgeClusterMemberAllocationPool `json:"allocation_pool,omitempty"`
	// Flag to enable the auto-relocation of standby service router running on edge cluster and node associated with the logical router. Only dynamically allocated tier1 logical routers are considered for the relocation. 
	EnableStandbyRelocation bool `json:"enable_standby_relocation,omitempty"`
	// Allocation type is used to specify the mode used to allocate the LR. This is populated only for TIER1 logical router and for TIER0 this will be null. 
	AllocationType string `json:"allocation_type,omitempty"`
}
