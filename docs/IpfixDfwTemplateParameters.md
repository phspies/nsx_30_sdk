# IpfixDfwTemplateParameters

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceIcmpType** | **bool** | Type of the IPv4 ICMP message.  | [optional] [default to true]
**IcmpCode** | **bool** | Code of the IPv4 ICMP message.  | [optional] [default to true]
**DestinationTransportPort** | **bool** | The destination transport port of a monitored network flow.  | [optional] [default to true]
**OctetDeltaCount** | **bool** | The number of octets since the previous report (if any) in incoming packets for this flow at the observation point. The number of octets include IP header(s) and payload.  | [optional] [default to true]
**VifUuid** | **bool** | VIF UUID - enterprise specific Information Element that uniquely identifies VIF.  | [optional] [default to true]
**ProtocolIdentifier** | **bool** | The value of the protocol number in the IP packet header.  | [optional] [default to true]
**FirewallEvent** | **bool** | Five valid values are allowed: 1. Flow Created. 2. Flow Deleted. 3. Flow Denied. 4. Flow Alert (not used in DropKick implementation). 5. Flow Update.  | [optional] [default to true]
**FlowDirection** | **bool** | Two valid values are allowed: 1. 0x00: igress flow to VM. 2. 0x01: egress flow from VM.  | [optional] [default to true]
**FlowEnd** | **bool** | The absolute timestamp (seconds) of the last packet of this flow.  | [optional] [default to true]
**SourceTransportPort** | **bool** | The source transport port of a monitored network flow.  | [optional] [default to true]
**PacketDeltaCount** | **bool** | The number of incoming packets since the previous report (if any) for this flow at the observation point.  | [optional] [default to true]
**DestinationAddress** | **bool** | The destination IP address of a monitored network flow.  | [optional] [default to true]
**SourceAddress** | **bool** | The source IP address of a monitored network flow.  | [optional] [default to true]
**RuleId** | **bool** | Firewall rule Id - enterprise specific Information Element that uniquely identifies firewall rule.  | [optional] [default to true]
**FlowStart** | **bool** | The absolute timestamp (seconds) of the first packet of this flow.  | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

