# {{classname}}

All URIs are relative to *https://nsxmanager.your.domain/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNotificationWatcher**](SystemAdministrationMonitoringApi.md#AddNotificationWatcher) | **Post** /notification-watchers | Add a new notification watcher.
[**AddUriFiltersAddUriFilters**](SystemAdministrationMonitoringApi.md#AddUriFiltersAddUriFilters) | **Post** /notification-watchers/{watcher-id}/notifications?action&#x3D;add_uri_filters | Add uri filters for the specified watcher ID.
[**BulkUpdateAlarmsSetStatus**](SystemAdministrationMonitoringApi.md#BulkUpdateAlarmsSetStatus) | **Post** /alarms?action&#x3D;set_status | Bulk update the status of zero or more Alarms.
[**CollectAlarms**](SystemAdministrationMonitoringApi.md#CollectAlarms) | **Get** /hpm/alarms | Collect alarms from all NSX nodes
[**CollectAuditLogs**](SystemAdministrationMonitoringApi.md#CollectAuditLogs) | **Post** /administration/audit-logs | Collect audit logs from registered manager nodes
[**CreateContainerClusterStatus**](SystemAdministrationMonitoringApi.md#CreateContainerClusterStatus) | **Post** /systemhealth/container-cluster/ncp/status | Create container cluster status list
[**CreateLatencyStatProfile**](SystemAdministrationMonitoringApi.md#CreateLatencyStatProfile) | **Post** /latency-profiles | Create a new latency profile
[**CreateSystemHealthAgentProfile**](SystemAdministrationMonitoringApi.md#CreateSystemHealthAgentProfile) | **Post** /systemhealth/profiles | Create a system health profile
[**DeleteContainerClusterSummary**](SystemAdministrationMonitoringApi.md#DeleteContainerClusterSummary) | **Delete** /systemhealth/container-cluster/{cluster-id}/ncp/status | Create container cluster status list
[**DeleteLatencyStatProfile**](SystemAdministrationMonitoringApi.md#DeleteLatencyStatProfile) | **Delete** /latency-profiles/{latency-profile-id} | Delete an existing latency profile
[**DeleteNotificationWatcher**](SystemAdministrationMonitoringApi.md#DeleteNotificationWatcher) | **Delete** /notification-watchers/{watcher-id} | Delete an existing Notification-Watcher.
[**DeleteSystemHealthAgentProfile**](SystemAdministrationMonitoringApi.md#DeleteSystemHealthAgentProfile) | **Delete** /systemhealth/profiles/{profile-id} | Delete an existing system health profile
[**DeleteUriFiltersDeleteUriFilters**](SystemAdministrationMonitoringApi.md#DeleteUriFiltersDeleteUriFilters) | **Post** /notification-watchers/{watcher-id}/notifications?action&#x3D;delete_uri_filters | Delete uri filters for the specified watcher ID.
[**GetAggregationServiceGlobalConfig**](SystemAdministrationMonitoringApi.md#GetAggregationServiceGlobalConfig) | **Get** /hpm/global-config | Read global health performance monitoring configuration
[**GetAlarm**](SystemAdministrationMonitoringApi.md#GetAlarm) | **Get** /alarms/{alarm-id} | Get Alarm identified by alarm-id.
[**GetAlarms**](SystemAdministrationMonitoringApi.md#GetAlarms) | **Get** /alarms | Get the list of all Alarms currently known to the system.
[**GetAutomaticHealthCheck**](SystemAdministrationMonitoringApi.md#GetAutomaticHealthCheck) | **Get** /automatic-health-checks/transport-zones/{transport-zone-id} | Get an automatic health check
[**GetAutomaticHealthCheckToggle**](SystemAdministrationMonitoringApi.md#GetAutomaticHealthCheckToggle) | **Get** /automatic-health-check-toggle | Get automatic health check toggle
[**GetErrorResolverInfo**](SystemAdministrationMonitoringApi.md#GetErrorResolverInfo) | **Get** /error-resolver/{error_id} | Fetches metadata about the given error_id
[**GetEvent**](SystemAdministrationMonitoringApi.md#GetEvent) | **Get** /events/{event-id} | Get Events identified by event-id.
[**GetFeatureStackConfiguration**](SystemAdministrationMonitoringApi.md#GetFeatureStackConfiguration) | **Get** /hpm/features/{feature-stack-name} | Read health performance monitoring configuration for feature stack
[**GetNotificationWatcher**](SystemAdministrationMonitoringApi.md#GetNotificationWatcher) | **Get** /notification-watchers/{watcher-id} | Returns notification watcher by watcher id.
[**GetNotifications**](SystemAdministrationMonitoringApi.md#GetNotifications) | **Get** /notification-watchers/{watcher-id}/notifications | Get notifications for the specified watcher ID.
[**ListAutomaticHealthChecks**](SystemAdministrationMonitoringApi.md#ListAutomaticHealthChecks) | **Get** /automatic-health-checks | List automatic health checks
[**ListErrorResolverInfo**](SystemAdministrationMonitoringApi.md#ListErrorResolverInfo) | **Get** /error-resolver | Fetches a list of metadata for all the registered error resolvers
[**ListEvents**](SystemAdministrationMonitoringApi.md#ListEvents) | **Get** /events | Get the list of all Events defined in NSX.
[**ListFeatureStackConfigurations**](SystemAdministrationMonitoringApi.md#ListFeatureStackConfigurations) | **Get** /hpm/features | List all health performance monitoring feature stacks
[**ListLatencyProfiles**](SystemAdministrationMonitoringApi.md#ListLatencyProfiles) | **Get** /latency-profiles | List latency profiles
[**ListNodeLogs**](SystemAdministrationMonitoringApi.md#ListNodeLogs) | **Get** /node/logs | List available node logs
[**ListNotificationWatchers**](SystemAdministrationMonitoringApi.md#ListNotificationWatchers) | **Get** /notification-watchers | Returns a list of registered notification watchers.
[**ListSystemHealthAgentProfiles**](SystemAdministrationMonitoringApi.md#ListSystemHealthAgentProfiles) | **Get** /systemhealth/profiles | List all system health profiles
[**ReadContainerClusterStatus**](SystemAdministrationMonitoringApi.md#ReadContainerClusterStatus) | **Get** /systemhealth/container-cluster/{cluster-id}/ncp/status | Get the container cluster status by given id
[**ReadContainerClusterStatusList**](SystemAdministrationMonitoringApi.md#ReadContainerClusterStatusList) | **Get** /systemhealth/container-cluster/ncp/status | Get all the container cluster status
[**ReadLatencyStatProfile**](SystemAdministrationMonitoringApi.md#ReadLatencyStatProfile) | **Get** /latency-profiles/{latency-profile-id} | Get an existing latency profile configuration
[**ReadNodeLog**](SystemAdministrationMonitoringApi.md#ReadNodeLog) | **Get** /node/logs/{log-name} | Read node log properties
[**ReadNodeLogData**](SystemAdministrationMonitoringApi.md#ReadNodeLogData) | **Get** /node/logs/{log-name}/data | Read node log contents
[**ReadTnContainerAgentStatus**](SystemAdministrationMonitoringApi.md#ReadTnContainerAgentStatus) | **Get** /systemhealth/transport-nodes/{node-id}/container/agent/status | Get the containter status on given node
[**ReadTnHyperbusStatus**](SystemAdministrationMonitoringApi.md#ReadTnHyperbusStatus) | **Get** /systemhealth/transport-nodes/{node-id}/container/hyperbus/status | Get the containter hyperbus status on given node
[**ResetAggregationServiceFeatureStackConfigurationResetCollectionFrequency**](SystemAdministrationMonitoringApi.md#ResetAggregationServiceFeatureStackConfigurationResetCollectionFrequency) | **Post** /hpm/features/{feature-stack-name}?action&#x3D;reset_collection_frequency | Reset the data collection frequency configuration setting to the default values
[**ResetEventValuesSetDefault**](SystemAdministrationMonitoringApi.md#ResetEventValuesSetDefault) | **Post** /events/{event-id}?action&#x3D;set_default | Reset all user configurable values to factory defaults.
[**ResolveErrorResolveError**](SystemAdministrationMonitoringApi.md#ResolveErrorResolveError) | **Post** /error-resolver?action&#x3D;resolve_error | Resolves the error
[**ShowSystemHealthAgentProfile**](SystemAdministrationMonitoringApi.md#ShowSystemHealthAgentProfile) | **Get** /systemhealth/profiles/{profile-id} | Show the details of a system health profile
[**UpdateAggregationServiceGlobalConfig**](SystemAdministrationMonitoringApi.md#UpdateAggregationServiceGlobalConfig) | **Put** /hpm/global-config | Set the global configuration for aggregation service related data collection
[**UpdateAlarmStatusSetStatus**](SystemAdministrationMonitoringApi.md#UpdateAlarmStatusSetStatus) | **Post** /alarms/{alarm-id}?action&#x3D;set_status | Update staus of alarm identified by alarm-id.
[**UpdateAutomaticHealthCheckToggle**](SystemAdministrationMonitoringApi.md#UpdateAutomaticHealthCheckToggle) | **Put** /automatic-health-check-toggle | Update automatic health check toggle
[**UpdateEvent**](SystemAdministrationMonitoringApi.md#UpdateEvent) | **Put** /events/{event-id} | Update event associated with event-id.
[**UpdateFeatureStackConfiguration**](SystemAdministrationMonitoringApi.md#UpdateFeatureStackConfiguration) | **Put** /hpm/features/{feature-stack-name} | Update health performance monitoring configuration for feature stack
[**UpdateLatencyProfile**](SystemAdministrationMonitoringApi.md#UpdateLatencyProfile) | **Put** /latency-profiles/{latency-profile-id} | Update an existing latency profile
[**UpdateNotificationWatcher**](SystemAdministrationMonitoringApi.md#UpdateNotificationWatcher) | **Put** /notification-watchers/{watcher-id} | Update notification watcher.
[**UpdateNotifications**](SystemAdministrationMonitoringApi.md#UpdateNotifications) | **Put** /notification-watchers/{watcher-id}/notifications | Update notifications for the specified watcher ID.
[**UpdateSystemHealthAgentProfile**](SystemAdministrationMonitoringApi.md#UpdateSystemHealthAgentProfile) | **Put** /systemhealth/profiles/{profile-id} | Update a system health profile

# **AddNotificationWatcher**
> NotificationWatcher AddNotificationWatcher(ctx, body)
Add a new notification watcher.

Add a new notification watcher.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NotificationWatcher**](NotificationWatcher.md)|  | 

### Return type

[**NotificationWatcher**](NotificationWatcher.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddUriFiltersAddUriFilters**
> NotificationsList AddUriFiltersAddUriFilters(ctx, body, watcherId)
Add uri filters for the specified watcher ID.

Add uri filters for the specified watcher ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Notification**](Notification.md)|  | 
  **watcherId** | **string**|  | 

### Return type

[**NotificationsList**](NotificationsList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkUpdateAlarmsSetStatus**
> BulkUpdateAlarmsSetStatus(ctx, newStatus, optional)
Bulk update the status of zero or more Alarms.

Bulk update the status of zero or more Alarms that match the specified filters. The new_status value can be OPEN, ACKNOWLEDGED, SUPPRESSED, or RESOLVED. If new_status is SUPPRESSED, the suppress_duration query parameter must also be specified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newStatus** | **string**| Status | 
 **optional** | ***SystemAdministrationMonitoringApiBulkUpdateAlarmsSetStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationMonitoringApiBulkUpdateAlarmsSetStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **optional.Int64**| Timestamp in milliseconds since epoch | 
 **before** | **optional.Int64**| Timestamp in milliseconds since epoch | 
 **cursor** | **optional.String**| Cursor for pagination | 
 **eventType** | **optional.String**| Event Type Filter | 
 **featureName** | **optional.String**| Feature Name | 
 **id** | **optional.String**| Alarm ID | 
 **intentPath** | **optional.String**| Intent Path for entity ID | 
 **nodeId** | **optional.String**| Node ID | 
 **nodeResourceType** | **optional.String**| Node Resource Type | 
 **pageSize** | **optional.Int64**| Page Size for pagination | 
 **severity** | **optional.String**| Severity | 
 **sortAscending** | **optional.Bool**| Represents order of sorting the values | [default to true]
 **sortBy** | **optional.String**| Key for sorting on this column | 
 **status** | **optional.String**| Status | 
 **suppressDuration** | **optional.Int64**| Duration in hours for which Alarm should be suppressed | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CollectAlarms**
> AlarmListResult CollectAlarms(ctx, optional)
Collect alarms from all NSX nodes

This API is executed on a manager node to return current alarms from all NSX nodes. This API is deprecated as part of alarm framework enhancements. Please use below new APIs: GET /alarms GET /alarms/<alarm-id> POST /alarms/<alarm-id>?action=set_status POST /alarms?action=set_status POST /alarms?action=set_and_verify_status GET /alarms/resource-types GET /events GET /events/<event-id> PUT /events/<event-id> POST /events/<event-id>?action=set_default 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationMonitoringApiCollectAlarmsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationMonitoringApiCollectAlarmsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.Int64**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **fields** | **optional.String**| Fields to include in query results | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 100]

### Return type

[**AlarmListResult**](AlarmListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CollectAuditLogs**
> AuditLogListResult CollectAuditLogs(ctx, body, optional)
Collect audit logs from registered manager nodes

This API is executed on a manager node to display audit logs from all nodes inside the management plane cluster. An audit log collection will be triggered if the local master audit log is outdated. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AuditLogRequest**](AuditLogRequest.md)|  | 
 **optional** | ***SystemAdministrationMonitoringApiCollectAuditLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationMonitoringApiCollectAuditLogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **fields** | **optional.**| Fields to include in query results | 
 **pageSize** | **optional.**| Maximum number of results to return in this page (server may return fewer) | [default to 100]

### Return type

[**AuditLogListResult**](AuditLogListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateContainerClusterStatus**
> ContainerClusterStatus CreateContainerClusterStatus(ctx, body)
Create container cluster status list

Create container cluster status list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ContainerClusterStatus**](ContainerClusterStatus.md)|  | 

### Return type

[**ContainerClusterStatus**](ContainerClusterStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLatencyStatProfile**
> LatencyStatProfile CreateLatencyStatProfile(ctx, body)
Create a new latency profile

Create a new latency profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LatencyStatProfile**](LatencyStatProfile.md)|  | 

### Return type

[**LatencyStatProfile**](LatencyStatProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSystemHealthAgentProfile**
> SystemHealthAgentProfile CreateSystemHealthAgentProfile(ctx, body)
Create a system health profile

Create a system health profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SystemHealthAgentProfile**](SystemHealthAgentProfile.md)|  | 

### Return type

[**SystemHealthAgentProfile**](SystemHealthAgentProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContainerClusterSummary**
> DeleteContainerClusterSummary(ctx, clusterId)
Create container cluster status list

Create container cluster status list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLatencyStatProfile**
> DeleteLatencyStatProfile(ctx, latencyProfileId)
Delete an existing latency profile

Delete an existing latency profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **latencyProfileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNotificationWatcher**
> DeleteNotificationWatcher(ctx, watcherId)
Delete an existing Notification-Watcher.

Delete notification watcher. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **watcherId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSystemHealthAgentProfile**
> DeleteSystemHealthAgentProfile(ctx, profileId)
Delete an existing system health profile

Delete an existing system health profile by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **profileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUriFiltersDeleteUriFilters**
> NotificationsList DeleteUriFiltersDeleteUriFilters(ctx, body, watcherId)
Delete uri filters for the specified watcher ID.

Delete uri filters for the specified watcher ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Notification**](Notification.md)|  | 
  **watcherId** | **string**|  | 

### Return type

[**NotificationsList**](NotificationsList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAggregationServiceGlobalConfig**
> GlobalCollectionConfiguration GetAggregationServiceGlobalConfig(ctx, )
Read global health performance monitoring configuration

Read global health performance monitoring configuration

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GlobalCollectionConfiguration**](GlobalCollectionConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlarm**
> Alarm GetAlarm(ctx, alarmId)
Get Alarm identified by alarm-id.

Returns alarm associated with alarm-id. If HTTP status 404 is returned, this means the specified alarm-id is invalid or the alarm with alarm-id has been deleted. An alarm is deleted by the system if it is RESOLVED and older than eight days. The system can also delete the remaining RESOLVED alarms sooner to free system resources when too many alarms are being generated. When this happens the oldest day's RESOLVED alarms are deleted first. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alarmId** | **string**|  | 

### Return type

[**Alarm**](Alarm.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlarms**
> AlarmsListResult GetAlarms(ctx, optional)
Get the list of all Alarms currently known to the system.

Returns a list of all Alarms currently known to the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationMonitoringApiGetAlarmsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationMonitoringApiGetAlarmsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **optional.Int64**| Timestamp in milliseconds since epoch | 
 **before** | **optional.Int64**| Timestamp in milliseconds since epoch | 
 **cursor** | **optional.String**| Cursor for pagination | 
 **eventType** | **optional.String**| Event Type Filter | 
 **featureName** | **optional.String**| Feature Name | 
 **id** | **optional.String**| Alarm ID | 
 **intentPath** | **optional.String**| Intent Path for entity ID | 
 **nodeId** | **optional.String**| Node ID | 
 **nodeResourceType** | **optional.String**| Node Resource Type | 
 **pageSize** | **optional.Int64**| Page Size for pagination | 
 **severity** | **optional.String**| Severity | 
 **sortAscending** | **optional.Bool**| Represents order of sorting the values | [default to true]
 **sortBy** | **optional.String**| Key for sorting on this column | 
 **status** | **optional.String**| Status | 

### Return type

[**AlarmsListResult**](AlarmsListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAutomaticHealthCheck**
> AutomaticHealthCheck GetAutomaticHealthCheck(ctx, transportZoneId)
Get an automatic health check

Get health check performed by system automatically for specific transport zone. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transportZoneId** | **string**|  | 

### Return type

[**AutomaticHealthCheck**](AutomaticHealthCheck.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAutomaticHealthCheckToggle**
> AutomaticHealthCheckToggle GetAutomaticHealthCheckToggle(ctx, )
Get automatic health check toggle

Get detailed info for automatic health check toggle.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AutomaticHealthCheckToggle**](AutomaticHealthCheckToggle.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetErrorResolverInfo**
> ErrorResolverInfo GetErrorResolverInfo(ctx, errorId)
Fetches metadata about the given error_id

Returns some metadata about the given error_id. This includes information of whether there is a resolver present for the given error_id and its associated user input data 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **errorId** | **string**|  | 

### Return type

[**ErrorResolverInfo**](ErrorResolverInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEvent**
> MonitoringEvent GetEvent(ctx, eventId)
Get Events identified by event-id.

Returns event associated with event-id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventId** | **string**|  | 

### Return type

[**MonitoringEvent**](MonitoringEvent.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatureStackConfiguration**
> FeatureStackCollectionConfiguration GetFeatureStackConfiguration(ctx, featureStackName)
Read health performance monitoring configuration for feature stack

Returns the complete set of client type data collection configuration records for the specified feature stack. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **featureStackName** | **string**|  | 

### Return type

[**FeatureStackCollectionConfiguration**](FeatureStackCollectionConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNotificationWatcher**
> NotificationWatcher GetNotificationWatcher(ctx, watcherId)
Returns notification watcher by watcher id.

Returns notification watcher by watcher id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **watcherId** | **string**|  | 

### Return type

[**NotificationWatcher**](NotificationWatcher.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNotifications**
> NotificationsList GetNotifications(ctx, watcherId)
Get notifications for the specified watcher ID.

Get notifications for the specified watcher ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **watcherId** | **string**|  | 

### Return type

[**NotificationsList**](NotificationsList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAutomaticHealthChecks**
> AutomaticHealthCheckListResult ListAutomaticHealthChecks(ctx, optional)
List automatic health checks

Query automatic health checks with list parameters. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationMonitoringApiListAutomaticHealthChecksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationMonitoringApiListAutomaticHealthChecksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**AutomaticHealthCheckListResult**](AutomaticHealthCheckListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListErrorResolverInfo**
> ErrorResolverInfoList ListErrorResolverInfo(ctx, )
Fetches a list of metadata for all the registered error resolvers

Returns a list of metadata for all the error resolvers registered. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ErrorResolverInfoList**](ErrorResolverInfoList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEvents**
> EventListResult ListEvents(ctx, )
Get the list of all Events defined in NSX.

Returns a list of all Events defined in NSX.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EventListResult**](EventListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFeatureStackConfigurations**
> FeatureStackCollectionConfigurationList ListFeatureStackConfigurations(ctx, )
List all health performance monitoring feature stacks

List all health performance monitoring feature stacks

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FeatureStackCollectionConfigurationList**](FeatureStackCollectionConfigurationList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLatencyProfiles**
> LatencyStatProfileListResult ListLatencyProfiles(ctx, optional)
List latency profiles

List latency profiles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationMonitoringApiListLatencyProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationMonitoringApiListLatencyProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**LatencyStatProfileListResult**](LatencyStatProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNodeLogs**
> NodeLogPropertiesListResult ListNodeLogs(ctx, )
List available node logs

Returns the number of log files and lists the log files that reside on the NSX virtual appliance. The list includes the filename, file size, and last-modified time in milliseconds since epoch (1 January 1970) for each log file. Knowing the last-modified time with millisecond accuracy since epoch is helpful when you are comparing two times, such as the time of a POST request and the end time on a server. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeLogPropertiesListResult**](NodeLogPropertiesListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNotificationWatchers**
> NotificationWatcherListResult ListNotificationWatchers(ctx, )
Returns a list of registered notification watchers.

Returns a list of registered notification watchers.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NotificationWatcherListResult**](NotificationWatcherListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSystemHealthAgentProfiles**
> SystemHealthAgentProfileListResult ListSystemHealthAgentProfiles(ctx, )
List all system health profiles

List all system health profiles. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SystemHealthAgentProfileListResult**](SystemHealthAgentProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadContainerClusterStatus**
> ContainerClusterSummary ReadContainerClusterStatus(ctx, clusterId)
Get the container cluster status by given id

Get the container cluster status by given id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 

### Return type

[**ContainerClusterSummary**](ContainerClusterSummary.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadContainerClusterStatusList**
> ContainerClusterStatusList ReadContainerClusterStatusList(ctx, optional)
Get all the container cluster status

Get all the container cluster status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationMonitoringApiReadContainerClusterStatusListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationMonitoringApiReadContainerClusterStatusListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ContainerClusterStatusList**](ContainerClusterStatusList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadLatencyStatProfile**
> LatencyStatProfile ReadLatencyStatProfile(ctx, latencyProfileId)
Get an existing latency profile configuration

Get an existing latency profile configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **latencyProfileId** | **string**|  | 

### Return type

[**LatencyStatProfile**](LatencyStatProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNodeLog**
> NodeLogProperties ReadNodeLog(ctx, logName)
Read node log properties

For a single specified log file, lists the filename, file size, and last-modified time. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logName** | **string**| Name of log file to read properties | 

### Return type

[**NodeLogProperties**](NodeLogProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNodeLogData**
> ReadNodeLogData(ctx, logName)
Read node log contents

For a single specified log file, returns the content of the log file. This method supports byte-range requests. To request just a portion of a log file, supply an HTTP Range header, e.g. \"Range: bytes=<start>-<end>\". <end> is optional, and, if omitted, the file contents from start to the end of the file are returned.' 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logName** | **string**| Name of log to read | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadTnContainerAgentStatus**
> TnNodeAgentStatusListResult ReadTnContainerAgentStatus(ctx, nodeId, optional)
Get the containter status on given node

Get the containter status on given node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 
 **optional** | ***SystemAdministrationMonitoringApiReadTnContainerAgentStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationMonitoringApiReadTnContainerAgentStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**TnNodeAgentStatusListResult**](TnNodeAgentStatusListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadTnHyperbusStatus**
> TnHyperbusStatus ReadTnHyperbusStatus(ctx, nodeId)
Get the containter hyperbus status on given node

Get the containter hyperbus status on given node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

[**TnHyperbusStatus**](TnHyperbusStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetAggregationServiceFeatureStackConfigurationResetCollectionFrequency**
> FeatureStackCollectionConfiguration ResetAggregationServiceFeatureStackConfigurationResetCollectionFrequency(ctx, featureStackName)
Reset the data collection frequency configuration setting to the default values

Reset the data collection frequency configuration setting to the default values

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **featureStackName** | **string**|  | 

### Return type

[**FeatureStackCollectionConfiguration**](FeatureStackCollectionConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetEventValuesSetDefault**
> MonitoringEvent ResetEventValuesSetDefault(ctx, eventId)
Reset all user configurable values to factory defaults.

Reset all user configurable values for event identified by event-id to factory defaults.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventId** | **string**|  | 

### Return type

[**MonitoringEvent**](MonitoringEvent.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResolveErrorResolveError**
> ResolveErrorResolveError(ctx, body)
Resolves the error

Invokes the corresponding error resolver for the given error(s) present in the payload 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ErrorResolverMetadataList**](ErrorResolverMetadataList.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowSystemHealthAgentProfile**
> SystemHealthAgentProfile ShowSystemHealthAgentProfile(ctx, profileId)
Show the details of a system health profile

Show the details of a system health profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **profileId** | **string**|  | 

### Return type

[**SystemHealthAgentProfile**](SystemHealthAgentProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAggregationServiceGlobalConfig**
> GlobalCollectionConfiguration UpdateAggregationServiceGlobalConfig(ctx, body)
Set the global configuration for aggregation service related data collection

Set the global configuration for aggregation service related data collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GlobalCollectionConfiguration**](GlobalCollectionConfiguration.md)|  | 

### Return type

[**GlobalCollectionConfiguration**](GlobalCollectionConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAlarmStatusSetStatus**
> Alarm UpdateAlarmStatusSetStatus(ctx, alarmId, newStatus, optional)
Update staus of alarm identified by alarm-id.

Update status of an Alarm. The new_status value can be OPEN, ACKNOWLEDGED, SUPPRESSED, or RESOLVED. If new_status is SUPPRESSED, the suppress_duration query parameter must also be specified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alarmId** | **string**|  | 
  **newStatus** | **string**| Status | 
 **optional** | ***SystemAdministrationMonitoringApiUpdateAlarmStatusSetStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationMonitoringApiUpdateAlarmStatusSetStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **suppressDuration** | **optional.Int64**| Duration in hours for which Alarm should be suppressed | 

### Return type

[**Alarm**](Alarm.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAutomaticHealthCheckToggle**
> AutomaticHealthCheckToggle UpdateAutomaticHealthCheckToggle(ctx, body)
Update automatic health check toggle

Change status of automatic health check toggle to enabled/disabled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AutomaticHealthCheckToggle**](AutomaticHealthCheckToggle.md)|  | 

### Return type

[**AutomaticHealthCheckToggle**](AutomaticHealthCheckToggle.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEvent**
> MonitoringEvent UpdateEvent(ctx, body, eventId)
Update event associated with event-id.

Update event identified by event-id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MonitoringEvent**](MonitoringEvent.md)|  | 
  **eventId** | **string**|  | 

### Return type

[**MonitoringEvent**](MonitoringEvent.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFeatureStackConfiguration**
> FeatureStackCollectionConfiguration UpdateFeatureStackConfiguration(ctx, body, featureStackName)
Update health performance monitoring configuration for feature stack

Apply the data collection configuration for the specified feature stack. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FeatureStackCollectionConfiguration**](FeatureStackCollectionConfiguration.md)|  | 
  **featureStackName** | **string**|  | 

### Return type

[**FeatureStackCollectionConfiguration**](FeatureStackCollectionConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLatencyProfile**
> LatencyStatProfile UpdateLatencyProfile(ctx, body, latencyProfileId)
Update an existing latency profile

Update an existing latency profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LatencyStatProfile**](LatencyStatProfile.md)|  | 
  **latencyProfileId** | **string**|  | 

### Return type

[**LatencyStatProfile**](LatencyStatProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNotificationWatcher**
> NotificationWatcher UpdateNotificationWatcher(ctx, body, watcherId)
Update notification watcher.

Update notification watcher.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NotificationWatcher**](NotificationWatcher.md)|  | 
  **watcherId** | **string**|  | 

### Return type

[**NotificationWatcher**](NotificationWatcher.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNotifications**
> NotificationsList UpdateNotifications(ctx, body, watcherId)
Update notifications for the specified watcher ID.

Update notifications for the specified watcher ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NotificationsList**](NotificationsList.md)|  | 
  **watcherId** | **string**|  | 

### Return type

[**NotificationsList**](NotificationsList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSystemHealthAgentProfile**
> SystemHealthAgentProfile UpdateSystemHealthAgentProfile(ctx, body, profileId)
Update a system health profile

Update a system health profile definition. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SystemHealthAgentProfile**](SystemHealthAgentProfile.md)|  | 
  **profileId** | **string**|  | 

### Return type

[**SystemHealthAgentProfile**](SystemHealthAgentProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

