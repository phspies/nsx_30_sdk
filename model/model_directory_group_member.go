package nsxt

type DirectoryGroupMember struct {
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
	// Directory group id this group member belongs to
	GroupId string `json:"group_id"`
	// Directory group name this group member owns
	MemberGroupDistinguishedName string `json:"member_group_distinguished_name"`
	// Directory group id this group member owns
	MemberGroupId string `json:"member_group_id"`
	// Directory group distinguished name this group member belongs to.
	GroupDistinguishedName string `json:"group_distinguished_name"`
	// Whether this member group is a directory member of the parent group speicified by group_id or a nested member group which parent group is also member group of the parent group speicified by group_id.
	Nested bool `json:"nested"`
}
