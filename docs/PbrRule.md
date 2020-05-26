# PbrRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | [***OwnerResourceLink**](OwnerResourceLink.md) |  | [optional] [default to null]
**DisplayName** | **string** | Defaults to ID if not set | [optional] [default to null]
**Id** | **string** | Identifier of the resource | [optional] [default to null]
**ResourceType** | **string** | The type of this resource. | [optional] [default to null]
**Description** | **string** | Description of this resource | [optional] [default to null]
**Disabled** | **bool** | Flag to disable rule. Disabled will only be persisted but never provisioned/realized. | [optional] [default to false]
**Sources** | [**[]ResourceReference**](ResourceReference.md) | List of sources. Null will be treated as any. | [optional] [default to null]
**RuleTag** | **string** | User level field which will be printed in CLI and packet logs. | [optional] [default to null]
**Services** | [**[]PbrService**](PBRService.md) | List of the services. Null will be treated as any. | [optional] [default to null]
**Notes** | **string** | User notes specific to the rule. | [optional] [default to null]
**AppliedTos** | [**[]ResourceReference**](ResourceReference.md) | List of object where rule will be enforced. field overrides this one. Null will be treated as any. | [optional] [default to null]
**Logged** | **bool** | Flag to enable packet logging. Default is disabled. | [optional] [default to false]
**Action** | **string** | Action enforced on the packets which matches the PBR rule. | [default to null]
**Destinations** | [**[]ResourceReference**](ResourceReference.md) | List of the destinations. Null will be treated as any. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

