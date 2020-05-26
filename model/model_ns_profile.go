package nsxt

// An entity that encapsulates attributes and sub-attributes of various network services (ex. L7 services,domain name,encryption algorithm) The entity will be consumed in DFW rules and can be added in new tuple called profile in DFW rules. This entity is design to be generic and can be consumed at other places as well where attributes and sub-attributes collection can be used. To get a list of supported attributes and sub-attributes fire the following REST API GET https://&lt;nsx-mgr&gt;/api/v1/ns-profiles/attributes 
type NsProfile struct {
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
	// Reference to the encapsulating object of attributes/sub-attributes for NSProfile. 
	NsprofileAttribute []NsAttributes `json:"nsprofile_attribute"`
	// If set to false, the NSProfile has some app ids which are unsupported. Those were allowed to be added in previous releases but in testing in later phases found that those app ids could not be detected. 
	IsValid bool `json:"is_valid,omitempty"`
}
