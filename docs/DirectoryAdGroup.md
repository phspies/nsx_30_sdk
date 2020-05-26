# DirectoryAdGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainSyncNodeId** | **string** | Domain sync node under which this directory group is located. We currently sync only from Root node and hence this attribute doesn&#x27;t have a specific value set. | [optional] [default to null]
**DistinguishedName** | **string** | Directory group distinguished name | [default to null]
**DomainId** | **string** | Domain ID this directory group belongs to. | [default to null]
**ResourceType** | **string** | Directory group resource type comes from multiple sub-classes extending this base class. For example, DirectoryAdGroup is one accepted resource_type. If there are more sub-classes defined, they will also be accepted resource_type. | [default to null]
**DomainName** | **string** | Each active directory domain has a domain naming context (NC), which contains domain-specific data. The root of this naming context is represented by a domain&#x27;s distinguished name (DN) and is typically referred to as the NC head. | [default to null]
**ObjectGuid** | **string** | GUID is a 128-bit value that is unique not only in the enterprise but also across the world. GUIDs are assigned to every object created by Active Directory, not just User and Group objects. | [default to null]
**SecureId** | **string** | A security identifier (SID) is a unique value of variable length used to identify a trustee. A SID consists of the following components - The revision level of the SID structure; A 48-bit identifier authority value that identifies the authority that issued the SID; A variable number of subauthority or relative identifier (RID) values that uniquely identify the trustee relative to the authority that issued the SID. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

