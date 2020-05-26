# EdgeClusterMemberAllocationProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocationPool** | [***EdgeClusterMemberAllocationPool**](EdgeClusterMemberAllocationPool.md) |  | [optional] [default to null]
**EnableStandbyRelocation** | **bool** | Flag to enable the auto-relocation of standby service router running on edge cluster and node associated with the logical router. Only dynamically allocated tier1 logical routers are considered for the relocation.  | [optional] [default to false]
**AllocationType** | **string** | Allocation type is used to specify the mode used to allocate the LR. This is populated only for TIER1 logical router and for TIER0 this will be null.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
