# TraceflowObservation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimestampMicro** | **int64** | Timestamp when the observation was created by the transport node (microseconds epoch) | [optional] [default to null]
**ComponentSubType** | **string** | The sub type of the component that issued the observation. | [optional] [default to null]
**ResourceType** | **string** |  | [default to null]
**ComponentType** | **string** | The type of the component that issued the observation. | [optional] [default to null]
**TransportNodeName** | **string** | name of the transport node that observed a traceflow packet | [optional] [default to null]
**Timestamp** | **int64** | Timestamp when the observation was created by the transport node (milliseconds epoch) | [optional] [default to null]
**TransportNodeId** | **string** | id of the transport node that observed a traceflow packet | [optional] [default to null]
**SequenceNo** | **int64** | the hop count for observations on the transport node that a traceflow packet is injected in will be 0. The hop count is incremented each time a subsequent transport node receives the traceflow packet. The sequence number of 999 indicates that the hop count could not be determined for the containing observation. | [optional] [default to null]
**TransportNodeType** | **string** | type of the transport node that observed a traceflow packet | [optional] [default to null]
**ComponentName** | **string** | The name of the component that issued the observation. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

