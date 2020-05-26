package nsxt

// Policy api will overwrite the fipsGlobalConfig set using MP api. Always use https://<policyIp>/policy/api/v1/infra/global-config to update fips configuration.
type FipsGlobalConfig struct {
	// Valid Global configuration types
	ResourceType string `json:"resource_type"`
	// When this flag is set to true FIPS mode will be set on ssl encryptions of load balancer feature.
	LbFipsEnabled bool `json:"lb_fips_enabled,omitempty"`
}
