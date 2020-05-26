# IPv6DhcpServer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpServerIp** | **string** | DHCP server ip in CIDR format. | [optional] [default to null]
**ServerId** | **string** | DHCP server id. | [optional] [default to null]
**DnsNameservers** | **[]string** | Primary and secondary DNS server address to assign host. They can be overridden by ip-pool or static-binding level property.  | [optional] [default to null]
**SntpServers** | **[]string** | SNTP server ips. | [optional] [default to null]
**DomainNames** | **[]string** | Host name or prefix to be assigned to host. It can be overridden by ip-pool or static-binding level property.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

