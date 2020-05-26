package nsxt

type TunnelStatusCount struct {
	// Roll-up status
	Status string `json:"status,omitempty"`
	// Degraded count
	DegradedCount int32 `json:"degraded_count,omitempty"`
	// Down count
	DownCount int32 `json:"down_count,omitempty"`
	// Up count
	UpCount int32 `json:"up_count,omitempty"`
	BfdDiagnostic *BfdDiagnosticCount `json:"bfd_diagnostic,omitempty"`
	BfdStatus *BfdStatusCount `json:"bfd_status,omitempty"`
}
