# Ipv6Header

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SrcIp** | **string** | The source ip address. | [optional] [default to null]
**DstIp** | **string** | The destination ip address. | [optional] [default to null]
**NextHeader** | **int64** | Identifies the type of header immediately following the IPv6 header. | [optional] [default to 58]
**HopLimit** | **int64** | Decremented by 1 by each node that forwards the packets. The packet is discarded if Hop Limit is decremented to zero. | [optional] [default to 64]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

