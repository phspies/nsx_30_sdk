# HealthCheckResultOnHostSwitch

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VlanMtuStatus** | **string** | Status of VLAN-MTU health check result on host switch.  | [optional] [default to null]
**HostSwitchName** | **string** | Name of the host switch. | [optional] [default to null]
**ResultsPerUplink** | [**[]HealthCheckResultPerUplink**](HealthCheckResultPerUplink.md) | List of health check results per uplink on current host switch of specific transport node.  | [optional] [default to null]
**UpdatedTime** | **int64** | Timestamp of check result updated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

