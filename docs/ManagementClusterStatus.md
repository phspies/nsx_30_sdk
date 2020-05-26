# ManagementClusterStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The current status of the management cluster | [optional] [default to null]
**OfflineNodes** | [**[]ManagementPlaneBaseNodeInfo**](ManagementPlaneBaseNodeInfo.md) | Current missing management plane nodes | [optional] [default to null]
**RequiredMembersForInitialization** | [**[]ClusterInitializationNodeInfo**](ClusterInitializationNodeInfo.md) | The details of the cluster nodes required for cluster initialization | [optional] [default to null]
**OnlineNodes** | [**[]ManagementPlaneBaseNodeInfo**](ManagementPlaneBaseNodeInfo.md) | Current alive management plane nodes | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

