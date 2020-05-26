package nsxt

// Match condition for client certficate subject DN
type LbClientCertificateSubjectDnCondition struct {
	// If true, case is significant when comparing subject DN value. 
	CaseSensitive bool `json:"case_sensitive,omitempty"`
	// Match type of subject DN
	MatchType string `json:"match_type,omitempty"`
	// Value of subject DN
	SubjectDn string `json:"subject_dn"`
}
