# IPv4DhcpServer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | [***DhcpOptions**](DhcpOptions.md) |  | [optional] [default to null]
**MonitorIppoolUsage** | **bool** | Enable or disable monitoring of DHCP ip-pools usage. When enabled, system events are generated when pool usage exceeds the configured thresholds. System events can be viewed in REST API /api/v2/hpm/alarms  | [optional] [default to false]
**DhcpServerIp** | **string** | DHCP server ip in CIDR format. | [default to null]
**DnsNameservers** | **[]string** | Primary and secondary DNS server address to assign host. They can be overridden by ip-pool or static-binding level property.  | [optional] [default to null]
**DomainName** | **string** | Host name or prefix to be assigned to host. It can be overridden by ip-pool or static-binding level property.  | [optional] [default to null]
**GatewayIp** | **string** | Gateway ip to be assigned to host. It can be overridden by ip-pool or static-binding level property.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

