# DistributedVirtualSwitch

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CmLocalId** | **string** | ID of the virtual switch in compute manager | [optional] [default to null]
**ExternalId** | **string** | External id of the virtual switch | [optional] [default to null]
**OriginType** | **string** | Switch type like VmwareDistributedVirtualSwitch | [optional] [default to null]
**OriginId** | **string** | ID of the compute manager where this virtual switch is discovered.  | [optional] [default to null]
**DiscoveredNodes** | [**[]DiscoveredNode**](DiscoveredNode.md) | Array of discovered nodes connected to this switch. | [optional] [default to null]
**UplinkPortgroup** | [***DistributedVirtualPortgroup**](DistributedVirtualPortgroup.md) |  | [optional] [default to null]
**Uuid** | **string** | UUID of the switch | [optional] [default to null]
**OriginProperties** | [**[]KeyValuePair**](KeyValuePair.md) | Key-Value map of additional properties of switch | [optional] [default to null]
**LacpGroupConfigs** | [**[]LacpGroupConfigInfo**](LacpGroupConfigInfo.md) | It contains information about VMware specific multiple dynamic LACP groups.  | [optional] [default to null]
**UplinkPortNames** | **[]string** | The uniform name of uplink ports on each host. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

