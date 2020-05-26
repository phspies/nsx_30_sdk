package nsxt

// An identity source service that runs Microsoft Active Directory. The service allows selected user accounts defined in Active Directory to log into and access NSX-T.
type ActiveDirectoryIdentitySource struct {
	// The list of LDAP servers that provide LDAP service for this identity source. Currently, only one LDAP server is supported.
	LdapServers []IdentitySourceLdapServer `json:"ldap_servers,omitempty"`
	// The name of the authentication domain. When users log into NSX using an identity of the form \"user@domain\", NSX uses the domain portion to determine which LDAP identity source to use. For Active Directory, this domain name must match the domain of the Active Directory.
	DomainName string `json:"domain_name"`
	ResourceType string `json:"resource_type"`
	// The subtree of the LDAP identity source to search when locating users and groups.
	BaseDn string `json:"base_dn"`
}
