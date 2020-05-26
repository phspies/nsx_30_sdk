package nsxt

// Tunnel endpoint configuration of preconfigured host switch
type PreconfiguredEndpoint struct {
	// Name of the virtual tunnel endpoint
	DeviceName string `json:"device_name"`
}
