package nsxt

// Port Connection Hypervisor/Transport Node Entity
type PortConnectionHypervisor struct {
	Resource *ManagedResource `json:"resource,omitempty"`
	// Resource ID is mapped to this. (ID is Generated for Edge node groups, since resource will be null)
	Id string `json:"id,omitempty"`
	Pnics []Pnic `json:"pnics,omitempty"`
	PnicsList []NodeInterfaceProperties `json:"pnics_list,omitempty"`
	NeighborsList []InterfaceNeighborProperties `json:"neighbors_list,omitempty"`
	Profiles []BaseHostSwitchProfile `json:"profiles,omitempty"`
}
