package nsxt

type PrincipalIdentity struct {
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
	// Indicator whether the entities created by this principal should be protected.
	IsProtected bool `json:"is_protected,omitempty"`
	// Role
	Role string `json:"role,omitempty"`
	// Name of the principal.
	Name string `json:"name"`
	// Use the 'role' field instead and pass in 'auditor' for read_only_api_users or 'enterprise_admin' for the others.
	PermissionGroup string `json:"permission_group,omitempty"`
	// Id of the stored certificate. When used with the deprecated POST /trust-management/principal-identities API this field is required.
	CertificateId string `json:"certificate_id,omitempty"`
	// Unique node-id of a principal.
	NodeId string `json:"node_id"`
}
