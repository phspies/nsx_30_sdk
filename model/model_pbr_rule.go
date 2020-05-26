package nsxt

type PbrRule struct {
	Owner *OwnerResourceLink `json:"_owner,omitempty"`
	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`
	// Identifier of the resource
	Id string `json:"id,omitempty"`
	// The type of this resource.
	ResourceType string `json:"resource_type,omitempty"`
	// Description of this resource
	Description string `json:"description,omitempty"`
	// Flag to disable rule. Disabled will only be persisted but never provisioned/realized.
	Disabled bool `json:"disabled,omitempty"`
	// List of sources. Null will be treated as any.
	Sources []ResourceReference `json:"sources,omitempty"`
	// User level field which will be printed in CLI and packet logs.
	RuleTag string `json:"rule_tag,omitempty"`
	// List of the services. Null will be treated as any.
	Services []PbrService `json:"services,omitempty"`
	// User notes specific to the rule.
	Notes string `json:"notes,omitempty"`
	// List of object where rule will be enforced. field overrides this one. Null will be treated as any.
	AppliedTos []ResourceReference `json:"applied_tos,omitempty"`
	// Flag to enable packet logging. Default is disabled.
	Logged bool `json:"logged,omitempty"`
	// Action enforced on the packets which matches the PBR rule.
	Action string `json:"action"`
	// List of the destinations. Null will be treated as any.
	Destinations []ResourceReference `json:"destinations,omitempty"`
}