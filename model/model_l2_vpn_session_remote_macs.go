package nsxt

// Remote MAC addresses for all logical switches for a L2VPN session.
type L2VpnSessionRemoteMacs struct {
	// List MAC addresses for all logical switch for a particular L2VPN session.
	RemoteMacAddresses []L2VpnSessionRemoteMacsForLs `json:"remote_mac_addresses,omitempty"`
	// L2VPN display name.
	DisplayName string `json:"display_name,omitempty"`
	// L2VPN session identifier.
	SessionId string `json:"session_id,omitempty"`
}
