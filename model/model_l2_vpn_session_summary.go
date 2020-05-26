package nsxt

// List summary of L2VPN sessions.
type L2VpnSessionSummary struct {
	// Total L2VPN sessions configured.
	TotalL2vpnSessions int64 `json:"total_l2vpn_sessions,omitempty"`
	// Number of established L2VPN sessions. L2VPN session is established when all the tunnels are up.
	EstablishedL2vpnSessions int64 `json:"established_l2vpn_sessions,omitempty"`
	// Number of failed L2VPN sessions. L2VPN session is failed when all the tunnels are down.
	FailedL2vpnSessions int64 `json:"failed_l2vpn_sessions,omitempty"`
}
