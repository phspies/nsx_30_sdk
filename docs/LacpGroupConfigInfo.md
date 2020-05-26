# LacpGroupConfigInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The key represents the identifier for the group that is unique across VC.  | [optional] [default to null]
**Name** | **string** | The display name of the LACP group. | [optional] [default to null]
**UplinkPortKeys** | **[]string** | Keys for the uplink ports in the group. Each uplink port is assigned a key that is unique across VC.  | [optional] [default to null]
**LoadBalanceAlgorithm** | **string** | Load balance algorithm used in LACP group. The possible values are dictated by the values available in VC. Please refer VMwareDvsLacpLoadBalanceAlgorithm documentation for a full list of values. A few examples are srcDestIp where source and destination IP are considered, srcIp where only source IP is considered.  | [optional] [default to null]
**UplinkNum** | **int64** | The number of uplink ports | [optional] [default to null]
**UplinkNames** | **[]string** | Names for the uplink ports in the group. | [optional] [default to null]
**Mode** | **string** | The mode of LACP can be ACTIVE or PASSIVE. If the mode is ACTIVE, LACP is enabled unconditionally. If the mode is PASSIVE, LACP is enabled only if LACP device is detected.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

