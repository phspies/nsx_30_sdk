package nsxt

type ServerSslProfileBinding struct {
	// A Certificate Revocation List (CRL) can be specified in the server-side SSL profile binding to disallow compromised server certificates. 
	ServerAuthCrlIds []string `json:"server_auth_crl_ids,omitempty"`
	// server authentication mode
	ServerAuth string `json:"server_auth,omitempty"`
	// authentication depth is used to set the verification depth in the server certificates chain. 
	CertificateChainDepth int64 `json:"certificate_chain_depth,omitempty"`
	// To support client authentication (load balancer acting as a client authenticating to the backend server), client certificate can be specified in the server-side SSL profile binding 
	ClientCertificateId string `json:"client_certificate_id,omitempty"`
	// If server auth type is REQUIRED, server certificate must be signed by one of the trusted Certificate Authorities (CAs), also referred to as root CAs, whose self signed certificates are specified. 
	ServerAuthCaIds []string `json:"server_auth_ca_ids,omitempty"`
	// Server SSL profile defines reusable, application-independent server side SSL properties. 
	SslProfileId string `json:"ssl_profile_id,omitempty"`
}