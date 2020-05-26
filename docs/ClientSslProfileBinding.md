# ClientSslProfileBinding

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientAuth** | **string** | client authentication mode | [optional] [default to CLIENT_AUTH.IGNORE]
**SslProfileId** | **string** | Client SSL profile defines reusable, application-independent client side SSL properties.  | [optional] [default to null]
**CertificateChainDepth** | **int64** | authentication depth is used to set the verification depth in the client certificates chain.  | [optional] [default to 3]
**ClientAuthCaIds** | **[]string** | If client auth type is REQUIRED, client certificate must be signed by one of the trusted Certificate Authorities (CAs), also referred to as root CAs, whose self signed certificates are specified.  | [optional] [default to null]
**DefaultCertificateId** | **string** | A default certificate should be specified which will be used if the server does not host multiple hostnames on the same IP address or if the client does not support SNI extension.  | [default to null]
**SniCertificateIds** | **[]string** | Client-side SSL profile binding allows multiple certificates, for different hostnames, to be bound to the same virtual server.  | [optional] [default to null]
**ClientAuthCrlIds** | **[]string** | A Certificate Revocation List (CRL) can be specified in the client-side SSL profile binding to disallow compromised client certificates.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

