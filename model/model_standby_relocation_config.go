package nsxt

// Standby service contexts relocation setting
type StandbyRelocationConfig struct {
	// The time interval (in minutes) to wait before starting the standby service context relocation process. 
	StandbyRelocationThreshold int64 `json:"standby_relocation_threshold,omitempty"`
}
