package nsxt

// Duplicate IP detection and control
type DuplicateIpDetection struct {
	// Indicates whether duplicate IP detection should be enabled
	DuplicateIpDetectionEnabled bool `json:"duplicate_ip_detection_enabled,omitempty"`
}
