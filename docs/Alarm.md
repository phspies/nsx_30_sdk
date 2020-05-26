# Alarm

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemOwned** | **bool** | Indicates system owned resource | [optional] [default to null]
**DisplayName** | **string** | Defaults to ID if not set | [optional] [default to null]
**Description** | **string** | Detailed description of Alarm. This is the same detailed description as the corresponding Event identified by feature_name.event_type.  | [optional] [default to null]
**Tags** | [**[]Tag**](Tag.md) | Opaque identifiers meaningful to the API user | [optional] [default to null]
**CreateUser** | **string** | ID of the user who created this resource | [optional] [default to null]
**Protection** | **string** | Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed             to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed                 to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super                    user and can modify it, but only when providing                    the request header X-Allow-Overwrite&#x3D;true. UNKNOWN - the _protection field could not be determined for this           entity.  | [optional] [default to null]
**CreateTime** | **int64** | Timestamp of resource creation | [optional] [default to null]
**LastModifiedTime** | **int64** | Timestamp of last modification | [optional] [default to null]
**LastModifiedUser** | **string** | ID of the user who last modified this resource | [optional] [default to null]
**Id** | **string** | ID that uniquely identifies an Alarm.  | [optional] [default to null]
**ResourceType** | **string** | The type of this resource. | [optional] [default to null]
**LastReportedTime** | **int64** | Indicates when the corresponding Event instance was last reported in milliseconds since epoch.  | [optional] [default to null]
**Status** | **string** | Indicate the status which the Alarm is in.  | [default to null]
**EntityId** | **string** | The entity that the Event instance applies to. Note entity_id may not be included in a response body. For example, the cpu_high Event may not return an entity_id.  | [optional] [default to null]
**EventType** | **string** | Name of Event, e.g. manager_cpu_usage_high, certificate_expired.  | [optional] [default to null]
**RecommendedAction** | **string** | Recommended action for Alarm. This is the same action as the corresponding Event identified by feature_name.event_type.  | [optional] [default to null]
**Severity** | **string** | Severity of the Alarm.Can be one of - CRITICAL, HIGH, MEDIUM, LOW.  | [optional] [default to null]
**NodeId** | **string** | The UUID of the node that the Event instance applies to.  | [optional] [default to null]
**FeatureName** | **string** | Feature defining this Event, e.g. manager_health, certificates.  | [optional] [default to null]
**ResolvedBy** | **string** | User ID of the user that set the status value to RESOLVED. This value can be SYSTEM to indicate that the system resolved the Alarm, for example when the system determines CPU usage is no longer high and the cpu_high Alarm is no longer applicable. This property is only returned when the status value is RESOLVED.  | [optional] [default to null]
**EventTypeDisplayName** | **string** | Display name of Event type.  | [optional] [default to null]
**NodeDisplayName** | **string** | Display name of node that the event instance applies to.  | [optional] [default to null]
**Summary** | **string** | Summary description of Alarm. This is the same summary description as the corresponding Event identified by feature_name.event_type.  | [optional] [default to null]
**AlarmSourceType** | **string** | Type of alarm source of the Event instance. Can be one of - INTENT_PATH, ENTITY_ID.  | [optional] [default to null]
**NodeResourceType** | **string** | The resource type of node that the Event instance applies to eg. ClusterNodeConfig, HostNode, EdgeNode.  | [optional] [default to null]
**AlarmSource** | **[]string** | If alarm_source_type &#x3D; INTENT_PATH, this field will contain a list of intent paths for the entity that the event instance applies to. If alarm_source_type &#x3D; ENTITY_ID, this field will contain a list with a single item identifying the entity id that the event instance applies to.  | [optional] [default to null]
**FeatureDisplayName** | **string** | Display name of feature defining this Event.  | [optional] [default to null]
**SuppressedBy** | **string** | User ID of the user that set the status value to SUPPRESSED. This property is only returned when the status value is SUPPRESSED.  | [optional] [default to null]
**SuppressStartTime** | **int64** | Indicates when the Alarm was suppressed in milliseconds since epoch. This property is only returned when the status value is SUPPRESSED.  | [optional] [default to null]
**ResolvedTime** | **int64** | Indicates when the Alarm was resolved in milliseconds since epoch. This property is only returned when the status value is RESOLVED.  | [optional] [default to null]
**EntityResourceType** | **string** | The entity type that the Event instance applies to.  | [optional] [default to null]
**SuppressDuration** | **int64** | The time period between suppress_start_time and suppress_start_time + suppress_duration (specified in hours) an Alarm is SUPPRESSED. This property is only returned when the status value is SUPPRESSED.  | [optional] [default to null]
**NodeIpAddresses** | **[]string** | IP addresses of node that the event instance applies to.  | [optional] [default to null]
**ReoccurrencesWhileSuppressed** | **int64** | The number of reoccurrences since this alarm has been SUPPRESSED.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

