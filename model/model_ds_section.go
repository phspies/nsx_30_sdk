package nsxt

type DsSection struct {
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
	// Stateful or Stateless nature of distributed service section is enforced on all rules inside the section. Layer3 sections can be stateful or stateless. Layer2 sections can only be stateless.
	Stateful bool `json:"stateful"`
	// It is a boolean flag which reflects whether a distributed service section is default section or not. Each Layer 3 and Layer 2 section will have at least and at most one default section.
	IsDefault bool `json:"is_default,omitempty"`
	// List of objects where the rules in this section will be enforced. This will take precedence over rule level appliedTo.
	AppliedTos []ResourceReference `json:"applied_tos,omitempty"`
	// Number of rules in this section.
	RuleCount int64 `json:"rule_count,omitempty"`
	// Type of the rules which a section can contain. Only homogeneous sections are supported.
	SectionType string `json:"section_type"`
}
