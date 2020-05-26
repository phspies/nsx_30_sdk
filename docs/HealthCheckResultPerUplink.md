# HealthCheckResultPerUplink

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UplinkName** | **string** | Name of the uplink. | [optional] [default to null]
**VlanAndMtuAllowed** | [**[]HealthCheckVlanRange**](HealthCheckVlanRange.md) | List of VLAN ID ranges which are allowed by VLAN and MTU settings.  | [optional] [default to null]
**VlanDisallowed** | [**[]HealthCheckVlanRange**](HealthCheckVlanRange.md) | List of VLAN ID ranges which may be disallowed by VLAN settings.  | [optional] [default to null]
**MtuDisallowed** | [**[]HealthCheckVlanRange**](HealthCheckVlanRange.md) | List of VLAN ID ranges which are allowed by VLAN settings but may be disallowed by MTU settings.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

