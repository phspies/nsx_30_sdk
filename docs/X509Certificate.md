# X509Certificate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EcdsaEcFieldF2mks** | **[]int64** | The order of the middle term(s) of the reduction polynomial in elliptic curve (EC) | characteristic 2 finite field.| Contents of this array are copied to protect against subsequent modification in ECDSA. | [optional] [default to null]
**Version** | **string** | Certificate version (default v1). | [optional] [default to null]
**IsCa** | **bool** | True if this is a CA certificate. | [optional] [default to null]
**SignatureAlgorithm** | **string** | The algorithm used by the Certificate Authority to sign the certificate. | [optional] [default to null]
**EcdsaPublicKeyA** | **string** | The first coefficient of this elliptic curve in ECDSA. | [optional] [default to null]
**RsaPublicKeyExponent** | **string** | An RSA public key is made up of the modulus and the public exponent. Exponent is a power number. | [optional] [default to null]
**EcdsaEcFieldF2mm** | **int64** | The first coefficient of this elliptic curve in elliptic curve (EC) | characteristic 2 finite field for ECDSA. | [optional] [default to null]
**IssuerCn** | **string** | The certificate issuer&#x27;s common name. | [optional] [default to null]
**SubjectCn** | **string** | The certificate owner&#x27;s common name. | [optional] [default to null]
**EcdsaPublicKeyOrder** | **string** | The order of generator G in ECDSA. | [optional] [default to null]
**EcdsaEcFieldF2mrp** | **string** | The value whose i-th bit corresponds to the i-th coefficient of the reduction polynomial | in elliptic curve (EC) characteristic 2 finite field for ECDSA. | [optional] [default to null]
**PublicKeyLength** | **int64** | Size measured in bits of the public/private keys used in a cryptographic algorithm. | [optional] [default to null]
**NotBefore** | **int64** | The time in epoch milliseconds at which the certificate becomes valid. | [optional] [default to null]
**EcdsaEcFieldF2pp** | **string** | The specified prime for the elliptic curve prime finite field in ECDSA. | [optional] [default to null]
**Issuer** | **string** | The certificate issuers complete distinguished name. | [optional] [default to null]
**EcdsaPublicKeyB** | **string** | The second coefficient of this elliptic curve in ECDSA. | [optional] [default to null]
**RsaPublicKeyModulus** | **string** | An RSA public key is made up of the modulus and the public exponent. Modulus is wrap around number. | [optional] [default to null]
**DsaPublicKeyY** | **string** | One of the DSA cryptogaphic algorithm&#x27;s strength parameters. | [optional] [default to null]
**EcdsaPublicKeyCofactor** | **int64** | The co-factor in ECDSA. | [optional] [default to null]
**NotAfter** | **int64** | The time in epoch milliseconds at which the certificate becomes invalid. | [optional] [default to null]
**DsaPublicKeyQ** | **string** | One of the DSA cryptogaphic algorithm&#x27;s strength parameters, sub-prime. | [optional] [default to null]
**DsaPublicKeyP** | **string** | One of the DSA cryptogaphic algorithm&#x27;s strength parameters, prime. | [optional] [default to null]
**EcdsaPublicKeyGeneratorY** | **string** | Y co-ordinate of G (the generator which is also known as the base point) in ECDSA. | [optional] [default to null]
**EcdsaPublicKeyGeneratorX** | **string** | X co-ordinate of G (the generator which is also known as the base point) in ECDSA. | [optional] [default to null]
**PublicKeyAlgo** | **string** | Cryptographic algorithm used by the public key for data encryption. | [optional] [default to null]
**IsValid** | **bool** | True if this certificate is valid. | [optional] [default to null]
**EcdsaPublicKeySeed** | **[]string** | The bytes used during curve generation for later validation in ECDSA.| Contents of this array are copied to protect against subsequent modification. | [optional] [default to null]
**Signature** | **string** | The signature value(the raw signature bits) used for signing and validate the cert. | [optional] [default to null]
**SerialNumber** | **string** | Certificate&#x27;s serial number. | [optional] [default to null]
**DsaPublicKeyG** | **string** | One of the DSA cryptogaphic algorithm&#x27;s strength parameters, base. | [optional] [default to null]
**Subject** | **string** | The certificate owners complete distinguished name. | [optional] [default to null]
**EcdsaEcField** | **string** | Represents an elliptic curve (EC) finite field in ECDSA. | [optional] [default to null]
**EcdsaCurveName** | **string** | The Curve name for the ECDSA certificate. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

