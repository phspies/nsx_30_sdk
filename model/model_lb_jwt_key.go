package nsxt

// LbJwtKey specifies the symmetric key or asymmetric public key used to decrypt the data in JWT. 
type LbJwtKey struct {
	// The property is used to identify JWT key type. 
	Type_ string `json:"type"`
}
