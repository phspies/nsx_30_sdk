# TraceflowObservationDroppedLogical

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NatRuleId** | **int64** | The ID of the NAT rule that was applied to forward the traceflow packet | [optional] [default to null]
**Reason** | **string** | The reason traceflow packet was dropped | [optional] [default to null]
**LportId** | **string** | The id of the logical port at which the traceflow packet was dropped | [optional] [default to null]
**LportName** | **string** | The name of the logical port at which the traceflow packet was dropped | [optional] [default to null]
**AclRuleId** | **int64** | The id of the acl rule that was applied to drop the traceflow packet | [optional] [default to null]
**ArpFailReason** | **string** | This field specifies the ARP fails reason ARP_TIMEOUT - ARP failure due to query control plane timeout ARP_CPFAIL - ARP failure due post ARP query message to control plane failure ARP_FROMCP - ARP failure due to deleting ARP entry from control plane ARP_PORTDESTROY - ARP failure due to port destruction ARP_TABLEDESTROY - ARP failure due to ARP table destruction ARP_NETDESTROY - ARP failure due to overlay network destruction | [optional] [default to null]
**ServicePathIndex** | **int64** | The index of service path that is a chain of services represents the point where the traceflow packet was dropped.  | [optional] [default to null]
**ComponentId** | **string** | The id of the component that dropped the traceflow packet. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

