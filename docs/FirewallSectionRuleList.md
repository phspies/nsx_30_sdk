# FirewallSectionRuleList

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** | Category from policy framework. | [optional] [default to null]
**EnforcedOn** | **string** | This attribute represents enforcement point of firewall section. For example, firewall section enforced on logical port with attachment type bridge endpoint will have &#x27;BRIDGEENDPOINT&#x27; value, firewall section enforced on logical router will have &#x27;LOGICALROUTER&#x27; value and rest have &#x27;VIF&#x27; value. | [optional] [default to null]
**Locked** | **bool** | Section is locked/unlocked. | [optional] [default to false]
**TcpStrict** | **bool** | If TCP strict is enabled on a section and a packet matches rule in it, the following check will be performed. If the packet does not belong to an existing session, the kernel will check to see if the SYN flag of the packet is set. If it is not, then it will drop the packet. | [optional] [default to false]
**LockModifiedBy** | **string** | ID of the user who last modified the lock for the section. | [optional] [default to null]
**LockModifiedTime** | **int64** | Section locked/unlocked time in epoch milliseconds. | [optional] [default to null]
**Priority** | **int64** | Priority of current section with respect to other sections. In case the field is empty, the list section api should be used to get section priority. | [optional] [default to null]
**FirewallSchedule** | [***ResourceReference**](ResourceReference.md) |  | [optional] [default to null]
**Comments** | **string** | Comments for section lock/unlock. | [optional] [default to null]
**Autoplumbed** | **bool** | This flag indicates whether it is an auto-plumbed section that is associated to a LogicalRouter. Auto-plumbed sections are system owned and cannot be updated via the API. | [optional] [default to false]
**Rules** | [**[]FirewallRule**](FirewallRule.md) | List of firewall rules in the section. Only homogenous rules are supported. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

