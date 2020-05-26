# LbServiceStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pools** | [**[]LbPoolStatus**](LbPoolStatus.md) | status of load balancer pools | [optional] [default to null]
**CpuUsage** | **int64** | Cpu usage in percentage | [optional] [default to null]
**ActiveTransportNodes** | **[]string** | Ids of load balancer service related active transport nodes | [optional] [default to null]
**MemoryUsage** | **int64** | Memory usage in percentage | [optional] [default to null]
**ServiceId** | **string** | Load balancer service identifier | [default to null]
**LastUpdateTimestamp** | **int64** | Timestamp when the data was last updated | [optional] [default to null]
**StandbyTransportNodes** | **[]string** | Ids of load balancer service related standby transport nodes | [optional] [default to null]
**ErrorMessage** | **string** | Error message, if available | [optional] [default to null]
**VirtualServers** | [**[]LbVirtualServerStatus**](LbVirtualServerStatus.md) | status of load balancer virtual servers | [optional] [default to null]
**ServiceStatus** | **string** | UP means the load balancer service is working fine on both transport-nodes(if have); DOWN means the load balancer service is down on both transport-nodes (if have), hence the load balancer will not respond to any requests; ERROR means error happens on transport-node(s) or no status is reported from transport-node(s). The load balancer service may be working (or not working); NO_STANDBY means load balancer service is working in one of the transport node while not in the other transport-node (if have). Hence if the load balancer service in the working transport-node goes down, the load balancer service will go down; DETACHED means that the load balancer service has no attachment setting and is not instantiated in any transport nodes; DISABLED means that admin state of load balancer service is DISABLED; UNKNOWN means that no status reported from transport-nodes.The load balancer service may be working(or not working).  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

