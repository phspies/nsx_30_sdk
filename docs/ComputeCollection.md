# ComputeCollection

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastSyncTime** | **int64** | Timestamp of last modification | [optional] [default to null]
**DisplayName** | **string** | Defaults to ID if not set | [optional] [default to null]
**Description** | **string** | Description of this resource | [optional] [default to null]
**ResourceType** | **string** | The type of this resource. | [default to null]
**Tags** | [**[]Tag**](Tag.md) | Opaque identifiers meaningful to the API user | [optional] [default to null]
**OriginId** | **string** | Id of the compute manager from where this Compute Collection was discovered | [optional] [default to null]
**OriginProperties** | [**[]KeyValuePair**](KeyValuePair.md) | Key-Value map of additional specific properties of compute collection in the Compute Manager  | [optional] [default to null]
**ExternalId** | **string** | External ID of the ComputeCollection in the source Compute manager, e.g. mo-ref in VC  | [optional] [default to null]
**OwnerId** | **string** | Id of the owner of compute collection in the Compute Manager | [optional] [default to null]
**OriginType** | **string** | ComputeCollection type like VC_Cluster. Here the Compute Manager type prefix would help in differentiating similar named Compute Collection types from different Compute Managers  | [optional] [default to null]
**CmLocalId** | **string** | Local Id of the compute collection in the Compute Manager | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

