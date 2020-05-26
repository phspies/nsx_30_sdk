# Ipv4Header

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SrcIp** | **string** | The source ip address. | [optional] [default to null]
**Flags** | **int64** | IP flags | [optional] [default to 0]
**DstIp** | **string** | The destination ip address. | [optional] [default to null]
**SrcSubnetPrefixLen** | **int64** | This is used together with src_ip to calculate dst_ip for broadcast when dst_ip is not given; not used in all other cases. | [optional] [default to null]
**Ttl** | **int64** | Time to live (ttl) | [optional] [default to 64]
**Protocol** | **int64** | IP protocol - defaults to ICMP | [optional] [default to 1]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

