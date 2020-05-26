package nsxt

// The credentials to login into the host node
type HostNodeLoginCredential struct {
	// The username of the account on the host node
	Username string `json:"username,omitempty"`
	// The authentication password of the host node
	Password string `json:"password,omitempty"`
	// For ESXi hosts, the thumbprint of the ESXi management service. For KVM hosts, the SSH key fingerprint. If thumbprint is not provided then connection to host may not be established and API call will fail. 
	Thumbprint string `json:"thumbprint,omitempty"`
}
