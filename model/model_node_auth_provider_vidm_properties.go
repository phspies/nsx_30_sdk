package nsxt

// Node AAA provider vIDM properties
type NodeAuthProviderVidmProperties struct {
	// vIDM client secret
	ClientSecret string `json:"client_secret,omitempty"`
	// Fully Qualified Domain Name(FQDN) of vIDM
	HostName string `json:"host_name"`
	// vIDM client id
	ClientId string `json:"client_id"`
	// vIDM enable flag
	VidmEnable bool `json:"vidm_enable,omitempty"`
	// Load Balancer enable flag
	LbEnable bool `json:"lb_enable,omitempty"`
	// Hexadecimal SHA256 hash of the vIDM server's X.509 certificate
	Thumbprint string `json:"thumbprint"`
	// host name to use when creating the redirect URL for clients to follow after authenticating to vIDM
	NodeHostName string `json:"node_host_name"`
}
