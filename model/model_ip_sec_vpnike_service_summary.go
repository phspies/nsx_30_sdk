package nsxt

// Summarized view of all IPSec VPN sessions for a specified service.
type IpSecVpnikeServiceSummary struct {
	// Traffic summary per session.
	TrafficSummaryPerSession []IpSecVpnSessionTrafficSummary `json:"traffic_summary_per_session,omitempty"`
	// Timestamp when the data was last updated.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	SessionSummary *IPsecVpnikeSessionSummary `json:"session_summary,omitempty"`
	AggregateTrafficCounters *IpSecVpnTrafficCounters `json:"aggregate_traffic_counters,omitempty"`
	// UUID for a vpn service.
	IpsecVpnServiceId string `json:"ipsec_vpn_service_id,omitempty"`
	// VPN service display name.
	DisplayName string `json:"display_name,omitempty"`
	// Logical router identifier associated with vpn service.
	LogicalRouterId string `json:"logical_router_id,omitempty"`
}
