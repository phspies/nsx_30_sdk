package nsxt

// Match condition for client certficate issuer DN
type LbClientCertificateIssuerDnCondition struct {
	// If true, case is significant when comparing issuer DN value. 
	CaseSensitive bool `json:"case_sensitive,omitempty"`
	// Match type of issuer DN
	MatchType string `json:"match_type,omitempty"`
	// Value of issuer DN
	IssuerDn string `json:"issuer_dn"`
}
