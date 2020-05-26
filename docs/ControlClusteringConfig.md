# ControlClusteringConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusteringType** | **string** | Specifies the type of clustering config to be used.  | [default to null]
**JoinToExistingCluster** | **bool** | Specifies whether or not the cluster node VM should try to join to the existing control cluster or initialize a new one. Only required in uncertainty case, i.e. when there are manually- deployed controllers that are registered but not connected to the cluster and no auto-deployed controllers are part of the cluster.  | [optional] [default to null]
**SharedSecret** | **string** | Shared secret to be used when joining the cluster node VM to a control cluster or for initializing a new cluster with the VM. Must contain at least 4 unique characters and be at least 6 characters long.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

