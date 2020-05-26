package nsxt

// BGP Graceful Restart timers configuration
type GracefulRestartTimer struct {
	// Maximum time BGP speaker will take for the BGP session to be re-established after a restart. Ranges from 1 sec to 3600 sec. This can be used to speed up routing convergence by its peer in case that the BGP speaker does not come back after a restart. If the session does not get re-established within the \"Restart Time\" that the Restarting Speaker advertised previously, the Receiving Speaker will delete all the stale routes from that peer. 
	RestartTimer int64 `json:"restart_timer,omitempty"`
	// Maximum time before stale routes are removed from the RIB when the local BGP process restarts. Ranges from 1 sec to 3600 sec. 
	StaleTimer int64 `json:"stale_timer,omitempty"`
}