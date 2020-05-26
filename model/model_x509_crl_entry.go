package nsxt

// Each revoked certificate is identified in a CRL by its certificate serial number.
type X509CrlEntry struct {
	// Revocation date.
	RevocationDate string `json:"revocation_date,omitempty"`
	// The revoked certificate's serial number.
	SerialNumber string `json:"serial_number,omitempty"`
}
