# AuthenticationPolicyProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [***SelfResourceLink**](SelfResourceLink.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](ResourceLink.md) | The server will populate this field when returing the resource. Ignored on PUT and POST. | [optional] [default to null]
**Schema** | **string** | Schema for this resource | [optional] [default to null]
**ApiFailedAuthResetPeriod** | **int64** | In order to trigger an account lockout, all authentication failures must occur in this time window. If the reset period expires, the failed login count is reset to zero. Only applies to NSX Manager nodes. Ignored on other node types. | [optional] [default to 900]
**MinimumPasswordLength** | **int64** | Minimum number of characters required in account passwords | [optional] [default to 8]
**CliFailedAuthLockoutPeriod** | **int64** | Once a lockout occurs, the account remains locked out of the CLI for this time period. While the lockout period is in effect, additional authentication attempts restart the lockout period, even if a valid password is specified. | [optional] [default to 900]
**ApiMaxAuthFailures** | **int64** | Only applies to NSX Manager nodes. Ignored on other node types. | [optional] [default to 5]
**ApiFailedAuthLockoutPeriod** | **int64** | Once a lockout occurs, the account remains locked out of the API for this time period. Only applies to NSX Manager nodes. Ignored on other node types. | [optional] [default to 900]
**CliMaxAuthFailures** | **int64** | Number of authentication failures that trigger CLI lockout | [optional] [default to 5]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

