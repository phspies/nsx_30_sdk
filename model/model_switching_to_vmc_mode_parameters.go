package nsxt

// Possible values of a mode in a \"/config/nsx_appliance_mode\" file
type SwitchingToVmcModeParameters struct {
	// Possible enum values in a \"/config/nsx_appliance_mode\" file
	ModeId string `json:"mode_id"`
	SreOrg *OrgInfo `json:"sre_org,omitempty"`
	// Org ID of a Client - commonly UUID.
	DefaultOrgId string `json:"default_org_id,omitempty"`
	EaOrg *OrgInfo `json:"ea_org,omitempty"`
	GssOrg *OrgInfo `json:"gss_org,omitempty"`
	// IP/host of PoP (Point-of-Presence) HTTP proxy server
	ProxyHost string `json:"proxy_host,omitempty"`
	// CSP time drift in milliseconds
	CspTimeDrift int64 `json:"csp_time_drift,omitempty"`
	// SDDC id
	SddcId string `json:"sddc_id,omitempty"`
	// List of whitelist IPs for basic auth
	BasicAuthWhitelistIps []string `json:"basic_auth_whitelist_ips,omitempty"`
	// Protocol and domain name (or IP address) of a CSP server, like \"https://console-stg.cloud.vmware.com\".
	BaseUrl string `json:"base_url,omitempty"`
	// Port of PoP (Point-of-Presence) Http proxy server
	ProxyPort int64 `json:"proxy_port,omitempty"`
	// Relative path on CSP server to the Org location. Can be \"/csp/gateway/am/api/orgs/\".
	CspOrgUri string `json:"csp_org_uri,omitempty"`
	CspClientCredential *Oauth2Credentials `json:"csp_client_credential,omitempty"`
	AuthCode *Oauth2Credentials `json:"auth_code,omitempty"`
	// When this parameter is set to true, only a change of the node mode happens without any update to the auth properties. When this param is not set to true i.e. set to false or not provided, mode change and update to the auth properties will both happen.
	ModeChangeOnly bool `json:"mode_change_only,omitempty"`
	// List of incoming client IDs
	CspClientIncomingCredentials []string `json:"csp_client_incoming_credentials,omitempty"`
	// Service definition id
	ServiceDefinitionId string `json:"service_definition_id,omitempty"`
	// Node Mode type
	ResourceType string `json:"resource_type,omitempty"`
}
