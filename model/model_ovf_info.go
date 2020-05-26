package nsxt

// Information related to OVF file.
type OvfInfo struct {
	// Version of the OVF.
	Version string `json:"version,omitempty"`
	// Name of OVF file.
	OvfName string `json:"ovf_name,omitempty"`
}
