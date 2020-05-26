# AlgTypeNsService

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | The specific type of NSServiceElement | [default to null]
**Alg** | **string** | The Application Layer Gateway (ALG) protocol. Please note, protocol NBNS_BROADCAST and NBDG_BROADCAST are  deprecated. Please use UDP protocol and create L4 Port Set type of service instead.  | [default to null]
**DestinationPorts** | **[]string** | The destination_port cannot be empty and must be a single value. | [default to null]
**SourcePorts** | **[]string** | Source ports | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

