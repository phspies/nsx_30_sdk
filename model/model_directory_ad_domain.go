package nsxt

// Active Directory Domain
type DirectoryAdDomain struct {
	// Directory domain LDAP servers' information including host, name, port, protocol and so on.
	LdapServers []DirectoryLdapServer `json:"ldap_servers"`
	// Directory domain name which best describes the domain. It could be unique fqdn name or it could also be descriptive. There is no unique contraint for domain name among different domains.
	Name string `json:"name"`
	// Domain resource type comes from multiple sub-classes extending this base class. For example, DirectoryAdDomain is one accepted resource_type. If there are more sub-classes defined, they will also be accepted resource_type.
	ResourceType string `json:"resource_type"`
	// Each active directory domain has a domain naming context (NC), which contains domain-specific data. The root of this naming context is represented by a domain's distinguished name (DN) and is typically referred to as the NC head.
	BaseDistinguishedName string `json:"base_distinguished_name"`
	SyncSettings *DirectoryDomainSyncSettings `json:"sync_settings,omitempty"`
	// NetBIOS names can contain all alphanumeric characters except for the certain disallowed characters. Names can contain a period, but names cannot start with a period. NetBIOS is similar to DNS in that it can serve as a directory service, but more limited as it has no provisions for a name hierarchy and names are limited to 15 characters. The netbios name is case insensitive and is stored in upper case regardless of input case.
	NetbiosName string `json:"netbios_name"`
	SelectiveSyncSettings *SelectiveSyncSettings `json:"selective_sync_settings,omitempty"`
}
