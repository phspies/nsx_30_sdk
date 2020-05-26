# NsGroupTagExpression

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** |  | [default to null]
**TagOp** | **string** | Target_type VirtualMachine supports all specified operators for tag expression while LogicalSwitch and LogicalPort supports only EQUALS operator. All operators perform a case insensitive match.  | [optional] [default to TAG_OP.EQUALS]
**Scope** | **string** | The tag.scope attribute of the object | [optional] [default to null]
**ScopeOp** | **string** | Operator of the scope expression eg- tag.scope &#x3D; \&quot;S1\&quot;. | [optional] [default to SCOPE_OP.EQUALS]
**Tag** | **string** | The tag.tag attribute of the object | [optional] [default to null]
**TargetType** | **string** | Type of the resource on which this expression is evaluated | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

