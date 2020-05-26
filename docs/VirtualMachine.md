# VirtualMachine

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastSyncTime** | **int64** | Timestamp of last modification | [optional] [default to null]
**DisplayName** | **string** | Defaults to ID if not set | [optional] [default to null]
**Description** | **string** | Description of this resource | [optional] [default to null]
**ResourceType** | **string** | The type of this resource. | [default to null]
**Tags** | [**[]Tag**](Tag.md) | Opaque identifiers meaningful to the API user | [optional] [default to null]
**Source** | [***ResourceReference**](ResourceReference.md) |  | [optional] [default to null]
**LocalIdOnHost** | **string** | Id of the vm unique within the host. | [default to null]
**Type_** | **string** | Virtual Machine type; Edge, Service VM or other. | [optional] [default to null]
**GuestInfo** | [***GuestInfo**](GuestInfo.md) |  | [optional] [default to null]
**PowerState** | **string** | Current power state of this virtual machine in the system. | [default to null]
**ComputeIds** | **[]string** | List of external compute ids of the virtual machine in the format &#x27;id-type-key:value&#x27; , list of external compute ids [&#x27;uuid:xxxx-xxxx-xxxx-xxxx&#x27;, &#x27;moIdOnHost:moref-11&#x27;, &#x27;instanceUuid:xxxx-xxxx-xxxx-xxxx&#x27;] | [default to null]
**HostId** | **string** | Id of the host in which this virtual machine exists. | [optional] [default to null]
**ExternalId** | **string** | Current external id of this virtual machine in the system. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

