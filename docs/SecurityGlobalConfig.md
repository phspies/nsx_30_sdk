# SecurityGlobalConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | Valid Global configuration types | [default to null]
**CaSignedOnly** | **bool** | When this flag is set to true (for NDcPP compliance) only ca-signed certificates will be allowed to be applied as server certificates. | [optional] [default to false]
**CrlCheckingEnabled** | **bool** | When this flag is set to true, during certificate checking the CRL is fetched and checked whether the certificate is revoked or not. | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

