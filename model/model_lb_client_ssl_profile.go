package nsxt

type LbClientSslProfile struct {
	// SSL session caching allows SSL client and server to reuse previously negotiated security parameters avoiding the expensive public key operation during handshake. 
	SessionCacheEnabled bool `json:"session_cache_enabled,omitempty"`
	// Session cache timeout specifies how long the SSL session parameters are held on to and can be reused. 
	SessionCacheTimeout int64 `json:"session_cache_timeout,omitempty"`
	// It is a label of cipher group which is mostly consumed by GUI. 
	CipherGroupLabel string `json:"cipher_group_label,omitempty"`
	// This flag is set to true when all the ciphers and protocols are FIPS compliant. It is set to false when one of the ciphers or protocols are not FIPS compliant.. 
	IsFips bool `json:"is_fips,omitempty"`
	// This flag is set to true when all the ciphers and protocols are secure. It is set to false when one of the ciphers or protocols is insecure. 
	IsSecure bool `json:"is_secure,omitempty"`
	// During SSL handshake as part of the SSL client Hello client sends an ordered list of ciphers that it can support (or prefers) and typically server selects the first one from the top of that list it can also support. For Perfect Forward Secrecy(PFS), server could override the client's preference. 
	PreferServerCiphers bool `json:"prefer_server_ciphers,omitempty"`
	// supported SSL cipher list to client side
	Ciphers []string `json:"ciphers,omitempty"`
	// SSL versions TLS1.1 and TLS1.2 are supported and enabled by default. SSLv2, SSLv3, and TLS1.0 are supported, but disabled by default. 
	Protocols []string `json:"protocols,omitempty"`
}
