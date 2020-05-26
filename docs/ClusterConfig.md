# ClusterConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **int32** | The _revision property describes the current revision of the resource. To prevent clients from overwriting each other&#x27;s changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected. | [optional] [default to null]
**ControlClusterChangesAllowed** | **bool** | True if control cluster nodes may be added or removed | [optional] [default to null]
**Nodes** | [**[]ClusterNodeInfo**](ClusterNodeInfo.md) | Configuration of each node in cluster | [optional] [default to null]
**MgmtClusterChangesAllowed** | **bool** | True if management cluster nodes may be added or removed | [optional] [default to null]
**ClusterId** | **string** | Unique identifier of this cluster | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

