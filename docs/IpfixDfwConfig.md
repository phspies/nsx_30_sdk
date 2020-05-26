# IpfixDfwConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedTos** | [**[]ResourceReference**](ResourceReference.md) | List of objects where the IPFIX Config will be enabled. | [optional] [default to null]
**ResourceType** | **string** | Supported IPFIX Config Types. | [default to null]
**Priority** | **int64** | This priority field is used to resolve conflicts in Logical Ports which are covered by more than one IPFIX profiles. The IPFIX exporter will send records to Collectors in highest priority profile (lowest number) only.  | [default to 0]
**Collector** | **string** | Each IPFIX DFW config can have its own collector config.  | [default to null]
**ActiveFlowExportTimeout** | **int64** | For long standing active flows, IPFIX records will be sent per timeout period  | [optional] [default to 1]
**TemplateParameters** | [***IpfixDfwTemplateParameters**](IpfixDfwTemplateParameters.md) |  | [optional] [default to null]
**ObservationDomainId** | **int64** | An identifier that is unique to the exporting process and used to meter the Flows.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

