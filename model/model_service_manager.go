package nsxt

// Partner console endpoint information for enabling NSX to callback with events and status.
type ServiceManager struct {
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
	// Integer port value to specify a standard/non-standard HTTPS port.
	Port int64 `json:"port"`
	// The IDs of services, provided by partner.
	ServiceIds []ResourceReference `json:"service_ids"`
	AuthenticationScheme *CallbackAuthenticationScheme `json:"authentication_scheme"`
	// Thumbprint (SHA-256 hash represented in lower case hex) for the certificate on the partner console. This will be required to establish secure communication with the console and to avoid man-in-the-middle attacks.
	Thumbprint string `json:"thumbprint,omitempty"`
	// Id which is unique to a vendor or partner for which the service is created.
	VendorId string `json:"vendor_id,omitempty"`
	// URI on which notification requests should be made on the specified server.
	Uri string `json:"uri"`
	// IP address or fully qualified domain name of the partner REST server.
	Server string `json:"server"`
}
