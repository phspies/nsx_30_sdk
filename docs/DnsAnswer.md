# DnsAnswer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsServer** | **string** | Dns server ip address and port, format is \&quot;ip address#port\&quot;.  | [default to null]
**RawAnswer** | **string** | It can be NXDOMAIN or error message which is not consisted of authoritative_answer or non_authoritative_answer.  | [optional] [default to null]
**NonAuthoritativeAnswers** | [**[]DnsQueryAnswer**](DnsQueryAnswer.md) | Non-authotitative answers of the query. This is a deprecated property, please use &#x27;answers&#x27; instead.  | [optional] [default to null]
**SourceIp** | **string** | The source ip used in this lookup.  | [default to null]
**EdgeNodeId** | **string** | ID of the edge node that performed the query.  | [default to null]
**AuthoritativeAnswers** | [**[]DnsQueryAnswer**](DnsQueryAnswer.md) | Authotitative answers of the query. This is a deprecated property, please use &#x27;answers&#x27; instead.  | [optional] [default to null]
**Answers** | [**[]DnsQueryAnswer**](DnsQueryAnswer.md) | The answers of the query.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

