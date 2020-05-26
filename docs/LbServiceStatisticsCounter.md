# LbServiceStatisticsCounter

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**L4TotalSessions** | **int64** | Number of l4 total sessions | [optional] [default to null]
**L7CurrentSessions** | **int64** | Number of l7 current sessions | [optional] [default to null]
**L4MaxSessions** | **int64** | L4 max sessions is used to show the peak L4 max session data since load balancer starts to provide service.  | [optional] [default to null]
**L4CurrentSessions** | **int64** | Number of l4 current sessions | [optional] [default to null]
**L7TotalSessions** | **int64** | Number of l7 total sessions | [optional] [default to null]
**L4CurrentSessionRate** | **float64** | The average number of l4 current sessions per second, the number is averaged over the last 5 one-second intervals.  | [optional] [default to null]
**L7CurrentSessionRate** | **float64** | The average number of l7 current requests per second, the number is averaged over the last 5 one-second intervals.  | [optional] [default to null]
**L7MaxSessions** | **int64** | L7 max sessions is used to show the peak L7 max session data since load balancer starts to provide service.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

