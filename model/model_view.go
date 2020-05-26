package nsxt

// Describes the configuration of a view to be displayed on the dashboard.
type View struct {
	// Indicates system owned resource
	SystemOwned bool `json:"_system_owned,omitempty"`
	// Title of the widget.
	DisplayName string `json:"display_name"`
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
	// Comma separated list of roles to which the shared view is visible. Allows user to specify the visibility of a shared view to the specified roles. User defined roles can also be specified in the list. The roles can be obtained via GET /api/v1/aaa/roles. Please visit API documentation for details about roles.
	IncludeRoles string `json:"include_roles,omitempty"`
	// Comma separated list of roles to which the shared view is not visible. Allows user to prevent the visibility of a shared view to the specified roles. User defined roles can also be specified in the list. The roles can be obtained via GET /api/v1/aaa/roles. Please visit API documentation for details about roles. If include_roles is specified then exclude_roles cannot be specified.
	ExcludeRoles string `json:"exclude_roles,omitempty"`
	// Determines placement of view relative to other views. The lower the weight, the higher it is in the placement order.
	Weight int32 `json:"weight,omitempty"`
	// Array of widgets that are part of the view.
	Widgets []WidgetItem `json:"widgets"`
	// Defaults to false. Set to true to publish the view to other users. The widgets of a shared view are visible to other users.
	Shared bool `json:"shared,omitempty"`
}
