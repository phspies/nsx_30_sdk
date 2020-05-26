package nsxt

// Service chain is a set of network Services. A Service chain is made up of ordered list of service profiles belonging to any same or different services.
type ServiceChain struct {
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
	// List of ServiceInsertionServiceProfiles id. Reverse path service profiles are applied to egress traffic and is optional. 2 different set of profiles can be defined for forward and reverse path. If not defined, the reverse of the forward path service profile is applied.
	ReversePathServiceProfiles []ResourceReference `json:"reverse_path_service_profiles,omitempty"`
	// Service attachment specifies the scope i.e Service plane at which the SVMs are deployed.
	ServiceAttachments []ResourceReference `json:"service_attachments"`
	// List of ServiceInsertionServiceProfiles that constitutes the the service chain. The forward path service profiles are applied to ingress traffic.
	ForwardPathServiceProfiles []ResourceReference `json:"forward_path_service_profiles"`
	// A unique id generated for every service chain. This is not a uuid.
	ServiceChainId string `json:"service_chain_id,omitempty"`
	// Path selection policy can be - ANY - Service Insertion is free to redirect to any service path regardless of any load balancing considerations or flow pinning. LOCAL - means to prefer local service insances. REMOTE - preference is to redirect to the SVM co-located on the same host.
	PathSelectionPolicy string `json:"path_selection_policy,omitempty"`
	// Failure policy for the service tells datapath, the action to take i.e to allow or block traffic during failure scenarios.
	OnFailurePolicy string `json:"on_failure_policy,omitempty"`
}
