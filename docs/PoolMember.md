# PoolMember

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxConcurrentConnections** | **int64** | To ensure members are not overloaded, connections to a member can be capped by the load balancer. When a member reaches this limit, it is skipped during server selection. If it is not specified, it means that connections are unlimited.  | [optional] [default to null]
**AdminState** | **string** | member admin state | [optional] [default to ADMIN_STATE.ENABLED]
**BackupMember** | **bool** | Backup servers are typically configured with a sorry page indicating to the user that the application is currently unavailable. While the pool is active (a specified minimum number of pool members are active) BACKUP members are skipped during server selection. When the pool is inactive, incoming connections are sent to only the BACKUP member(s).  | [optional] [default to false]
**Weight** | **int64** | Pool member weight is used for WEIGHTED_ROUND_ROBIN balancing algorithm. The weight value would be ignored in other algorithms.  | [optional] [default to 1]
**DisplayName** | **string** | pool member name | [optional] [default to null]
**IpAddress** | **string** | pool member IP address | [default to null]
**Port** | **string** | If port is specified, all connections will be sent to this port. Only single port is supported. If unset, the same port the client connected to will be used, it could be overrode by default_pool_member_port setting in virtual server. The port should not specified for port range case.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

