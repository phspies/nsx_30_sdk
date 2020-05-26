package nsxt

type PrincipalIdentityWithCertificate struct {
	// Indicator whether the entities created by this principal should be protected.
	IsProtected bool `json:"is_protected,omitempty"`
	// Role
	Role string `json:"role,omitempty"`
	// Name of the principal.
	Name string `json:"name"`
	// Use the 'role' field instead and pass in 'auditor' for read_only_api_users or 'enterprise_admin' for the others.
	PermissionGroup string `json:"permission_group,omitempty"`
	// Id of the stored certificate. When used with the deprecated POST /trust-management/principal-identities API this field is required.
	CertificateId string `json:"certificate_id,omitempty"`
	// Unique node-id of a principal.
	NodeId string `json:"node_id"`
	// PEM encoding of the new certificate.
	CertificatePem string `json:"certificate_pem"`
}
