# IdentitySourceLdapServer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificates** | **[]string** | If using LDAPS or STARTTLS, provide the X.509 certificate of the LDAP server in PEM format. This property is not required when connecting without TLS encryption and is ignored in that case. | [optional] [default to null]
**BindIdentity** | **string** | A username used to authenticate to the directory when admnistering roles in NSX. This user should have privileges to search the LDAP directory for groups and users. This user is also used in some cases (OpenLDAP) to look up an NSX user&#x27;s distinguished name based on their NSX login name. If omitted, NSX will authenticate to the LDAP server using an LDAP anonymous bind operation. For Active Directory, provide a userPrincipalName (e.g. administrator@airius.com) or the full distinguished nane. For OpenLDAP, provide the distinguished name of the user (e.g. uid&#x3D;admin, cn&#x3D;airius, dc&#x3D;com). | [optional] [default to null]
**UseStarttls** | **bool** | If set to true, Use the StartTLS extended operation to upgrade the connection to TLS before sending any sensitive information. The LDAP server must support the StartTLS extended operation in order for this protocol to operate correctly. This option is ignored if the URL scheme is LDAPS.  | [optional] [default to false]
**Url** | **string** | The URL for the LDAP server. Supported URL schemes are LDAP and LDAPS. Either a hostname or an IP address may be given, and the port number is optional and defaults to 389 for the LDAP scheme and 636 for the LDAPS scheme. | [default to null]
**Password** | **string** | A password used when authenticating to the directory. | [optional] [default to null]
**Enabled** | **bool** | Allows the LDAP server to be enabled or disabled. When disabled, this LDAP server will not be used to authenticate users. | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

