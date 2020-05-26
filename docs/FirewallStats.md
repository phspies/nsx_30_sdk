# FirewallStats

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [***SelfResourceLink**](SelfResourceLink.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](ResourceLink.md) | The server will populate this field when returing the resource. Ignored on PUT and POST. | [optional] [default to null]
**Schema** | **string** | Schema for this resource | [optional] [default to null]
**TotalSessionCount** | **int64** | Aggregated number of sessions processed by the all firewall rules. This is aggregated statistic which are computed with lower frequency compared to individual generic rule statistics. It may have a computation delay up to 15 minutes in response to this API. | [optional] [default to null]
**PacketCount** | **int64** | Aggregated number of packets processed by the rule. | [optional] [default to null]
**PopularityIndex** | **int64** | This is calculated by sessions count divided by age of the rule. | [optional] [default to null]
**MaxSessionCount** | **int64** | Maximum value of sessions count of all firewall rules of the type. This is aggregated statistic which are computed with lower frequency compared to generic rule statistics. It may have a computation delay up to 15 minutes in response to this API. | [optional] [default to null]
**ByteCount** | **int64** | Aggregated number of bytes processed by the rule. | [optional] [default to null]
**MaxPopularityIndex** | **int64** | Maximum value of popularity index of all firewall rules of the type. This is aggregated statistic which are computed with lower frequency compared to individual generic rule statistics. It may have a computation delay up to 15 minutes in response to this API. | [optional] [default to null]
**HitCount** | **int64** | Aggregated number of hits received by the rule. | [optional] [default to null]
**SessionCount** | **int64** | Aggregated number of sessions processed by the rule. | [optional] [default to null]
**RuleId** | **string** | Rule Identifier of the Firewall rule. This is a globally unique number. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

