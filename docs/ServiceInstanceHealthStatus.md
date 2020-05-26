# ServiceInstanceHealthStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSvaMuxIncompatible** | **bool** | Protocol version might be different in both Mux and SVA. | [optional] [default to null]
**ConnectTimestamp** | **string** | Latest timestamp when mux was connected to SVA. | [optional] [default to null]
**MuxIncompatibleVersion** | **string** | Mux version when Mux and SVA are incompatible | [optional] [default to null]
**SolutionVersion** | **string** | Version of third party partner solution application. | [optional] [default to null]
**SyncTime** | **string** | Latest timestamp when health status is received. | [optional] [default to null]
**SolutionStatus** | **string** | Status of third party partner solution application. | [optional] [default to null]
**IsStale** | **bool** | The parameter is set if the last received health status is older than the predefined interval.  | [optional] [default to null]
**MuxConnectedStatus** | **string** | Status of multiplexer which forwards the events from guest virtual machines to the partner appliance. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

