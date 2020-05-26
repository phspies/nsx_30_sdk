package nsxt

// A list of LDA entries returned from a search of an LDAP identity source.
type LdapIdentitySourceSearchResultList struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	Results []LdapIdentitySourceSearchResultItem `json:"results,omitempty"`
}