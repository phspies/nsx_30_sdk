# UpgradeUnitGroupAggregateInfo

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
**Status** | **string** | Upgrade status of upgrade unit group | [optional] [default to null]
**UpgradeUnitCount** | **int32** | Number of upgrade units in the group | [optional] [default to null]
**FailedCount** | **int32** | Number of nodes in the upgrade unit group that failed upgrade | [optional] [default to null]
**Type_** | **string** | Component type | [default to null]
**PercentComplete** | **float64** | Indicator of upgrade progress in percentage | [optional] [default to null]
**PostUpgradeStatus** | [***UpgradeChecksExecutionStatus**](UpgradeChecksExecutionStatus.md) |  | [optional] [default to null]
**Enabled** | **bool** | Flag to indicate whether upgrade of this group is enabled or not | [optional] [default to true]
**UpgradeUnits** | [**[]UpgradeUnit**](UpgradeUnit.md) | List of upgrade units in the group | [optional] [default to null]
**ExtendedConfiguration** | [**[]KeyValuePair**](KeyValuePair.md) | Extended configuration for the group | [optional] [default to null]
**Parallel** | **bool** | Upgrade method to specify whether the upgrade is to be performed in parallel or serially | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
