# DirectoryLdapServer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemOwned** | **bool** | Indicates system owned resource | [optional] [default to null]
**DisplayName** | **string** | Defaults to ID if not set | [optional] [default to null]
**Description** | **string** | Description of this resource | [optional] [default to null]
**Tags** | [**[]Tag**](Tag.md) | Opaque identifiers meaningful to the API user | [optional] [default to null]
**CreateUser** | **string** | ID of the user who created this resource | [optional] [default to null]
**Protection** | **string** | Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed             to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed                 to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super                    user and can modify it, but only when providing                    the request header X-Allow-Overwrite&#x3D;true. UNKNOWN - the _protection field could not be determined for this           entity.  | [optional] [default to null]
**CreateTime** | **int64** | Timestamp of resource creation | [optional] [default to null]
**LastModifiedTime** | **int64** | Timestamp of last modification | [optional] [default to null]
**LastModifiedUser** | **string** | ID of the user who last modified this resource | [optional] [default to null]
**Id** | **string** | Unique identifier of this resource | [optional] [default to null]
**ResourceType** | **string** | The type of this resource. | [optional] [default to null]
**Username** | **string** | Directory LDAP server connection user name. | [optional] [default to null]
**Host** | **string** | Directory LDAP server DNS host name or ip address which is reachable by NSX manager to be connected and do object synchronization. | [default to null]
**Protocol** | **string** | Directory LDAP server connection protocol which is either LDAP or LDAPS. | [optional] [default to PROTOCOL.LDAP]
**Thumbprint** | **string** | Directory LDAP server certificate thumbprint used in secure LDAPS connection. | [optional] [default to null]
**Password** | **string** | Directory LDAP server connection password. | [optional] [default to null]
**DomainName** | **string** | Directory domain name which best describes the domain. It could be unique fqdn name or it could also be descriptive. There is no unique contraint for domain name among different domains. | [optional] [default to null]
**Port** | **int64** | Directory LDAP server connection TCP/UDP port. | [optional] [default to 389]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

