package nsxt

// Reference to a CRL Distribution Point where to fetch a CRL
type CrlDistributionPointStatus struct {
	// Status of the fetched CRL for this CrlDistributionPoint
	Status string `json:"status,omitempty"`
	// Error message when fetching the CRL failed.
	ErrorMessage string `json:"error_message,omitempty"`
}
