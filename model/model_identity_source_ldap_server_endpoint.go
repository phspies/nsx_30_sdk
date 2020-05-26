package nsxt

// Information about a single LDAP server endpoint.
type IdentitySourceLdapServerEndpoint struct {
	// The URL for the LDAP server. Supported URL schemes are LDAP and LDAPS. Either a hostname or an IP address may be given, and the port number is optional and defaults to 389 for the LDAP scheme and 636 for the LDAPS scheme.
	Url string `json:"url"`
	// If set to true, Use the StartTLS extended operation to upgrade the connection to TLS before sending any sensitive information. The LDAP server must support the StartTLS extended operation in order for this protocol to operate correctly. This option is ignored if the URL scheme is LDAPS. 
	UseStarttls bool `json:"use_starttls,omitempty"`
}
