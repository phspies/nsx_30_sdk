# IdfwUserSessionData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | AD user ID (may not exist). | [optional] [default to null]
**UserSessionId** | **int32** | User session ID.  This also indicates whether this is VDI / RDSH. | [default to null]
**VmExtId** | **string** | Virtual machine (external ID or BIOS UUID) where login/logout events occurred. | [optional] [default to null]
**Id** | **string** | Identifier of user session data. | [optional] [default to null]
**LoginTime** | **int64** | Login time. | [default to null]
**UserName** | **string** | AD user name. | [default to null]
**LogoutTime** | **int64** | Logout time if applicable.  An active user session has no logout time. Non-active user session is stored (up to last 5 most recent entries) per VM and per user.  | [optional] [default to null]
**DomainName** | **string** | AD Domain of user. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

