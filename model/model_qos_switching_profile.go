package nsxt

type QosSwitchingProfile struct {
	RequiredCapabilities []string `json:"required_capabilities,omitempty"`
	ResourceType string `json:"resource_type"`
	ShaperConfiguration []QosBaseRateShaper `json:"shaper_configuration,omitempty"`
	// Class of service
	ClassOfService int32 `json:"class_of_service,omitempty"`
	Dscp *Dscp `json:"dscp,omitempty"`
}
