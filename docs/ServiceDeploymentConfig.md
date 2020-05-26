# ServiceDeploymentConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageId** | **string** | Moref of the datastore in VC. If it is to be taken from &#x27;Agent VM Settings&#x27;, then it should be empty. | [optional] [default to null]
**HostId** | **string** | The service VM will be deployed on the specified host in the specified server within the cluster if host_id is specified. Note: You must ensure that storage and specified networks are accessible       by this host.  | [optional] [default to null]
**ComputeCollectionId** | **string** | Resource Pool or cluster Id. | [default to null]
**VmNicInfo** | [***VmNicInfo**](VmNicInfo.md) |  | [optional] [default to null]
**ComputeManagerId** | **string** | Context Id or VCenter Id. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

