# PortMirroringSession

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
**Direction** | **string** | Port mirroring session direction | [default to null]
**MirrorSources** | [**[]MirrorSource**](MirrorSource.md) | Mirror sources | [default to null]
**EncapsulationVlanId** | **int64** | Only for Remote SPAN Port Mirror. | [optional] [default to null]
**SessionType** | **string** | If this property is unset, this session will be treated as LocalPortMirrorSession.  | [optional] [default to SESSION_TYPE.LOCAL_PORT_MIRROR_SESSION]
**SnapLength** | **int64** | If this property is set, the packet will be truncated to the provided length. If this property is unset, entire packet will be mirrored.  | [optional] [default to null]
**PortMirroringFilters** | [**[]PortMirroringFilter**](PortMirroringFilter.md) | An array of 5-tuples used to filter packets for the mirror session, if not provided, all the packets will be mirrored. | [optional] [default to null]
**PreserveOriginalVlan** | **bool** | Only for Remote SPAN Port Mirror. Whether to preserve original VLAN. | [optional] [default to false]
**MirrorDestination** | [***MirrorDestination**](MirrorDestination.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

