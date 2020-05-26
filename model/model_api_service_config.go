package nsxt

// Properties that affect the configuration of the NSX API service. 
type ApiServiceConfig struct {
	// The maximum number of concurrent API requests that will be serviced. If the number of API requests being processed exceeds this limit, new API requests will be refused and a 503 Service Unavailable response will be returned to the client.  To disable API concurrency limiting, set this value to 0.
	GlobalApiConcurrencyLimit int64 `json:"global_api_concurrency_limit,omitempty"`
	// The list of IP addresses which are not subjected to a lockout on failed login attempts.
	LockoutImmuneAddresses []string `json:"lockout_immune_addresses,omitempty"`
	// Identifies whether basic authentication is enabled or disabled in API calls.
	BasicAuthenticationEnabled bool `json:"basic_authentication_enabled,omitempty"`
	// Host name or IP address to use for redirect location headers, or empty string to derive from current request. To disable, set redirect_host to the empty string (\"\").
	RedirectHost string `json:"redirect_host,omitempty"`
	// The TLS cipher suites that the API service will negotiate.
	CipherSuites []CipherSuite `json:"cipher_suites,omitempty"`
	// Identifies whether cookie-based authentication is enabled or disabled in API calls. When cookie-based authentication is disabled, new sessions cannot be created via /api/session/create.
	CookieBasedAuthenticationEnabled bool `json:"cookie_based_authentication_enabled,omitempty"`
	// NSX session inactivity timeout
	SessionTimeout int64 `json:"session_timeout,omitempty"`
	// The maximum number of API requests that will be serviced per second for a given authenticated client.  If more API requests are received than can be serviced, a 429 Too Many Requests HTTP response will be returned. To disable API rate limiting, set this value to 0.
	ClientApiRateLimit int64 `json:"client_api_rate_limit,omitempty"`
	// The maximum number of concurrent API requests that will be serviced for a given authenticated client.  If the number of API requests being processed exceeds this limit, new API requests will be refused and a 503 Service Unavailable response will be returned to the client. To disable API concurrency limiting, set this value to 0.
	ClientApiConcurrencyLimit int64 `json:"client_api_concurrency_limit,omitempty"`
	// The TLS protocol versions that the API service will negotiate.
	ProtocolVersions []ProtocolVersion `json:"protocol_versions,omitempty"`
	// NSX connection timeout, in seconds. To disable timeout, set to 0.
	ConnectionTimeout int64 `json:"connection_timeout,omitempty"`
}
