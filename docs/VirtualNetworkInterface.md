# VirtualNetworkInterface

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastSyncTime** | **int64** | Timestamp of last modification | [optional] [default to null]
**DisplayName** | **string** | Defaults to ID if not set | [optional] [default to null]
**Description** | **string** | Description of this resource | [optional] [default to null]
**ResourceType** | **string** | The type of this resource. | [default to null]
**Tags** | [**[]Tag**](Tag.md) | Opaque identifiers meaningful to the API user | [optional] [default to null]
**MacAddress** | **string** | MAC address of the virtual network interface. | [default to null]
**OwnerVmType** | **string** | Owner virtual machine type; Edge, Service VM or other. | [optional] [default to null]
**DeviceKey** | **string** | Device key of the virtual network interface. | [default to null]
**HostId** | **string** | Id of the host on which the vm exists. | [default to null]
**OwnerVmId** | **string** | Id of the vm to which this virtual network interface belongs. | [default to null]
**VmLocalIdOnHost** | **string** | Id of the vm unique within the host. | [default to null]
**ExternalId** | **string** | External Id of the virtual network inferface. | [default to null]
**LportAttachmentId** | **string** | LPort Attachment Id of the virtual network interface. | [optional] [default to null]
**IpAddressInfo** | [**[]IpAddressInfo**](IpAddressInfo.md) | IP Addresses of the the virtual network interface, from various sources. | [optional] [default to null]
**DeviceName** | **string** | Device name of the virtual network interface. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

