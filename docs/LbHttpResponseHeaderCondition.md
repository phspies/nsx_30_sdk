# LbHttpResponseHeaderCondition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inverse** | **bool** | A flag to indicate whether reverse the match result of this condition | [optional] [default to false]
**Type_** | **string** | Type of load balancer rule condition | [default to null]
**HeaderValue** | **string** | Value of HTTP header field | [default to null]
**CaseSensitive** | **bool** | If true, case is significant when comparing HTTP header value.  | [optional] [default to true]
**MatchType** | **string** | Match type of HTTP header value | [optional] [default to MATCH_TYPE.REGEX]
**HeaderName** | **string** | Name of HTTP header field | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

