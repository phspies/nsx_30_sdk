# LbIpHeaderCondition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inverse** | **bool** | A flag to indicate whether reverse the match result of this condition | [optional] [default to false]
**Type_** | **string** | Type of load balancer rule condition | [default to null]
**GroupId** | **string** | Source IP address of HTTP message should match IP addresses which are configured in Group in order to perform actions.  | [optional] [default to null]
**SourceAddress** | **string** | Source IP address of HTTP message. IP Address can be expressed as a single IP address like 10.1.1.1, or a range of IP addresses like 10.1.1.101-10.1.1.160. Both IPv4 and IPv6 addresses are supported.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

