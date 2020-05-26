# RoutingGlobalConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | Valid Global configuration types | [default to null]
**LogicalUplinkMtu** | **int32** | This is the global default MTU for all the logical uplinks in a NSX domain. Currently logical uplink MTU can only be set globally and applies to the entire NSX domain. There is no option to override this value at transport zone level or transport node level. If this value is not set, the default value of 1500 will be used. | [optional] [default to 1500]
**VdrMacNested** | **string** | This is the global default MAC address for all VDRs in all transport nodes in a NSX system nested in another NSX system. It can be changed only when there is no transport node in the NSX system. All transport zones in such a nested NSX system will have the \&quot;nested_nsx\&quot; property being true so that all transport nodes will use this MAC for the VDR ports to avoid conflict with the VDR MAC in the outer NSX system. | [optional] [default to 02:50:56:56:44:53]
**L3ForwardingMode** | **string** | This setting does not restrict configuration as per other modes. But the forwarding will only work as per the mode set here. | [default to L3_FORWARDING_MODE.ONLY]
**VdrMac** | **string** | This is the global default MAC address for all VDRs in all transport nodes in a NSX system. It can be changed only when there is no transport node in the NSX system. | [optional] [default to 02:50:56:56:44:52]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

