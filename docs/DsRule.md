# DsRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | [***OwnerResourceLink**](OwnerResourceLink.md) |  | [optional] [default to null]
**DisplayName** | **string** | Defaults to ID if not set | [optional] [default to null]
**Id** | **string** | Identifier of the resource | [optional] [default to null]
**ResourceType** | **string** | The type of this resource. | [optional] [default to null]
**Description** | **string** | Description of this resource | [optional] [default to null]
**IsDefault** | **bool** | Flag to indicate whether rule is default. | [optional] [default to null]
**Direction** | **string** | Rule direction in case of stateless distributed service rules. This will only considered if section level parameter is set to stateless. Default to IN_OUT if not specified. | [optional] [default to DIRECTION.IN_OUT]
**RuleTag** | **string** | User level field which will be printed in CLI and packet logs. | [optional] [default to null]
**IpProtocol** | **string** | Type of IP packet that should be matched while enforcing the rule. | [optional] [default to IP_PROTOCOL.IPV4_IPV6]
**Notes** | **string** | User notes specific to the rule. | [optional] [default to null]
**AppliedTos** | [**[]ResourceReference**](ResourceReference.md) | List of object where rule will be enforced. The section level field overrides this one. Null will be treated as any. | [optional] [default to null]
**Logged** | **bool** | Flag to enable packet logging. Default is disabled. | [optional] [default to false]
**Disabled** | **bool** | Flag to disable rule. Disabled will only be persisted but never provisioned/realized. | [optional] [default to false]
**Sources** | [**[]ResourceReference**](ResourceReference.md) | List of sources. Null will be treated as any. | [optional] [default to null]
**Action** | **string** | Action enforced on the packets which matches the distributed service rule. Currently DS Layer supports below actions. ALLOW           - Forward any packet when a rule with this action gets a match (Used by Firewall). DROP            - Drop any packet when a rule with this action gets a match. Packets won&#x27;t go further(Used by Firewall). REJECT          - Terminate TCP connection by sending TCP reset for a packet when a rule with this action gets a match (Used by Firewall). REDIRECT        - Redirect any packet to a partner appliance when a rule with this action gets a match (Used by Service Insertion). DO_NOT_REDIRECT - Do not redirect any packet to a partner appliance when a rule with this action gets a match (Used by Service Insertion). DETECT          - Detect IDS Signatures. | [default to null]
**Priority** | **int64** | Priority of the rule. | [optional] [default to null]
**SourcesExcluded** | **bool** | Negation of the source. | [optional] [default to false]
**DestinationsExcluded** | **bool** | Negation of the destination. | [optional] [default to false]
**Destinations** | [**[]ResourceReference**](ResourceReference.md) | List of the destinations. Null will be treated as any. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

