# PrincipalIdentity

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
**IsProtected** | **bool** | Indicator whether the entities created by this principal should be protected. | [optional] [default to null]
**Role** | **string** | Role | [optional] [default to null]
**Name** | **string** | Name of the principal. | [default to null]
**PermissionGroup** | **string** | Use the &#x27;role&#x27; field instead and pass in &#x27;auditor&#x27; for read_only_api_users or &#x27;enterprise_admin&#x27; for the others. | [optional] [default to null]
**CertificateId** | **string** | Id of the stored certificate. When used with the deprecated POST /trust-management/principal-identities API this field is required. | [optional] [default to null]
**NodeId** | **string** | Unique node-id of a principal. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
