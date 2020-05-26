package nsxt

type LbHttpsMonitor struct {
	// If the monitor port is specified, it would override pool member port setting for healthcheck. A port range is not supported. For ICMP monitor, monitor_port is not required. 
	MonitorPort string `json:"monitor_port,omitempty"`
	// num of consecutive checks must fail before marking it down
	FallCount int64 `json:"fall_count,omitempty"`
	// the frequency at which the system issues the monitor check (in second)
	Interval int64 `json:"interval,omitempty"`
	// num of consecutive checks must pass before marking it up
	RiseCount int64 `json:"rise_count,omitempty"`
	// the number of seconds the target has in which to respond to the monitor request 
	Timeout int64 `json:"timeout,omitempty"`
	// The HTTP response status code should be a valid HTTP status code. 
	ResponseStatusCodes []int32 `json:"response_status_codes,omitempty"`
	// A Certificate Revocation List (CRL) can be specified in the server-side SSL profile binding to disallow compromised server certificates. 
	ServerAuthCrlIds []string `json:"server_auth_crl_ids,omitempty"`
	// If server auth type is REQUIRED, server certificate must be signed by one of the trusted Certificate Authorities (CAs), also referred to as root CAs, whose self signed certificates are specified. 
	ServerAuthCaIds []string `json:"server_auth_ca_ids,omitempty"`
	// server authentication mode
	ServerAuth string `json:"server_auth,omitempty"`
	// String to send as part of HTTP health check request body. Valid only for certain HTTP methods like POST. 
	RequestBody string `json:"request_body,omitempty"`
	// If HTTP response body match string (regular expressions not supported) is specified (using LbHttpMonitor.response_body) then the healthcheck HTTP response body is matched against the specified string and server is considered healthy only if there is a match. If the response body string is not specified, HTTP healthcheck is considered successful if the HTTP response status code is 2xx, but it can be configured to accept other status codes as successful. 
	ResponseBody string `json:"response_body,omitempty"`
	// supported SSL cipher list to servers
	Ciphers []string `json:"ciphers,omitempty"`
	// Array of HTTP request headers
	RequestHeaders []LbHttpRequestHeader `json:"request_headers,omitempty"`
	// client certificate can be specified to support client authentication. 
	ClientCertificateId string `json:"client_certificate_id,omitempty"`
	// the health check method for HTTP monitor type
	RequestMethod string `json:"request_method,omitempty"`
	// This flag is set to true when all the ciphers and protocols are FIPS compliant. It is set to false when one of the ciphers or protocols are not FIPS compliant.. 
	IsFips bool `json:"is_fips,omitempty"`
	// authentication depth is used to set the verification depth in the server certificates chain. 
	CertificateChainDepth int64 `json:"certificate_chain_depth,omitempty"`
	// This flag is set to true when all the ciphers and protocols are secure. It is set to false when one of the ciphers or protocols is insecure. 
	IsSecure bool `json:"is_secure,omitempty"`
	// URL used for HTTP monitor
	RequestUrl string `json:"request_url,omitempty"`
	// It is a label of cipher group which is mostly consumed by GUI. 
	CipherGroupLabel string `json:"cipher_group_label,omitempty"`
	// HTTP request version
	RequestVersion string `json:"request_version,omitempty"`
	// SSL versions TLS1.1 and TLS1.2 are supported and enabled by default. SSLv2, SSLv3, and TLS1.0 are supported, but disabled by default. 
	Protocols []string `json:"protocols,omitempty"`
}
