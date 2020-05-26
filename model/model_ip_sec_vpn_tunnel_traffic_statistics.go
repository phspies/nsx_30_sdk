package nsxt

// IPSec VPN tunnel traffic statistics.
type IpSecVpnTunnelTrafficStatistics struct {
	// Total number of packets dropped while sending for any reason.
	PacketsSentOtherError int64 `json:"packets_sent_other_error,omitempty"`
	// Total number of outgoing packets on outbound Security association (SA).
	PacketsOut int64 `json:"packets_out,omitempty"`
	// Total number of outgoing packets dropped on outbound security association.
	DroppedPacketsOut int64 `json:"dropped_packets_out,omitempty"`
	// Total number of packets dropped due to integrity failures.
	IntegrityFailures int64 `json:"integrity_failures,omitempty"`
	// Number of packets dropped because of no matching policy is available.
	NomatchingPolicyErrors int64 `json:"nomatching_policy_errors,omitempty"`
	// Totoal number of security association (SA) mismatch errors on incoming packets.
	SaMismatchErrorsIn int64 `json:"sa_mismatch_errors_in,omitempty"`
	// Peer subnet to which a tunnel belongs.
	PeerSubnet string `json:"peer_subnet,omitempty"`
	// Total number of packets dropped due to replay check on that Security association (SA).
	ReplayErrors int64 `json:"replay_errors,omitempty"`
	// Total number of outgoing bytes on outbound Security association (SA).
	BytesOut int64 `json:"bytes_out,omitempty"`
	// Local subnet to which a tunnel belongs.
	LocalSubnet string `json:"local_subnet,omitempty"`
	// Total number of incoming packets dropped on inbound security association.
	DroppedPacketsIn int64 `json:"dropped_packets_in,omitempty"`
	// Total number of packets dropped because of failure in encryption.
	EncryptionFailures int64 `json:"encryption_failures,omitempty"`
	// Totoal number of security association (SA) mismatch errors on outgoing packets.
	SaMismatchErrorsOut int64 `json:"sa_mismatch_errors_out,omitempty"`
	// Gives the detailed reason about the tunnel when it is down. If tunnel is UP tunnel down reason will be empty.
	TunnelDownReason string `json:"tunnel_down_reason,omitempty"`
	// Total number of incoming packets dropped on inbound Security association (SA)(misc).
	PacketsReceiveOtherError int64 `json:"packets_receive_other_error,omitempty"`
	// Total number of incoming bytes on inbound Security association (SA).
	BytesIn int64 `json:"bytes_in,omitempty"`
	// Total number of packets dropped due to decryption failures.
	DecryptionFailures int64 `json:"decryption_failures,omitempty"`
	// Total number of packets dropped while sending due to overflow in sequence number.
	SeqNumberOverflowError int64 `json:"seq_number_overflow_error,omitempty"`
	// Total number of incoming packets on inbound Security association (SA).
	PacketsIn int64 `json:"packets_in,omitempty"`
	// Specifies the status of tunnel. If all the SA (Security association) are negotiated then tunnels status will be UP. If negotiation fails for the SAs status will be DOWN, if SAs are in negotiating phase tunnels status will be NEGOTIATING.
	TunnelStatus string `json:"tunnel_status,omitempty"`
}
