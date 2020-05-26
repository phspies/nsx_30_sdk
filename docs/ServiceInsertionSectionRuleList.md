# ServiceInsertionSectionRuleList

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TcpStrict** | **bool** | Ensures that a three way TCP handshake is done before the data packets are sent if the value is set to be true. tcp_strict&#x3D;true is supported only for stateful sections. | [optional] [default to false]
**Rules** | [**[]ServiceInsertionRule**](ServiceInsertionRule.md) | List of Service Insertion rules in the section. Only homogeneous rules are supported. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

