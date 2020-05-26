# IpfixSwitchConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedTos** | [**[]ResourceReference**](ResourceReference.md) | List of objects where the IPFIX Config will be enabled. | [optional] [default to null]
**ResourceType** | **string** | Supported IPFIX Config Types. | [default to null]
**ActiveTimeout** | **int32** | The time in seconds after a Flow is expired even if more packets matching this Flow are received by the cache.  | [optional] [default to 300]
**IdleTimeout** | **int32** | The time in seconds after a Flow is expired if no more packets matching this Flow are received by the cache.  | [optional] [default to 300]
**PacketSampleProbability** | **float64** | The probability in percentage that a packet is sampled. The value should be  in range (0,100] and can only have three decimal places at most. The probability  is equal for every packet.  | [optional] [default to null]
**MaxFlows** | **int64** | The maximum number of flow entries in each exporter flow cache.  | [optional] [default to 16384]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

