# IdentitySourceLdapServerEndpoint

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL for the LDAP server. Supported URL schemes are LDAP and LDAPS. Either a hostname or an IP address may be given, and the port number is optional and defaults to 389 for the LDAP scheme and 636 for the LDAPS scheme. | [default to null]
**UseStarttls** | **bool** | If set to true, Use the StartTLS extended operation to upgrade the connection to TLS before sending any sensitive information. The LDAP server must support the StartTLS extended operation in order for this protocol to operate correctly. This option is ignored if the URL scheme is LDAPS.  | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

