package nsxt

// BGP configuration for Tier0 logical router. We create BGP configuration while creation of Tier0 logical router. 
type BgpConfig struct {
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
	InterSrIbgp *InterSrRoutingConfig `json:"inter_sr_ibgp,omitempty"`
	// This is a deprecated property, Please use 'as_num' instead. For VRF logical router, the as_number from parent logical router will be effective.
	AsNumber int64 `json:"as_number,omitempty"`
	// List of routes to be aggregated
	RouteAggregation []BgpRouteAggregation `json:"route_aggregation,omitempty"`
	// Logical router id
	LogicalRouterId string `json:"logical_router_id,omitempty"`
	// Flag to enable graceful restart. This field is deprecated, kindly use graceful_restart_config parameter for graceful restart configuration. If both parameters are set and consistent with each other [i.e. graceful_restart=false and graceful_restart_mode=HELPER_ONLY OR graceful_restart=true and graceful_restart_mode=GR_AND_HELPER] then this is allowed, but if inconsistent with each other then this is not allowed and validation error will be thrown. For VRF logical router, the settings from parent logical router will be effective. 
	GracefulRestart bool `json:"graceful_restart,omitempty"`
	// For VRF logical router, the as_num from parent logical router will be effective.
	AsNum string `json:"as_num,omitempty"`
	// While creation of BGP config this flag will be set to - true for Tier0 logical router with Active-Active high-availability mode - false for Tier0 logical router with Active-Standby high-availanility mode. User can change this value while updating the config. If this property is not specified in the payload, the default value will be considered as false irrespective of the high-availability mode. 
	Enabled bool `json:"enabled,omitempty"`
	GracefulRestartConfig *GracefulRestartConfig `json:"graceful_restart_config,omitempty"`
	// For TIER0 logical router, default is true. For VRF logical router, the settings from parent logical router will be effective.
	MultipathRelax bool `json:"multipath_relax,omitempty"`
	// While creation of BGP config this flag will be set to true User can change this value while updating BGP config. If this property is not specified in the payload, the default value will be considered as true. 
	Ecmp bool `json:"ecmp,omitempty"`
}