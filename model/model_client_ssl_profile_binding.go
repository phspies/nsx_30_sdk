package nsxt

type ClientSslProfileBinding struct {
	// client authentication mode
	ClientAuth string `json:"client_auth,omitempty"`
	// Client SSL profile defines reusable, application-independent client side SSL properties. 
	SslProfileId string `json:"ssl_profile_id,omitempty"`
	// authentication depth is used to set the verification depth in the client certificates chain. 
	CertificateChainDepth int64 `json:"certificate_chain_depth,omitempty"`
	// If client auth type is REQUIRED, client certificate must be signed by one of the trusted Certificate Authorities (CAs), also referred to as root CAs, whose self signed certificates are specified. 
	ClientAuthCaIds []string `json:"client_auth_ca_ids,omitempty"`
	// A default certificate should be specified which will be used if the server does not host multiple hostnames on the same IP address or if the client does not support SNI extension. 
	DefaultCertificateId string `json:"default_certificate_id"`
	// Client-side SSL profile binding allows multiple certificates, for different hostnames, to be bound to the same virtual server. 
	SniCertificateIds []string `json:"sni_certificate_ids,omitempty"`
	// A Certificate Revocation List (CRL) can be specified in the client-side SSL profile binding to disallow compromised client certificates. 
	ClientAuthCrlIds []string `json:"client_auth_crl_ids,omitempty"`
}
