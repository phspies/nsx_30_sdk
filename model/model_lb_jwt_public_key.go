package nsxt

// The key is used to specify the public key content which is used to verify the signature of JWT tokens. 
type LbJwtPublicKey struct {
	// The property is used to identify JWT key type. 
	Type_ string `json:"type"`
	// Content of public key
	PublicKeyContent string `json:"public_key_content"`
}
