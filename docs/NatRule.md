# NatRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemOwned** | **bool** | Indicates system owned resource | [optional] [default to null]
**DisplayName** | **string** | Defaults to ID if not set | [optional] [default to null]
**Description** | **string** | Description of this resource | [optional] [default to null]
**Tags** | [**[]Tag**](Tag.md) | Opaque identifiers meaningful to the API user | [optional] [default to null]
**CreateUser** | **string** | ID of the user who created this resource | [optional] [default to null]
**Protection** | **string** | Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed             to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed                 to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super                    user and can modify it, but only when providing                    the request header X-Allow-Overwrite&#x3D;true. UNKNOWN - the _protection field could not be determined for this           entity.  | [optional] [default to null]
**CreateTime** | **int64** | Timestamp of resource creation | [optional] [default to null]
**LastModifiedTime** | **int64** | Timestamp of last modification | [optional] [default to null]
**LastModifiedUser** | **string** | ID of the user who last modified this resource | [optional] [default to null]
**Id** | **string** | Unique identifier of this resource | [optional] [default to null]
**ResourceType** | **string** | The type of this resource. | [optional] [default to null]
**MatchDestinationNetwork** | **string** | IP Address | CIDR | (null implies Any)  | [optional] [default to null]
**TranslatedNetwork** | **string** | The translated address for the matched IP packet. For a SNAT, it can be a single ip address, an ip range, or a CIDR block. For a DNAT and a REFLEXIVE, it can be a single ip address or a CIDR block. Translated network is not supported for NO_SNAT or NO_DNAT.  | [optional] [default to null]
**RulePriority** | **int64** | Ascending, valid range [0-2147483647]. If multiple rules have the same priority, evaluation sequence is undefined.  | [optional] [default to 1024]
**MatchService** | [***NsServiceElement**](NSServiceElement.md) |  | [optional] [default to null]
**AppliedTos** | [**[]ResourceReference**](ResourceReference.md) | Holds the list of LogicalRouterPort Ids that a NAT rule can be applied to. The LogicalRouterPort used must belong to the same LogicalRouter for which the NAT Rule is created. As of now a NAT rule can only have a single LogicalRouterPort as applied_tos. When applied_tos is not set, the NAT rule is applied to all LogicalRouterPorts beloging to the LogicalRouter. | [optional] [default to null]
**Enabled** | **bool** | Indicator to enable/disable the rule.  | [optional] [default to true]
**InternalRuleId** | **string** | Internal NAT rule uuid for debug used in Controller and backend. | [optional] [default to null]
**Logging** | **bool** | Enable/disable the logging of rule.  | [optional] [default to false]
**TranslatedPorts** | **string** | The translated port(s) for the mtached IP packet. It can be a single port or a port range. Please note, port translating is supported only for DNAT.  | [optional] [default to null]
**Action** | **string** | Valid actions: SNAT, DNAT, NO_SNAT, NO_DNAT, REFLEXIVE, NAT64. All rules in a logical router are either stateless or stateful. Mix is not supported. SNAT and DNAT are stateful, can NOT be supported when the logical router is running at active-active HA mode; REFLEXIVE is stateless. NO_SNAT and NO_DNAT have no translated_fields, only match fields are supported.  | [default to null]
**FirewallMatch** | **string** | Indicate how firewall is applied to a traffic packet. Firewall can be bypassed, or be applied to external/internal address of NAT rule.  The firewall_match will take priority over nat_pass. If the firewall_match is not provided, the nat_pass will be picked up.  | [optional] [default to null]
**NatPass** | **bool** | Default is true. If the nat_pass is set to true, the following firewall stage will be skipped. Please note, if action is NO_SNAT or NO_DNAT, then nat_pass must be set to true or omitted.  Nat_pass was deprecated with an alternative firewall_match. Please stop using nat_pass to specify whether firewall stage is skipped. if you want to skip, please set firewall_match to BYPASS. If you do not want to skip, please set the firewall_match to MATCH_EXTERNAL_ADDRESS or MATCH_INTERNAL_ADDRESS.  Please note, the firewall_match will take priority over the nat_pass. If both are provided, the nat_pass is ignored. If firewall_match is not provided while the nat_pass is specified, the nat_pass will still be picked up. In this case, if nat_pass is set to false, firewall rule will be applied on internall address of a packet, i.e. MATCH_INTERNAL_ADDRESS.  | [optional] [default to true]
**LogicalRouterId** | **string** | The logical router id which the nat rule runs on. | [optional] [default to null]
**MatchSourceNetwork** | **string** | IP Address | CIDR | (null implies Any)  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

