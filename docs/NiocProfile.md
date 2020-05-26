# NiocProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequiredCapabilities** | **[]string** |  | [optional] [default to null]
**ResourceType** | **string** | Supported HostSwitch profiles. | [default to null]
**HostInfraTrafficRes** | [**[]ResourceAllocation**](ResourceAllocation.md) | host_infra_traffic_res specifies bandwidth allocation for various traffic resources.  | [optional] [default to null]
**Enabled** | **bool** | The enabled property specifies the status of NIOC feature. When enabled is set to true, NIOC feature is turned on and the bandwidth allocations specified for the traffic resources are enforced. When enabled is set to false, NIOC feature is turned off and no bandwidth allocation is guaranteed. By default, enabled will be set to true.  | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

