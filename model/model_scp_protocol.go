package nsxt

type ScpProtocol struct {
	// Protocol name
	Name string `json:"name"`
	// SSH fingerprint of server
	SshFingerprint string `json:"ssh_fingerprint"`
	AuthenticationScheme *PasswordAuthenticationScheme `json:"authentication_scheme"`
}
