# IpfixSwitchUpmProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | All IPFIX profile types. | [default to null]
**Priority** | **int32** | This priority field is used to resolve conflicts in logical ports/switch  which inherit multiple switch IPFIX profiles from NSGroups.  Override rule is : for multiple profiles inherited from NSGroups, the one with highest priority (lowest number) overrides others; the profile directly applied to logical switch overrides profiles inherited from NSGroup; the profile directly applied to logical port overides profiles inherited from logical switch and/or nsgroup;  The IPFIX exporter will send records to collectors of final effective profile only.  | [default to null]
**CollectorProfile** | **string** | Each IPFIX switching profile can have its own collector profile.  | [default to null]
**IdleTimeout** | **int32** | The time in seconds after a flow is expired if no more packets matching this flow are received by the cache.  | [optional] [default to 300]
**MaxFlows** | **int64** | The maximum number of flow entries in each exporter flow cache.  | [optional] [default to 16384]
**ObservationDomainId** | **int64** | An identifier that is unique to the exporting process and used to meter the Flows.  | [default to null]
**ActiveTimeout** | **int32** | The time in seconds after a flow is expired even if more packets matching this Flow are received by the cache.  | [optional] [default to 300]
**ExportOverlayFlow** | **bool** | It controls whether sample result includes overlay flow info.  | [optional] [default to true]
**AppliedTos** | [***AppliedTos**](AppliedTos.md) |  | [optional] [default to null]
**PacketSampleProbability** | **float64** | The probability in percentage that a packet is sampled. The value should be  in range (0,100] and can only have three decimal places at most. The probability  is equal for every packet.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

