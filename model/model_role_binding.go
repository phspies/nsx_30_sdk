package nsxt

// User/Group's role binding
type RoleBinding struct {
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
	// Identity source type
	IdentitySourceType string `json:"identity_source_type,omitempty"`
	// User/Group's name
	Name string `json:"name,omitempty"`
	// Roles
	Roles []Role `json:"roles,omitempty"`
	// Type
	Type_ string `json:"type,omitempty"`
	// Property 'stale' can be considered to have these values - absent  - This type of rolebinding does not support stale property TRUE    - Rolebinding is stale in vIDM meaning the user is no longer present in vIDM FALSE   - Rolebinding is available in vIDM UNKNOWN - Rolebinding's state of staleness in unknown Once rolebindings become stale, they can be deleted using the API POST /aaa/role-bindings?action=delete_stale_bindings
	Stale string `json:"stale,omitempty"`
	// The ID of the external identity source that holds the referenced external entity. Currently, only external LDAP servers are allowed.
	IdentitySourceId string `json:"identity_source_id,omitempty"`
}
