package nsxt

type SupportBundleFileTransferProtocol struct {
	// SSH fingerprint of server
	SshFingerprint string `json:"ssh_fingerprint"`
	// Protocol name
	Name string `json:"name"`
	AuthenticationScheme *SupportBundleFileTransferAuthenticationScheme `json:"authentication_scheme"`
}
