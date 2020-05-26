package nsxt

// IPSec VPN session status summary, gives total, failed, degraded and established IPSec VPN sessions.
type IPsecVpnikeSessionSummary struct {
	// Number of established sessions.
	EstablishedSessions int64 `json:"established_sessions,omitempty"`
	// Total sessions configured.
	TotalSessions int64 `json:"total_sessions,omitempty"`
	// Number of failed sessions.
	FailedSessions int64 `json:"failed_sessions,omitempty"`
	// Number of degraded sessions.
	DegradedSessions int64 `json:"degraded_sessions,omitempty"`
}
