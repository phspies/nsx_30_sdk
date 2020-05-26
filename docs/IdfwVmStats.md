# IdfwVmStats

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmExtId** | **string** | Virtual machine (external ID or BIOS UUID) where login/logout event occurred. | [default to null]
**ActiveSessions** | [**[]IdfwUserSessionData**](IdfwUserSessionData.md) | List of active (still logged in) user login/sessions data (no limit) | [default to null]
**ArchivedSessions** | [**[]IdfwUserSessionData**](IdfwUserSessionData.md) | Optional list of up to 5 most recent archived (previously logged in) user login/session data. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

