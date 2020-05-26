package nsxt

type LdapIdentitySourceSearchResultItem struct {
	// Distinguished name (DN) of the entry.
	Dn string `json:"dn,omitempty"`
	// The Common Name (CN) of the entry, if available.
	CommonName string `json:"common_name,omitempty"`
	// For Active Directory (AD) users, this will be the user principal name (UPN), in the format user@domain. For non-AD users, this will be the user's uid property, followed by \"@\" and the domain of the directory. For groups, this will be the group's common name, followed by \"@\" and the domain of the directory.
	PrincipalName string `json:"principal_name,omitempty"`
	// Describes the type of the entry
	Type_ string `json:"type,omitempty"`
}
