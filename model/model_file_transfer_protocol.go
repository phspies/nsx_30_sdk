package nsxt

// Protocol to transfer backup file to remote server
type FileTransferProtocol struct {
	// Protocol name
	ProtocolName string `json:"protocol_name"`
	// The expected SSH fingerprint of the server. If the server's fingerprint does not match this fingerprint, the connection will be terminated.  Only ECDSA fingerprints hashed with SHA256 are supported. To obtain the host's ssh fingerprint, you should connect via some method other than SSH to obtain this information. You can use one of these commands to view the key's fingerprint: 1. ssh-keygen -l -E sha256 -f ssh_host_ecdsa_key.pub 2. awk '{print $2}' ssh_host_ecdsa_key.pub | base64 -d | sha256sum -b |    sed 's/ .*$//' | xxd -r -p | base64 | sed 's/.//44g' |    awk '{print \"SHA256:\"$1}' 
	SshFingerprint string `json:"ssh_fingerprint"`
	AuthenticationScheme *FileTransferAuthenticationScheme `json:"authentication_scheme"`
}
