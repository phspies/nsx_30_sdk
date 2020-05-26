package nsxt

type ServiceEndpoint struct {
	CertificateSha256Thumbprint string `json:"certificate_sha256_thumbprint,omitempty"`
	// Certificate or certificate chain
	Certificate string `json:"certificate,omitempty"`
	// List of entities hosted on accessible through the service endpoint
	EntitiesHosted []HostedEntityInfo `json:"entities_hosted,omitempty"`
	// IPv4 or IPv6 address
	IpAddress string `json:"ip_address"`
	Fqdn string `json:"fqdn,omitempty"`
	// Unique identifier of this service endpoint
	ServiceEndpointUuid string `json:"service_endpoint_uuid,omitempty"`
	Port int64 `json:"port"`
}
