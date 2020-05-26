# DhcpIpPool

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LeaseTime** | **int64** | Lease time, in seconds, [60-(2^32-1)]. Default is 86400. | [optional] [default to 86400]
**GatewayIp** | **string** | Gateway ip address of the allocation. | [optional] [default to null]
**Options** | [***DhcpOptions**](DhcpOptions.md) |  | [optional] [default to null]
**AllocationRanges** | [**[]IpPoolRange**](IpPoolRange.md) | Ip-ranges to define dynamic ip allocation ranges. | [default to null]
**WarningThreshold** | **int64** | Warning threshold. Alert will be raised if the pool usage reaches the given threshold.  | [optional] [default to 80]
**ErrorThreshold** | **int64** | Error threshold. Alert will be raised if the pool usage reaches the given threshold.  | [optional] [default to 100]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

