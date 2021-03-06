# HostSwitchState

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostSwitchType** | **string** | VDS represents VMware vSphere Distributed Switch from vSphere that is used as HostSwitch through TransportNode or TransportNodeProfile configuration. When VDS is used as a HostSwitch, Hosts have to be added to VDS from vSphere and VDS instance is created on Hosts. To configure NSX on such hosts, you can use this VDS as a HostSwitch from NSX manager. vCenter has the ownership of MTU, LAG, NIOC and LLDP configuration of such VDS backed HostSwitch. Remaining configuration (e.g. UplinkHostswitchProfile) will be managed by NSX. NVDS represents NSX Virtual Switch which is NSX native HostSwitch. All configurations of NVDS will be managed by NSX. | [optional] [default to HOST_SWITCH_TYPE.NVDS]
**HostSwitchId** | **string** | External ID of the HostSwitch | [optional] [default to null]
**Endpoints** | [**[]Endpoint**](Endpoint.md) | List of virtual tunnel endpoints which are configured on this switch | [optional] [default to null]
**TransportZoneIds** | **[]string** | List of Ids of TransportZones this HostSwitch belongs to | [optional] [default to null]
**HostSwitchName** | **string** | The name must be unique among all host switches specified in a given Transport Node. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

