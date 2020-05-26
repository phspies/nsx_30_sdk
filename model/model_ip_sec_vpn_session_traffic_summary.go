package nsxt

// IPSec VPN session traffic summary.
type IpSecVpnSessionTrafficSummary struct {
	TrafficCounters *IpSecVpnTrafficCounters `json:"traffic_counters,omitempty"`
	IpsecVpnSession *ResourceReference `json:"ipsec_vpn_session,omitempty"`
}
