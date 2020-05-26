# ActiveDirectoryIdentitySource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LdapServers** | [**[]IdentitySourceLdapServer**](IdentitySourceLdapServer.md) | The list of LDAP servers that provide LDAP service for this identity source. Currently, only one LDAP server is supported. | [optional] [default to null]
**DomainName** | **string** | The name of the authentication domain. When users log into NSX using an identity of the form \&quot;user@domain\&quot;, NSX uses the domain portion to determine which LDAP identity source to use. For Active Directory, this domain name must match the domain of the Active Directory. | [default to null]
**ResourceType** | **string** |  | [default to null]
**BaseDn** | **string** | The subtree of the LDAP identity source to search when locating users and groups. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

