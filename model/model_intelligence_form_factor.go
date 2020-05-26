package nsxt

// NSX Intelligence node form factor
type IntelligenceFormFactor struct {
	// Number of virtual cpus on the Intelligence nodes
	Vcpu int64 `json:"vcpu,omitempty"`
	// Disk size of the Intelligence nodes in GBs.
	Disk int64 `json:"disk,omitempty"`
	// NSX Intelligence node form factor type 
	Type_ string `json:"type,omitempty"`
	// Memory size of the Intelligence nodes in GBs
	Memory int64 `json:"memory,omitempty"`
}
