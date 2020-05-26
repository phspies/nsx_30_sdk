package nsxt

// Physical NIC specification
type Pnic struct {
	// Uplink name for this Pnic. This name will be used to reference this Pnic in other configurations.
	UplinkName string `json:"uplink_name"`
	// device name or key
	DeviceName string `json:"device_name"`
}
