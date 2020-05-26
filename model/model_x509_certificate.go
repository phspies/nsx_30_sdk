package nsxt

type X509Certificate struct {
	// The order of the middle term(s) of the reduction polynomial in elliptic curve (EC) | characteristic 2 finite field.| Contents of this array are copied to protect against subsequent modification in ECDSA.
	EcdsaEcFieldF2mks []int64 `json:"ecdsa_ec_field_f2mks,omitempty"`
	// Certificate version (default v1).
	Version string `json:"version,omitempty"`
	// True if this is a CA certificate.
	IsCa bool `json:"is_ca,omitempty"`
	// The algorithm used by the Certificate Authority to sign the certificate.
	SignatureAlgorithm string `json:"signature_algorithm,omitempty"`
	// The first coefficient of this elliptic curve in ECDSA.
	EcdsaPublicKeyA string `json:"ecdsa_public_key_a,omitempty"`
	// An RSA public key is made up of the modulus and the public exponent. Exponent is a power number.
	RsaPublicKeyExponent string `json:"rsa_public_key_exponent,omitempty"`
	// The first coefficient of this elliptic curve in elliptic curve (EC) | characteristic 2 finite field for ECDSA.
	EcdsaEcFieldF2mm int64 `json:"ecdsa_ec_field_f2mm,omitempty"`
	// The certificate issuer's common name.
	IssuerCn string `json:"issuer_cn,omitempty"`
	// The certificate owner's common name.
	SubjectCn string `json:"subject_cn,omitempty"`
	// The order of generator G in ECDSA.
	EcdsaPublicKeyOrder string `json:"ecdsa_public_key_order,omitempty"`
	// The value whose i-th bit corresponds to the i-th coefficient of the reduction polynomial | in elliptic curve (EC) characteristic 2 finite field for ECDSA.
	EcdsaEcFieldF2mrp string `json:"ecdsa_ec_field_f2mrp,omitempty"`
	// Size measured in bits of the public/private keys used in a cryptographic algorithm.
	PublicKeyLength int64 `json:"public_key_length,omitempty"`
	// The time in epoch milliseconds at which the certificate becomes valid.
	NotBefore int64 `json:"not_before,omitempty"`
	// The specified prime for the elliptic curve prime finite field in ECDSA.
	EcdsaEcFieldF2pp string `json:"ecdsa_ec_field_f2pp,omitempty"`
	// The certificate issuers complete distinguished name.
	Issuer string `json:"issuer,omitempty"`
	// The second coefficient of this elliptic curve in ECDSA.
	EcdsaPublicKeyB string `json:"ecdsa_public_key_b,omitempty"`
	// An RSA public key is made up of the modulus and the public exponent. Modulus is wrap around number.
	RsaPublicKeyModulus string `json:"rsa_public_key_modulus,omitempty"`
	// One of the DSA cryptogaphic algorithm's strength parameters.
	DsaPublicKeyY string `json:"dsa_public_key_y,omitempty"`
	// The co-factor in ECDSA.
	EcdsaPublicKeyCofactor int64 `json:"ecdsa_public_key_cofactor,omitempty"`
	// The time in epoch milliseconds at which the certificate becomes invalid.
	NotAfter int64 `json:"not_after,omitempty"`
	// One of the DSA cryptogaphic algorithm's strength parameters, sub-prime.
	DsaPublicKeyQ string `json:"dsa_public_key_q,omitempty"`
	// One of the DSA cryptogaphic algorithm's strength parameters, prime.
	DsaPublicKeyP string `json:"dsa_public_key_p,omitempty"`
	// Y co-ordinate of G (the generator which is also known as the base point) in ECDSA.
	EcdsaPublicKeyGeneratorY string `json:"ecdsa_public_key_generator_y,omitempty"`
	// X co-ordinate of G (the generator which is also known as the base point) in ECDSA.
	EcdsaPublicKeyGeneratorX string `json:"ecdsa_public_key_generator_x,omitempty"`
	// Cryptographic algorithm used by the public key for data encryption.
	PublicKeyAlgo string `json:"public_key_algo,omitempty"`
	// True if this certificate is valid.
	IsValid bool `json:"is_valid,omitempty"`
	// The bytes used during curve generation for later validation in ECDSA.| Contents of this array are copied to protect against subsequent modification.
	EcdsaPublicKeySeed []string `json:"ecdsa_public_key_seed,omitempty"`
	// The signature value(the raw signature bits) used for signing and validate the cert.
	Signature string `json:"signature,omitempty"`
	// Certificate's serial number.
	SerialNumber string `json:"serial_number,omitempty"`
	// One of the DSA cryptogaphic algorithm's strength parameters, base.
	DsaPublicKeyG string `json:"dsa_public_key_g,omitempty"`
	// The certificate owners complete distinguished name.
	Subject string `json:"subject,omitempty"`
	// Represents an elliptic curve (EC) finite field in ECDSA.
	EcdsaEcField string `json:"ecdsa_ec_field,omitempty"`
	// The Curve name for the ECDSA certificate.
	EcdsaCurveName string `json:"ecdsa_curve_name,omitempty"`
}
