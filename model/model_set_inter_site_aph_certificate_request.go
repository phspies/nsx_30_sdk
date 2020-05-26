package nsxt

// Data for setting Appliance Proxy certificate for inter-site communication
type SetInterSiteAphCertificateRequest struct {
	// ID of the certificate that is already imported.
	CertId string `json:"cert_id,omitempty"`
	// ID of the node that this certificate is used on.
	UsedById string `json:"used_by_id,omitempty"`
}
