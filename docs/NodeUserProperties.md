# NodeUserProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [***SelfResourceLink**](SelfResourceLink.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](ResourceLink.md) | The server will populate this field when returing the resource. Ignored on PUT and POST. | [optional] [default to null]
**Schema** | **string** | Schema for this resource | [optional] [default to null]
**Username** | **string** | User login name (must be \&quot;root\&quot; if userid is 0) | [optional] [default to null]
**Status** | **string** | Status of the user. This value can be ACTIVE indicating authentication attempts will be successful if the correct credentials are specified. The value can also be PASSWORD_EXPIRED indicating authentication attempts will fail because the user&#x27;s password has expired and must be changed. Or, this value can be NOT_ACTIVATED indicating the user&#x27;s password has not yet been set and must be set before the user can authenticate. | [optional] [default to null]
**LastPasswordChange** | **int64** | Number of days since password was last changed | [optional] [default to null]
**FullName** | **string** | Full name for the user | [optional] [default to null]
**PasswordChangeFrequency** | **int64** | Number of days password is valid before it must be changed. This can be set to 0 to indicate no password change is required or a positive integer up to 9999. By default local user passwords must be changed every 90 days. | [optional] [default to null]
**Password** | **string** | Password for the user (optionally specified on PUT, unspecified on GET) | [optional] [default to null]
**Userid** | **int64** | Numeric id for the user | [optional] [default to null]
**OldPassword** | **string** | Old password for the user (required on PUT if password specified) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

