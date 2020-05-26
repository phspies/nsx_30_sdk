# VmToolsInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastSyncTime** | **int64** | Timestamp of last modification | [optional] [default to null]
**DisplayName** | **string** | Defaults to ID if not set | [optional] [default to null]
**Description** | **string** | Description of this resource | [optional] [default to null]
**ResourceType** | **string** | The type of this resource. | [default to null]
**Tags** | [**[]Tag**](Tag.md) | Opaque identifiers meaningful to the API user | [optional] [default to null]
**Source** | [***ResourceReference**](ResourceReference.md) |  | [optional] [default to null]
**VmType** | **string** | Type of VM - Edge, Service or other. | [optional] [default to null]
**NetworkAgentVersion** | **string** | Version of network agent on the VM of a third party partner solution. | [optional] [default to null]
**HostLocalId** | **string** | Id of the VM which is assigned locally by the host. It is the VM-moref on ESXi hosts, in other environments it is VM UUID. | [optional] [default to null]
**ExternalId** | **string** | Current external id of this virtual machine in the system. | [optional] [default to null]
**ToolsVersion** | **string** | Version of VMTools installed on the VM. | [optional] [default to null]
**FileAgentVersion** | **string** | Version of file agent on the VM of a third party partner solution. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

