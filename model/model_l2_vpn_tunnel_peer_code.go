package nsxt

// L2Vpn tunnel peer code
type L2VpnTunnelPeerCode struct {
	TransportTunnel *ResourceReference `json:"transport_tunnel"`
	// Copy this code to paste on the remote end of the tunnel. This is a base64 encoded string which has all the configuration for tunnel. E.g tap device local/peer ips and protocol, encryption algorithm, etc. The peer code also contains a pre-shared key; be careful when sharing or storing it.
	PeerCode string `json:"peer_code"`
}
