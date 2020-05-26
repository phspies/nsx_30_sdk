package nsxt

// HTTP Service properties
type HttpServiceProperties struct {
	// The maximum number of concurrent API requests that will be serviced. If the number of API requests being processed exceeds this limit, new API requests will be refused and a 503 Service Unavailable response will be returned to the client.  To disable API concurrency limiting, set this value to 0.
	GlobalApiConcurrencyLimit int64 `json:"global_api_concurrency_limit,omitempty"`
	Certificate *Certificate `json:"certificate,omitempty"`
	// Identifies whether basic authentication is enabled or disabled in API calls.
	BasicAuthenticationEnabled bool `json:"basic_authentication_enabled,omitempty"`
	// Identifies whether cookie-based authentication is enabled or disabled in API calls. When cookie-based authentication is disabled, new sessions cannot be created via /api/session/create.
	CookieBasedAuthenticationEnabled bool `json:"cookie_based_authentication_enabled,omitempty"`
	// Cipher suites used to secure contents of connection
	CipherSuites []CipherSuite `json:"cipher_suites,omitempty"`
	// Host name or IP address to use for redirect location headers, or empty string to derive from current request
	RedirectHost string `json:"redirect_host,omitempty"`
	// NSX session inactivity timeout, set to 0 to configure no timeout
	SessionTimeout int64 `json:"session_timeout,omitempty"`
	// The maximum number of API requests that will be serviced per second for a given authenticated client.  If more API requests are received than can be serviced, a 429 Too Many Requests HTTP response will be returned. To disable API rate limiting, set this value to 0.
	ClientApiRateLimit int64 `json:"client_api_rate_limit,omitempty"`
	// The maximum number of concurrent API requests that will be serviced for a given authenticated client.  If the number of API requests being processed exceeds this limit, new API requests will be refused and a 503 Service Unavailable response will be returned to the client. To disable API concurrency limiting, set this value to 0.
	ClientApiConcurrencyLimit int64 `json:"client_api_concurrency_limit,omitempty"`
	// TLS protocol versions
	ProtocolVersions []ProtocolVersion `json:"protocol_versions,omitempty"`
	// NSX connection timeout, set to 0 to configure no timeout
	ConnectionTimeout int64 `json:"connection_timeout,omitempty"`
	// Service logging level
	LoggingLevel string `json:"logging_level,omitempty"`
}
