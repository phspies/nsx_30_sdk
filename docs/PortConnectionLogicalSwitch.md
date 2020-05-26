# PortConnectionLogicalSwitch

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | [***ManagedResource**](ManagedResource.md) |  | [optional] [default to null]
**Id** | **string** | Resource ID is mapped to this. (ID is Generated for Edge node groups, since resource will be null) | [optional] [default to null]
**VmPortsStates** | [**[]LogicalPortState**](LogicalPortState.md) | States of Logical Ports that are attached to a VIF/VM | [optional] [default to null]
**VmPorts** | [**[]LogicalPort**](LogicalPort.md) | Logical Ports that are attached to a VIF/VM | [optional] [default to null]
**VmVnics** | [**[]VirtualNetworkInterface**](VirtualNetworkInterface.md) | Virutal Network Interfaces that are attached to the Logical Ports | [optional] [default to null]
**RouterPorts** | [**[]LogicalPort**](LogicalPort.md) | Logical Ports that are attached to a router | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

