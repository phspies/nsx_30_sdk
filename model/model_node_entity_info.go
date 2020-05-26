package nsxt

type NodeEntityInfo struct {
	// IP address of service provider
	IpAddress string `json:"ip_address,omitempty"`
	// Port number of service provider
	Port int64 `json:"port,omitempty"`
	// Entity type of this service endpoint
	EntityType string `json:"entity_type,omitempty"`
}
