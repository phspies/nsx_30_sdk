# LdapIdentitySourceSearchResultItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dn** | **string** | Distinguished name (DN) of the entry. | [optional] [default to null]
**CommonName** | **string** | The Common Name (CN) of the entry, if available. | [optional] [default to null]
**PrincipalName** | **string** | For Active Directory (AD) users, this will be the user principal name (UPN), in the format user@domain. For non-AD users, this will be the user&#x27;s uid property, followed by \&quot;@\&quot; and the domain of the directory. For groups, this will be the group&#x27;s common name, followed by \&quot;@\&quot; and the domain of the directory. | [optional] [default to null]
**Type_** | **string** | Describes the type of the entry | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

