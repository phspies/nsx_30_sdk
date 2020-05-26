package nsxt

// Get the peer_code for each tunnel to paste on the remote end of the tunnel. Currently only stand-along/unmanaged edge is supported on the remote end of the tunnel.
type L2VpnSessionPeerCodes struct {
	// List of peer codes per transport tunnel.
	PeerCodes []L2VpnTunnelPeerCode `json:"peer_codes"`
}
