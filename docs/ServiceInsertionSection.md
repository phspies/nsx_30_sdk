# ServiceInsertionSection

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stateful** | **bool** | Stateful or Stateless nature of distributed service section is enforced on all rules inside the section. Layer3 sections can be stateful or stateless. Layer2 sections can only be stateless. | [default to null]
**IsDefault** | **bool** | It is a boolean flag which reflects whether a distributed service section is default section or not. Each Layer 3 and Layer 2 section will have at least and at most one default section. | [optional] [default to null]
**AppliedTos** | [**[]ResourceReference**](ResourceReference.md) | List of objects where the rules in this section will be enforced. This will take precedence over rule level appliedTo. | [optional] [default to null]
**RuleCount** | **int64** | Number of rules in this section. | [optional] [default to null]
**SectionType** | **string** | Type of the rules which a section can contain. Only homogeneous sections are supported. | [default to null]
**TcpStrict** | **bool** | Ensures that a three way TCP handshake is done before the data packets are sent if the value is set to be true. tcp_strict&#x3D;true is supported only for stateful sections. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

