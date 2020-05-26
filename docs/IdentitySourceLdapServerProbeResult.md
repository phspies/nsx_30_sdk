# IdentitySourceLdapServerProbeResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | THe URL of the probed LDAP host. | [optional] [default to null]
**Errors** | [**[]LdapProbeError**](LdapProbeError.md) | Detail about errors encountered during the probe. | [optional] [default to null]
**Result** | **string** | Overall result of the probe. If the probe was able to connect to the LDAP service, authenticate using the provided credentials, and perform searches of the configured user and group search bases without error, the result is SUCCESS.  Otherwise, the result is FAILURE, and additional details may be found in the errors property. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

