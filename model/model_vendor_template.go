package nsxt

// Vendor Templates are registered by the partner service manager to be used in the service profile. They contain named (k-v) pairs.
type VendorTemplate struct {
	// Indicates system owned resource
	SystemOwned bool `json:"_system_owned,omitempty"`
	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`
	// Description of this resource
	Description string `json:"description,omitempty"`
	// Opaque identifiers meaningful to the API user
	Tags []Tag `json:"tags,omitempty"`
	// ID of the user who created this resource
	CreateUser string `json:"_create_user,omitempty"`
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed             to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed                 to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super                    user and can modify it, but only when providing                    the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this           entity. 
	Protection string `json:"_protection,omitempty"`
	// Timestamp of resource creation
	CreateTime int64 `json:"_create_time,omitempty"`
	// Timestamp of last modification
	LastModifiedTime int64 `json:"_last_modified_time,omitempty"`
	// ID of the user who last modified this resource
	LastModifiedUser string `json:"_last_modified_user,omitempty"`
	// Unique identifier of this resource
	Id string `json:"id,omitempty"`
	// The type of this resource.
	ResourceType string `json:"resource_type,omitempty"`
	// Different VMs in data center can have Different protection levels as specified by administrator in the policy. The identifier for the policy with which the partner appliance identifies this policy. This identifier will be passed to the partner appliance at runtime to specify which protection level is applicable for the VM being protected.
	VendorTemplateKey string `json:"vendor_template_key,omitempty"`
	// List of attributes specific to a partner for which the vendor template is created. There attributes are passed on to the partner appliance and is opaque to the NSX Manager. Attributes are not supported by guest introspection service.
	Attributes []Attribute `json:"attributes,omitempty"`
	// The redirection action represents if the packet is exclusively redirected to the service, or if a copy is forwarded to the service. Service profile inherits the redirection action specified at the vendor template and cannot override the action specified at the vendor template. Redirection action is not applicable to guest introspection service.
	RedirectionAction string `json:"redirection_action,omitempty"`
	// The capabilities provided by the services. Needs to be one of the following | NG_FW - Next Generation Firewall | IDS_IPS - Intrusion detection System / Intrusion Prevention System | NET_MON - Network Monitoring | HCX - Hybrid Cloud Exchange | BYOD - Bring Your Own Device | EPP - Endpoint Protection.(Third party AntiVirus partners using NXGI should use this functionality for the service)
	Functionality string `json:"functionality,omitempty"`
	// The service to which the vendor template belongs.
	ServiceId string `json:"service_id,omitempty"`
}
