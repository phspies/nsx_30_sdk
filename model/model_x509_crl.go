package nsxt

// A CRL is a time-stamped list identifying revoked certificates.
type X509Crl struct {
	// Next update time for the CRL.
	NextUpdate string `json:"next_update,omitempty"`
	// CRL's version number either 1 or 2.
	Version string `json:"version,omitempty"`
	// List of X509CrlEntry.
	CrlEntries []X509CrlEntry `json:"crl_entries,omitempty"`
	// Issuer's distinguished name. (DN)
	Issuer string `json:"issuer,omitempty"`
}