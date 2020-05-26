package nsxt

// GI Specific service profile
type GiServiceProfile struct {
	// Service Profile type, for example 'GiServiceProfile', 'ServiceInsertionServiceProfile'
	ResourceType string `json:"resource_type"`
	// Different VMs in data center can have Different protection levels as specified by administrator in the policy. The identifier for the policy with which the partner appliance identifies this policy. This identifier will be passed to the partner appliance at runtime to specify which protection level is applicable for the VM being protected.
	VendorTemplateKey string `json:"vendor_template_key,omitempty"`
	// The service to which the service profile belongs.
	ServiceId string `json:"service_id,omitempty"`
	// ID of the vendor template, created by partner while registering the service.
	VendorTemplateId string `json:"vendor_template_id"`
}
