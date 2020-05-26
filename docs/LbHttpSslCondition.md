# LbHttpSslCondition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inverse** | **bool** | A flag to indicate whether reverse the match result of this condition | [optional] [default to false]
**Type_** | **string** | Type of load balancer rule condition | [default to null]
**ClientSupportedSslCiphers** | **[]string** | Cipher list which supported by client | [optional] [default to null]
**ClientCertificateIssuerDn** | [***LbClientCertificateIssuerDnCondition**](LbClientCertificateIssuerDnCondition.md) |  | [optional] [default to null]
**ClientCertificateSubjectDn** | [***LbClientCertificateSubjectDnCondition**](LbClientCertificateSubjectDnCondition.md) |  | [optional] [default to null]
**UsedSslCipher** | **string** | Cipher used for an established SSL connection | [optional] [default to null]
**SessionReused** | **string** | The type of SSL session reused | [optional] [default to SESSION_REUSED.IGNORE]
**UsedProtocol** | **string** | Protocol of an established SSL connection | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

