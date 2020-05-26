package nsxt

// Information about a single LDAP server.
type IdentitySourceLdapServer struct {
	// If using LDAPS or STARTTLS, provide the X.509 certificate of the LDAP server in PEM format. This property is not required when connecting without TLS encryption and is ignored in that case.
	Certificates []string `json:"certificates,omitempty"`
	// A username used to authenticate to the directory when admnistering roles in NSX. This user should have privileges to search the LDAP directory for groups and users. This user is also used in some cases (OpenLDAP) to look up an NSX user's distinguished name based on their NSX login name. If omitted, NSX will authenticate to the LDAP server using an LDAP anonymous bind operation. For Active Directory, provide a userPrincipalName (e.g. administrator@airius.com) or the full distinguished nane. For OpenLDAP, provide the distinguished name of the user (e.g. uid=admin, cn=airius, dc=com).
	BindIdentity string `json:"bind_identity,omitempty"`
	// If set to true, Use the StartTLS extended operation to upgrade the connection to TLS before sending any sensitive information. The LDAP server must support the StartTLS extended operation in order for this protocol to operate correctly. This option is ignored if the URL scheme is LDAPS. 
	UseStarttls bool `json:"use_starttls,omitempty"`
	// The URL for the LDAP server. Supported URL schemes are LDAP and LDAPS. Either a hostname or an IP address may be given, and the port number is optional and defaults to 389 for the LDAP scheme and 636 for the LDAPS scheme.
	Url string `json:"url"`
	// A password used when authenticating to the directory.
	Password string `json:"password,omitempty"`
	// Allows the LDAP server to be enabled or disabled. When disabled, this LDAP server will not be used to authenticate users.
	Enabled bool `json:"enabled,omitempty"`
}
