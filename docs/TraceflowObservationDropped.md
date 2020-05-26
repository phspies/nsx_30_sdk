# TraceflowObservationDropped

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
**NatRuleId** | **int64** | The ID of the NAT rule that was applied to forward the traceflow packet | [optional] [default to null]
**Reason** | **string** | The reason traceflow packet was dropped | [optional] [default to null]
**LportId** | **string** | The id of the logical port at which the traceflow packet was dropped | [optional] [default to null]
**LportName** | **string** | The name of the logical port at which the traceflow packet was dropped | [optional] [default to null]
**AclRuleId** | **int64** | The id of the acl rule that was applied to drop the traceflow packet | [optional] [default to null]
**ArpFailReason** | **string** | This field specifies the ARP fails reason ARP_TIMEOUT - ARP failure due to query control plane timeout ARP_CPFAIL - ARP failure due post ARP query message to control plane failure ARP_FROMCP - ARP failure due to deleting ARP entry from control plane ARP_PORTDESTROY - ARP failure due to port destruction ARP_TABLEDESTROY - ARP failure due to ARP table destruction ARP_NETDESTROY - ARP failure due to overlay network destruction | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

