package nsxt

// The deployment of a registered service. service instance is instantiation of service.
type BaseServiceInstance struct {
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
	// ServiceInstance is used when NSX handles the lifecyle of   appliance. Deployment and appliance related all the information is necessary. ByodServiceInstance is a custom instance to be used when NSX is not handling   the lifecycles of appliance/s. User will manage their own appliance (BYOD)   to connect with NSX. VirtualServiceInstance is a a custom instance to be used when NSX is not   handling the lifecycle of an appliance and when the user is not bringing   their own appliance. 
	ResourceType string `json:"resource_type"`
	// Failure policy of the service instance - if it has to be different from the service. By default the service instance inherits the FailurePolicy of the service it belongs to.
	OnFailurePolicy string `json:"on_failure_policy,omitempty"`
	// Transport to be used by this service instance for deploying the Service-VM. This field is to be set Not Applicable(NA) if the service only caters to functionality EPP(Endpoint Protection).
	TransportType string `json:"transport_type"`
	// The Service to which the service instance is associated.
	ServiceId string `json:"service_id,omitempty"`
}