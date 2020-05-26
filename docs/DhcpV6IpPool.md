# DhcpV6IpPool

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LeaseTime** | **int64** | Lease time, in seconds. | [optional] [default to 86400]
**SntpServers** | **[]string** | SNTP server ips. | [optional] [default to null]
**PreferredTime** | **int64** | Preferred time, in seconds. If this value is not provided, the value of lease_time*0.8 will be used.  | [optional] [default to null]
**DnsNameservers** | **[]string** | Primary and secondary DNS server address to assign host. They can be overridden by ip-pool or static-binding level property.  | [optional] [default to null]
**DomainNames** | **[]string** | Host name or prefix to be assigned to host. It can be overridden by ip-pool or static-binding level property.  | [optional] [default to null]
**Ranges** | [**[]IpPoolRange**](IpPoolRange.md) | Ip address ranges to define dynamic ip allocation ranges. | [optional] [default to null]
**ExcludedRanges** | [**[]IpPoolRange**](IpPoolRange.md) | Excluded addresses to define dynamic ip allocation ranges. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

