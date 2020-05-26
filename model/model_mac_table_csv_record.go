package nsxt

type MacTableCsvRecord struct {
	// The virtual tunnel endpoint MAC address
	VtepMacAddress string `json:"vtep_mac_address,omitempty"`
	// The virtual tunnel endpoint IP address
	VtepIp string `json:"vtep_ip,omitempty"`
	// The MAC address
	MacAddress string `json:"mac_address"`
}
