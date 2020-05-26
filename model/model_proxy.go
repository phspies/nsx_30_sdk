package nsxt

// Type to define the Proxy configuration.
type Proxy struct {
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
	// Specify the user name used to authenticate with the proxy server, if required. 
	Username string `json:"username,omitempty"`
	// Specify the fully qualified domain name, or ip address, of the proxy server. 
	Host string `json:"host"`
	// Specify the password used to authenticate with the proxy server, if required. 
	Password string `json:"password,omitempty"`
	// The scheme accepted by the proxy server. Specify one of HTTP and HTTPS. 
	Scheme string `json:"scheme"`
	// Flag to indicate if proxy is enabled. When set to true, a scheme, host and port must be provided. 
	Enabled bool `json:"enabled"`
	// Specify the port of the proxy server.
	Port int32 `json:"port"`
}
