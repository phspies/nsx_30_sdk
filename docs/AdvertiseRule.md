# AdvertiseRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | ALLOW action enables the advertisment and DENY action disables the advertisement of a filtered routes to the connected TIER0 router. | [optional] [default to ACTION.ALLOW]
**RuleFilter** | [***AdvertisementRuleFilter**](AdvertisementRuleFilter.md) |  | [optional] [default to null]
**DisplayName** | **string** | Display name | [optional] [default to null]
**Networks** | **[]string** | network(CIDR) to be routed | [default to null]
**Description** | **string** | Description | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

