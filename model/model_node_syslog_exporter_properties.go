package nsxt

// Node syslog exporter properties
type NodeSyslogExporterProperties struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// CA certificate PEM of TLS server to export to
	TlsCaPem string `json:"tls_ca_pem,omitempty"`
	// Export protocol
	Protocol string `json:"protocol"`
	// Syslog exporter name
	ExporterName string `json:"exporter_name"`
	// Logging level to export
	Level string `json:"level"`
	// CA certificate PEM of the rsyslog client
	TlsClientCaPem string `json:"tls_client_ca_pem,omitempty"`
	// Certificate PEM of the rsyslog client
	TlsCertPem string `json:"tls_cert_pem,omitempty"`
	// IP address or hostname of server to export to
	Server string `json:"server"`
	// Facilities to export
	Facilities []string `json:"facilities,omitempty"`
	// MSGIDs to export
	Msgids []string `json:"msgids,omitempty"`
	// Structured data to export
	StructuredData []string `json:"structured_data,omitempty"`
	// Port to export to
	Port int64 `json:"port,omitempty"`
	// Private key PEM of the rsyslog client
	TlsKeyPem string `json:"tls_key_pem,omitempty"`
}
