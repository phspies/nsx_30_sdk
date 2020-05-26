package nsxt

// Service profile represents a specialization of vendor template.
type ServiceInsertionServiceProfile struct {
	// Service Profile type, for example 'GiServiceProfile', 'ServiceInsertionServiceProfile'
	ResourceType string `json:"resource_type"`
	// List of attributes specific to a partner for which the service is created. These attributes are passed on to the partner appliance and are opaque to the NSX Manager. If a vendor template exposes configurables, then the values are specified here.
	Attributes []Attribute `json:"attributes,omitempty"`
	// The service to which the service profile belongs.
	ServiceId string `json:"service_id,omitempty"`
	// The redirection action represents if the packet is exclusively redirected to the service, or if a copy is forwarded to the service. The service insertion profile inherits the redirection action if already specified at the vendor template. However the service profile cannot overide the action specified at the vendor template.
	RedirectionAction string `json:"redirection_action,omitempty"`
	// Id of the vendor template to be used by the servive profile.
	VendorTemplateId string `json:"vendor_template_id"`
}
