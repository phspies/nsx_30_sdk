package nsxt

// The key is used to specify the symmetric key which is used to verify the signature of JWT tokens. 
type LbJwtSymmetricKey struct {
	// The property is used to identify JWT key type. 
	Type_ string `json:"type"`
}
