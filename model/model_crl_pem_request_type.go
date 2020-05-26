package nsxt

// Request Type to get a CRL's PEM file.
type CrlPemRequestType struct {
	// CRL Distribution Point URI where to fetch the CRL.
	CdpUri string `json:"cdp_uri,omitempty"`
}
