package nsxt

type ComputeManager struct {
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
	Credential *LoginCredential `json:"credential,omitempty"`
	// If the compute manager is VC and need to set set as OIDC provider for NSX then this flag should be set as true. This is specific to wcp feature, should be enabled when this feature is being used. 
	SetAsOidcProvider bool `json:"set_as_oidc_provider,omitempty"`
	// Specifies https port of the reverse proxy to connect to compute manager. For e.g. In case of VC, this port can be retrieved from this config file /etc/vmware-rhttpproxy/config.xml. 
	ReverseProxyHttpsPort int64 `json:"reverse_proxy_https_port,omitempty"`
	// Key-Value map of additional specific properties of compute manager
	OriginProperties []KeyValuePair `json:"origin_properties,omitempty"`
	// Compute manager type like vCenter
	OriginType string `json:"origin_type"`
	// IP address or hostname of compute manager
	Server string `json:"server"`
}
