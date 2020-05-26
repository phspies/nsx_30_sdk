package nsxt

// Object to identify an uplink based on its type and name
type Uplink struct {
	// Name of this uplink
	UplinkName string `json:"uplink_name"`
	// Type of the uplink
	UplinkType string `json:"uplink_type"`
}
