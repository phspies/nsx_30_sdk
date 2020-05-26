# FirewallSection

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stateful** | **bool** | Stateful or Stateless nature of distributed service section is enforced on all rules inside the section. Layer3 sections can be stateful or stateless. Layer2 sections can only be stateless. | [default to null]
**IsDefault** | **bool** | It is a boolean flag which reflects whether a distributed service section is default section or not. Each Layer 3 and Layer 2 section will have at least and at most one default section. | [optional] [default to null]
**AppliedTos** | [**[]ResourceReference**](ResourceReference.md) | List of objects where the rules in this section will be enforced. This will take precedence over rule level appliedTo. | [optional] [default to null]
**RuleCount** | **int64** | Number of rules in this section. | [optional] [default to null]
**SectionType** | **string** | Type of the rules which a section can contain. Only homogeneous sections are supported. | [default to null]
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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

