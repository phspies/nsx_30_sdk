# ResourceAllocation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reservation** | **float64** | Minimum guaranteed bandwidth percentage | [default to null]
**TrafficType** | [***HostInfraTrafficType**](HostInfraTrafficType.md) |  | [default to null]
**Limit** | **float64** | The limit property specifies the maximum bandwidth allocation for a given traffic type and is expressed in percentage. The default value for this field is set to -1 which means the traffic is unbounded for the traffic type. All other negative values for this property is not supported and will be rejected by the API.  | [default to null]
**Shares** | **int32** | Shares | [default to 50]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

