# FirewallFloodProtectionProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | Resource type to use as profile type | [default to null]
**IcmpActiveFlowLimit** | **int64** | The maximum limit of active icmp connections. If this property is omitted, or set to null, then there is no limit on active icmp connections for those components if it&#x27;s applied to ESX components (such as segment, segment port, virtual machine, etc); on the other side, if it&#x27;s applied to EDGE components (such as, gateway), it will be set to default limit (10,000) on the specific components. | [optional] [default to null]
**OtherActiveConnLimit** | **int64** | The maximum limit of other active connections besides udp, icmp and half open tcp connections. If this property is omitted, or set to null, then there is no limit on other active connections besides udp, icmp and tcp half open connections for those components if it&#x27;s applied to ESX components (such as segment, segment port, virtual machine, etc); on the other side, if it&#x27;s applied to EDGE components (such as, gateway), it will be set to default limit (10,000) on the specific components. | [optional] [default to null]
**EnableSyncache** | **bool** | The flag to indicate syncache is enabled or not. This option does not apply to EDGE components. | [optional] [default to false]
**UdpActiveFlowLimit** | **int64** | The maximum limit of active udp connections. If this property is omitted, or set to null, then there is no limit on active udp connections for those components if it&#x27;s applied to ESX components (such as segment, segment port, virtual machine, etc); on the other side, if it&#x27;s applied to EDGE components (such as, gateway), it will be set to default limit (100,000) on the specific component. | [optional] [default to null]
**TcpHalfOpenConnLimit** | **int64** | The maximum limit of tcp half open connections. If this property is omitted, or set to null, then there is no limit on active tcp half open connections for those components if it&#x27;s applied to ESX components (such as segment, segment port, virtual machine, etc); on the other side, if it&#x27;s applied to EDGE components (such as, gateway), it will be set to default limit (1,000,000) on the specific components. | [optional] [default to null]
**EnableRstSpoofing** | **bool** | The flag to indicate RST spoofing is enabled or not. This option does not apply to EDGE components. This can be enabled only if syncache is enabled. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

