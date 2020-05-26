package nsxt

// pNIC/bond statuses
type PnicBondStatus struct {
	// Status of pNIC/bond
	Status string `json:"status,omitempty"`
	// type, whether the object is a pNIC or a bond
	Type_ string `json:"type,omitempty"`
	// Name of the pNIC/bond
	Name string `json:"name,omitempty"`
}
