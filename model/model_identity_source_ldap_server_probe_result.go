package nsxt

// The results of probing an individual LDAP server.
type IdentitySourceLdapServerProbeResult struct {
	// THe URL of the probed LDAP host.
	Url string `json:"url,omitempty"`
	// Detail about errors encountered during the probe.
	Errors []LdapProbeError `json:"errors,omitempty"`
	// Overall result of the probe. If the probe was able to connect to the LDAP service, authenticate using the provided credentials, and perform searches of the configured user and group search bases without error, the result is SUCCESS.  Otherwise, the result is FAILURE, and additional details may be found in the errors property.
	Result string `json:"result,omitempty"`
}
