package nsxt

// Summarized view of all selected IPSec VPN sessions.
type IpSecVpnSessionSummary struct {
	// Traffic summary per session.
	TrafficSummaryPerSession []IpSecVpnSessionTrafficSummary `json:"traffic_summary_per_session,omitempty"`
	// Timestamp when the data was last updated.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	SessionSummary *IPsecVpnikeSessionSummary `json:"session_summary,omitempty"`
	AggregateTrafficCounters *IpSecVpnTrafficCounters `json:"aggregate_traffic_counters,omitempty"`
}
