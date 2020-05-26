package nsxt

type BfdDiagnosticCount struct {
	// Number of tunnels with concatenated path down diagnostic message
	ConcatenatedPathDownCount int64 `json:"concatenated_path_down_count,omitempty"`
	// Number of tunnels with administratively down diagnostic message
	AdministrativelyDownCount int64 `json:"administratively_down_count,omitempty"`
	// Number of tunnels with no diagnostic
	NoDiagnosticCount int64 `json:"no_diagnostic_count,omitempty"`
	// Number of tunnels with path down diagnostic message
	PathDownCount int64 `json:"path_down_count,omitempty"`
	// Number of tunnels with reverse concatenated path down diagnostic message
	ReverseConcatenatedPathDownCount int64 `json:"reverse_concatenated_path_down_count,omitempty"`
	// Number of tunnels neighbor signaled session down
	NeighborSignaledSessionDownCount int64 `json:"neighbor_signaled_session_down_count,omitempty"`
	// Number of tunnels with control detection time expired diagnostic message
	ControlDetectionTimeExpiredCount int64 `json:"control_detection_time_expired_count,omitempty"`
	// Number of tunnels with echo function failed diagnostic message
	EchoFunctionFailedCount int64 `json:"echo_function_failed_count,omitempty"`
	// Number of tunnels with forwarding plane reset diagnostic message
	ForwardingPlaneResetCount int64 `json:"forwarding_plane_reset_count,omitempty"`
}
