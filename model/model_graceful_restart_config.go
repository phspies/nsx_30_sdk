package nsxt

// BGP Graceful Restart configuration parameters
type GracefulRestartConfig struct {
	// BGP Graceful Restart mode
	GracefulRestartMode string `json:"graceful_restart_mode,omitempty"`
	GracefulRestartTimer *GracefulRestartTimer `json:"graceful_restart_timer,omitempty"`
}
