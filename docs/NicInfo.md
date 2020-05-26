# NicInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubnetMask** | **string** | Subnet mask associated with the NIC metadata. | [optional] [default to null]
**GatewayAddress** | **string** | Gateway address associated with the NIC metadata. | [optional] [default to null]
**IpAllocationType** | **string** | IP allocation type with values STATIC, DHCP, or NONE indicating that IP address is not required. | [optional] [default to null]
**NicMetadata** | [***NicMetadata**](NicMetadata.md) |  | [optional] [default to null]
**NetworkId** | **string** | Network Id associated with the NIC metadata. It can be a moref, or a logical switch ID. If it is to be taken from &#x27;Agent VM Settings&#x27;, then it should be empty. | [optional] [default to null]
**IpPoolId** | **string** | If the nic should get IP using a static IP pool then IP pool id should be provided here. | [optional] [default to null]
**IpAddress** | **string** | IP address associated with the NIC metadata. Required only when assigning IP statically for a deployment that is for a single VM instance. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

