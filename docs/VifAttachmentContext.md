# VifAttachmentContext

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocateAddresses** | **string** | A flag to indicate whether to allocate addresses from allocation     pools bound to the parent logical switch.  | [optional] [default to null]
**ResourceType** | **string** | Used to identify which concrete class it is | [default to null]
**VifType** | **string** | Type of the VIF attached to logical port | [default to null]
**ParentVifId** | **string** | VIF ID of the parent VIF if vif_type is CHILD | [optional] [default to null]
**AppId** | **string** | An application ID used to identify / look up a child VIF behind a parent VIF. Only effective when vif_type is CHILD.  | [optional] [default to null]
**TrafficTag** | **int32** | Current we use VLAN id as the traffic tag. Only effective when vif_type is CHILD. Each logical port inside a container must have a unique traffic tag. If the traffic_tag is not unique, no error is generated, but traffic will not be delivered to any port with a non-unique tag.  | [optional] [default to null]
**TransportNodeUuid** | **string** | Only effective when vif_type is INDEPENDENT. Each logical port inside a bare metal server or container must have a transport node UUID. We use transport node ID as transport node UUID.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

