# LogicalPortState

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransportNodeIds** | **[]string** | Identifiers of the transport node where the port is located | [optional] [default to null]
**Attachment** | [***LogicalPortAttachmentState**](LogicalPortAttachmentState.md) |  | [optional] [default to null]
**DuplicateBindings** | [**[]DuplicateAddressBindingEntry**](DuplicateAddressBindingEntry.md) | If any address binding discovered on the port is also found on other port on the same logical switch, then it is included in the duplicate bindings list along with the ID of the port with which it conflicts.  | [optional] [default to null]
**DiscoveredBindings** | [**[]AddressBindingEntry**](AddressBindingEntry.md) | Contains the list of address bindings for a logical port that were automatically dicovered using various snooping methods like ARP, DHCP etc.  | [optional] [default to null]
**Id** | **string** | Id of the logical port | [default to null]
**RealizedBindings** | [**[]AddressBindingEntry**](AddressBindingEntry.md) | List of logical port bindings that are realized. This list may be populated from the discovered bindings or manual user specified bindings. This binding configuration can be used by features such as firewall, spoof-guard, traceflow etc.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

