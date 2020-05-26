# MigrationUnit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [***SelfResourceLink**](SelfResourceLink.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](ResourceLink.md) | The server will populate this field when returing the resource. Ignored on PUT and POST. | [optional] [default to null]
**Schema** | **string** | Schema for this resource | [optional] [default to null]
**Group** | [***ResourceReference**](ResourceReference.md) |  | [optional] [default to null]
**Warnings** | **[]string** | List of warnings indicating issues with the migration unit that may result in migration failure | [optional] [default to null]
**CurrentVersion** | **string** | This is component version e.g. if migration unit is of type HOST, then this is host version. | [optional] [default to null]
**Metadata** | [**[]KeyValuePair**](KeyValuePair.md) | Metadata about migration unit | [optional] [default to null]
**Type_** | **string** | Migration unit type | [optional] [default to null]
**Id** | **string** | Identifier of the migration unit | [optional] [default to null]
**DisplayName** | **string** | Name of the migration unit | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

