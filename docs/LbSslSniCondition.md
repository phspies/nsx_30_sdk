# LbSslSniCondition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inverse** | **bool** | A flag to indicate whether reverse the match result of this condition | [optional] [default to false]
**Type_** | **string** | Type of load balancer rule condition | [default to null]
**CaseSensitive** | **bool** | If true, case is significant when comparing SNI value.  | [optional] [default to true]
**MatchType** | **string** | Determine how a specified string value is used to match SNI.  | [optional] [default to MATCH_TYPE.REGEX]
**Sni** | **string** | The SNI(Server Name indication) in client hello message.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

