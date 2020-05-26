# MonitoringEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemOwned** | **bool** | Indicates system owned resource | [optional] [default to null]
**DisplayName** | **string** | Defaults to ID if not set | [optional] [default to null]
**Description** | **string** | Detailed description of the event.  | [optional] [default to null]
**Tags** | [**[]Tag**](Tag.md) | Opaque identifiers meaningful to the API user | [optional] [default to null]
**CreateUser** | **string** | ID of the user who created this resource | [optional] [default to null]
**Protection** | **string** | Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed             to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed                 to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super                    user and can modify it, but only when providing                    the request header X-Allow-Overwrite&#x3D;true. UNKNOWN - the _protection field could not be determined for this           entity.  | [optional] [default to null]
**CreateTime** | **int64** | Timestamp of resource creation | [optional] [default to null]
**LastModifiedTime** | **int64** | Timestamp of last modification | [optional] [default to null]
**LastModifiedUser** | **string** | ID of the user who last modified this resource | [optional] [default to null]
**Id** | **string** | Unique identifier in the form of feature_name.event_type.  | [optional] [default to null]
**ResourceType** | **string** | The type of this resource. | [optional] [default to null]
**IsThresholdFixed** | **bool** | Indicates if the threshold property is configurable via the API.  | [optional] [default to null]
**EventFalseSnmpOid** | **string** | Optional field containing OID for SNMP trap sent when Event instance is False. This value is null if suppress_snmp_trap or suppress_clear_oid is True.  | [optional] [default to null]
**DescriptionOnClear** | **string** | Description of Event when an Event instance transitions from True to False.  | [optional] [default to null]
**EventType** | **string** | Name of Event, e.g. manager_cpu_usage_high, certificate_expired.  | [optional] [default to null]
**RecommendedAction** | **string** | Recommended action for the alarm condition.  | [optional] [default to null]
**Severity** | **string** | Severity of the Event.Can be one of - CRITICAL, HIGH, MEDIUM, LOW.  | [optional] [default to null]
**Sensitivity** | **int64** | Percentage of samples to consider and used in combination with threshold when determining whether an Event instance status is True or False. Event evaluation uses sampling to determine Event instance status. A higher sensitivity value specifies that more samples are used to ensure accuracy and ignore infrequent or rare spikes in sampled data.  | [default to null]
**IsDisabled** | **bool** | Flag to indicate whether sampling for this Event is off or on.  | [optional] [default to false]
**SuppressAlarm** | **bool** | Flag to suppress Alarm generation. Alarms are not generated for this Event when this is set to True.  | [optional] [default to false]
**Summary** | **string** | Summary description of the event.  | [optional] [default to null]
**FeatureDisplayName** | **string** | Display name of feature defining this Event.  | [optional] [default to null]
**EntityResourceType** | **string** | Resource Type of entity where this event is applicable eg. LogicalSwitch, LogicalPort etc.  | [optional] [default to null]
**FeatureName** | **string** | Feature defining this Event, e.g. manager_health, certificates.  | [optional] [default to null]
**Threshold** | **int64** | Threshold to determine if a single sample is True. For example, if the configured threshold is 95% and the current CPU sample is 99%, then the current sample is considered True.  | [default to null]
**EventTypeDisplayName** | **string** | Display name of Event type.  | [optional] [default to null]
**SuppressSnmpTrap** | **bool** | Flag to suppress SNMP trap generation. SNMP traps are not sent for this Event when this is set to True.  | [optional] [default to false]
**EventTrueSnmpOid** | **string** | Optional field containing OID for SNMP trap sent when Event instance is True. This value is null if suppress_snmp_trap is True.  | [optional] [default to null]
**NodeTypes** | **[]string** | Array identifying the nodes on which this Event is applicable. Can be one or more of the following values - nsx_public_cloud_gateway, nsx_edge, nsx_esx, nsx_kvm, nsx_manager.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

