# InstanceDeploymentConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextId** | **string** | Context Id or VCenter Id. | [default to null]
**VmNicInfos** | [**[]VmNicInfo**](VmNicInfo.md) | List of NIC information for VMs | [default to null]
**StorageId** | **string** | Storage Id. | [default to null]
**HostId** | **string** | The service VM will be deployed on the specified host in the specified server within the cluster if host_id is specified. Note: You must ensure that storage and specified networks are accessible by this host.  | [optional] [default to null]
**ComputeId** | **string** | Resource Pool or Compute Id. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

