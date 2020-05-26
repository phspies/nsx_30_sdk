package nsxt

type MigrationUnitGroup struct {
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
	// Number of migration units in the group
	MigrationUnitCount int32 `json:"migration_unit_count,omitempty"`
	// List of migration units in the group
	MigrationUnits []MigrationUnit `json:"migration_units,omitempty"`
	// Flag to indicate whether migration of this group is enabled or not
	Enabled bool `json:"enabled,omitempty"`
	// Component type
	Type_ string `json:"type"`
	// Migration method to specify whether the migration is to be performed in parallel or serially
	Parallel bool `json:"parallel,omitempty"`
	// Extended configuration for the group
	ExtendedConfiguration []KeyValuePair `json:"extended_configuration,omitempty"`
}
