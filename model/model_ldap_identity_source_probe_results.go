package nsxt

// Results from probing all LDAP servers in an LDAP identity source configuration.
type LdapIdentitySourceProbeResults struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Probe results for all probed LDAP servers.
	Results []IdentitySourceLdapServerProbeResult `json:"results,omitempty"`
}
