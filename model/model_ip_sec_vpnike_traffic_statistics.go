package nsxt

// Traffic statistics for IPSec VPN IKE session. Note - Not supported in this release.
type IpSecVpnikeTrafficStatistics struct {
	// Number of packets out.
	PacketsOut int64 `json:"packets_out,omitempty"`
	// Fail count.
	FailCount int64 `json:"fail_count,omitempty"`
	// Number of packets in.
	PacketsIn int64 `json:"packets_in,omitempty"`
	// Number of bytes out.
	BytesOut int64 `json:"bytes_out,omitempty"`
	// Number of bytes in.
	BytesIn int64 `json:"bytes_in,omitempty"`
}
