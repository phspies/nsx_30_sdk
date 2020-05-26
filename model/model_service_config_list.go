package nsxt

// List of Service config objects that needs to be either created or updated with the respective profiles and precedence. 
type ServiceConfigList struct {
	// An Array of ServiceConfig objects containing details of profiles to be applied, entities on which these profiles will be applied and precedence. 
	ServiceConfigs []ServiceConfig `json:"service_configs"`
}
