# DhcpStaticBinding

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LeaseTime** | **int64** | Lease time, in seconds, [60-(2^32-1)]. Default is 86400. | [optional] [default to 86400]
**GatewayIp** | **string** | Gateway ip address of the allocation. | [optional] [default to null]
**Options** | [***DhcpOptions**](DhcpOptions.md) |  | [optional] [default to null]
**IpAddress** | **string** | The ip address to be assigned to the host. | [default to null]
**HostName** | **string** | The host name to be assigned to the host. | [optional] [default to null]
**MacAddress** | **string** | The MAC address of the host. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

