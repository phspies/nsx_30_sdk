package nsxt

// Cluster virtual IP properties
type ClusterVirtualIpProperties struct {
	// Virtual IP address, 0.0.0.0 if not configured
	IpAddress string `json:"ip_address"`
}
