# SiMacLearningCounters

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MacsLearned** | **int64** | Number of MACs learned | [optional] [default to null]
**MacNotLearnedPacketsDropped** | **int64** | The number of packets with unknown source MAC address that are dropped without learning the source MAC address. Applicable only when the MAC limit is reached and MAC Limit policy is MAC_LEARNING_LIMIT_POLICY_DROP. | [optional] [default to null]
**MacNotLearnedPacketsAllowed** | **int64** | The number of packets with unknown source MAC address that are dispatched without learning the source MAC address. Applicable only when the MAC limit is reached and MAC Limit policy is MAC_LEARNING_LIMIT_POLICY_ALLOW. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

