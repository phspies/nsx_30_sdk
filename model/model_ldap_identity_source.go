package nsxt

// This is the base type for all identity sources that use LDAP for authentication and group membership.
type LdapIdentitySource struct {
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
	ResourceType string `json:"resource_type"`
	// The list of LDAP servers that provide LDAP service for this identity source. Currently, only one LDAP server is supported.
	LdapServers []IdentitySourceLdapServer `json:"ldap_servers,omitempty"`
	// The name of the authentication domain. When users log into NSX using an identity of the form \"user@domain\", NSX uses the domain portion to determine which LDAP identity source to use. For Active Directory, this domain name must match the domain of the Active Directory.
	DomainName string `json:"domain_name"`
	// The subtree of the LDAP identity source to search when locating users and groups.
	BaseDn string `json:"base_dn"`
}
