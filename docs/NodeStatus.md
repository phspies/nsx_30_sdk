# NodeStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MpaConnectivityStatus** | **string** | Indicates the fabric node&#x27;s MP&amp;lt;-&amp;gt;MPA channel connectivity status, UP, DOWN, UNKNOWN. | [optional] [default to null]
**LcpConnectivityStatusDetails** | [**[]ControlConnStatus**](ControlConnStatus.md) | Details, if any, about the current LCP&amp;lt;-&amp;gt;CCP channel connectivity status of the fabric node. | [optional] [default to null]
**MpaConnectivityStatusDetails** | **string** | Details, if any, about the current MP&amp;lt;-&amp;gt;MPA channel connectivity status of the fabric node. | [optional] [default to null]
**ExternalId** | **string** | HostNode external id | [optional] [default to null]
**SoftwareVersion** | **string** | Software version of the fabric node. | [optional] [default to null]
**MaintenanceMode** | **string** | Indicates the fabric node&#x27;s status of maintenance mode, OFF, ENTERING, ON, EXITING. | [optional] [default to null]
**InventorySyncPaused** | **bool** | Is true if inventory sync is paused else false | [optional] [default to null]
**SystemStatus** | [***NodeStatusProperties**](NodeStatusProperties.md) |  | [optional] [default to null]
**InventorySyncReenableTime** | **int64** | Inventory sync auto re-enable target time, in epoch milis | [optional] [default to null]
**LcpConnectivityStatus** | **string** | Indicates the fabric node&#x27;s LCP&amp;lt;-&amp;gt;CCP channel connectivity status, UP, DOWN, DEGRADED, UNKNOWN. | [optional] [default to LCP_CONNECTIVITY_STATUS.UNKNOWN]
**LastHeartbeatTimestamp** | **int64** | Timestamp of the last heartbeat status change, in epoch milliseconds. | [optional] [default to null]
**LastSyncTime** | **int64** | Timestamp of the last successful update of Inventory, in epoch milliseconds. | [optional] [default to null]
**HostNodeDeploymentStatus** | **string** | This enum specifies the current nsx install state for host node or current deployment and ready state for edge node. The ready status &#x27;NODE_READY&#x27; indicates whether edge node is ready to become a transport node. The status &#x27;EDGE_CONFIG_ERROR&#x27; indicates that edge hardware or underlying host is not supported. After all fabric level operations are done for an edge node, this value indicates transport node related configuration issues and state as relevant.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

