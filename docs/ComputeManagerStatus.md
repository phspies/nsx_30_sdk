# ComputeManagerStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | Version of the compute manager | [optional] [default to null]
**ConnectionStatus** | **string** | Status of connection with the compute manager | [optional] [default to null]
**ConnectionErrors** | [**[]ErrorInfo**](ErrorInfo.md) | Errors when connecting with compute manager | [optional] [default to null]
**OidcEndPointId** | **string** | If Compute manager is trusted as authorization server, then this Id will be Id of corresponding oidc end point.  | [optional] [default to null]
**LastSyncTime** | **int64** | Timestamp of the last successful update of Inventory, in epoch milliseconds. | [optional] [default to null]
**ConnectionStatusDetails** | **string** | Details about connection status | [optional] [default to null]
**RegistrationErrors** | [**[]ErrorInfo**](ErrorInfo.md) | Errors when registering with compute manager | [optional] [default to null]
**RegistrationStatus** | **string** | Registration status of compute manager | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

