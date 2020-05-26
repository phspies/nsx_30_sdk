package nsxt

type VtepTableCsvRecord struct {
	// The virtual tunnel endpoint label
	VtepLabel int64 `json:"vtep_label"`
	// The virtual tunnel endpoint MAC address
	VtepMacAddress string `json:"vtep_mac_address"`
	// The virtual tunnel endpoint IP address
	VtepIp string `json:"vtep_ip,omitempty"`
	// The segment Id
	SegmentId string `json:"segment_id,omitempty"`
}
