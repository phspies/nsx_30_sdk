package nsxt

// NTP Service properties
type NtpServiceProperties struct {
	// Start NTP service when system boots
	StartOnBoot bool `json:"start_on_boot,omitempty"`
	// NTP servers
	Servers []string `json:"servers"`
}
