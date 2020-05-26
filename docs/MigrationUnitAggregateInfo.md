# MigrationUnitAggregateInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [***SelfResourceLink**](SelfResourceLink.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](ResourceLink.md) | The server will populate this field when returing the resource. Ignored on PUT and POST. | [optional] [default to null]
**Schema** | **string** | Schema for this resource | [optional] [default to null]
**Status** | **string** | Status of migration unit | [optional] [default to null]
**PercentComplete** | **float64** | Indicator of migration progress in percentage | [optional] [default to null]
**Errors** | **[]string** | List of errors occurred during migration of this migration unit | [optional] [default to null]
**Unit** | [***MigrationUnit**](MigrationUnit.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

