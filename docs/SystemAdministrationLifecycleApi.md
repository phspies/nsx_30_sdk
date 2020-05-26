# {{classname}}

All URIs are relative to *https://nsxmanager.your.domain/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AbortMigrationAbort**](SystemAdministrationLifecycleApi.md#AbortMigrationAbort) | **Post** /migration/plan?action&#x3D;abort | Abort migration
[**AbortPreUpgradeChecksAbortPreUpgradeChecks**](SystemAdministrationLifecycleApi.md#AbortPreUpgradeChecksAbortPreUpgradeChecks) | **Post** /upgrade?action&#x3D;abort_pre_upgrade_checks | Abort pre-upgrade checks
[**AcceptRecommendedValueInFeedbackAcceptRecommended**](SystemAdministrationLifecycleApi.md#AcceptRecommendedValueInFeedbackAcceptRecommended) | **Post** /migration/feedback-response?action&#x3D;accept-recommended | Accept default action for feedbacks
[**AcceptUpgradeEULA**](SystemAdministrationLifecycleApi.md#AcceptUpgradeEULA) | **Post** /upgrade/eula/accept | Accept end user license agreement 
[**AddMigrationUnitsToGroupAddMigrationUnits**](SystemAdministrationLifecycleApi.md#AddMigrationUnitsToGroupAddMigrationUnits) | **Post** /migration/migration-unit-groups/{group-id}?action&#x3D;add_migration_units | Add migration units to specified migration unit group
[**AddUpgradeUnitsToGroupAddUpgradeUnits**](SystemAdministrationLifecycleApi.md#AddUpgradeUnitsToGroupAddUpgradeUnits) | **Post** /upgrade/upgrade-unit-groups/{group-id}?action&#x3D;add_upgrade_units | Add upgrade units to specified upgrade unit group
[**AdvanceClusterRestoreAdvance**](SystemAdministrationLifecycleApi.md#AdvanceClusterRestoreAdvance) | **Post** /cluster/restore?action&#x3D;advance | Advance any suspended restore operation
[**CancelClusterRestoreCancel**](SystemAdministrationLifecycleApi.md#CancelClusterRestoreCancel) | **Post** /cluster/restore?action&#x3D;cancel | Cancel any running restore operation
[**CancelUpgradeBundleUploadCancelUpload**](SystemAdministrationLifecycleApi.md#CancelUpgradeBundleUploadCancelUpload) | **Post** /upgrade/bundles/{bundle-id}?action&#x3D;cancel_upload | Cancel upgrade bundle upload
[**ConfigureBackupConfig**](SystemAdministrationLifecycleApi.md#ConfigureBackupConfig) | **Put** /cluster/backups/config | Configure backup
[**ConfigureRestoreConfig**](SystemAdministrationLifecycleApi.md#ConfigureRestoreConfig) | **Put** /cluster/restore/config | Deprecated. Configure Restore SFTP server credentials
[**ContinueMigrationContinue**](SystemAdministrationLifecycleApi.md#ContinueMigrationContinue) | **Post** /migration/plan?action&#x3D;continue | Continue migration
[**ContinueUpgradeContinue**](SystemAdministrationLifecycleApi.md#ContinueUpgradeContinue) | **Post** /upgrade/plan?action&#x3D;continue | Continue upgrade
[**CopyFromRemoteFileCopyFromRemoteFile**](SystemAdministrationLifecycleApi.md#CopyFromRemoteFileCopyFromRemoteFile) | **Post** /node/file-store/{file-name}?action&#x3D;copy_from_remote_file | Copy a remote file to the file store
[**CopyToRemoteFileCopyToRemoteFile**](SystemAdministrationLifecycleApi.md#CopyToRemoteFileCopyToRemoteFile) | **Post** /node/file-store/{file-name}?action&#x3D;copy_to_remote_file | Copy file in the file store to a remote file store
[**CreateFile**](SystemAdministrationLifecycleApi.md#CreateFile) | **Post** /node/file-store/{file-name} | Upload a file to the file store
[**CreateMigrationUnitGroup**](SystemAdministrationLifecycleApi.md#CreateMigrationUnitGroup) | **Post** /migration/migration-unit-groups | Create a group
[**CreateRemoteDirectoryCreateRemoteDirectory**](SystemAdministrationLifecycleApi.md#CreateRemoteDirectoryCreateRemoteDirectory) | **Post** /node/file-store?action&#x3D;create_remote_directory | Create directory in remote file server
[**CreateUpgradeUnitGroup**](SystemAdministrationLifecycleApi.md#CreateUpgradeUnitGroup) | **Post** /upgrade/upgrade-unit-groups | Create a group
[**CreateView**](SystemAdministrationLifecycleApi.md#CreateView) | **Post** /ui-views | Creates a new View.
[**CreateWidgetConfiguration**](SystemAdministrationLifecycleApi.md#CreateWidgetConfiguration) | **Post** /ui-views/{view-id}/widgetconfigurations | Creates a new Widget Configuration.
[**DeletView**](SystemAdministrationLifecycleApi.md#DeletView) | **Delete** /ui-views/{view-id} | Delete View
[**DeleteFile**](SystemAdministrationLifecycleApi.md#DeleteFile) | **Delete** /node/file-store/{file-name} | Delete file
[**DeleteMigrationUnitGroup**](SystemAdministrationLifecycleApi.md#DeleteMigrationUnitGroup) | **Delete** /migration/migration-unit-groups/{group-id} | Delete the migration unit group
[**DeleteUpgradeUnitGroup**](SystemAdministrationLifecycleApi.md#DeleteUpgradeUnitGroup) | **Delete** /upgrade/upgrade-unit-groups/{group-id} | Delete the upgrade unit group
[**DeleteWidgetConfiguration**](SystemAdministrationLifecycleApi.md#DeleteWidgetConfiguration) | **Delete** /ui-views/{view-id}/widgetconfigurations/{widgetconfiguration-id} | Delete Widget Configuration
[**ExecutePostUpgradeChecksExecutePostUpgradeChecks**](SystemAdministrationLifecycleApi.md#ExecutePostUpgradeChecksExecutePostUpgradeChecks) | **Post** /upgrade/{component-type}?action&#x3D;execute_post_upgrade_checks | Execute post-upgrade checks
[**ExecutePreUpgradeChecksExecutePreUpgradeChecks**](SystemAdministrationLifecycleApi.md#ExecutePreUpgradeChecksExecutePreUpgradeChecks) | **Post** /upgrade?action&#x3D;execute_pre_upgrade_checks | Execute pre-upgrade checks
[**ExecuteUpgradeTask_**](SystemAdministrationLifecycleApi.md#ExecuteUpgradeTask_) | **Post** /node/upgrade/performtask?action&#x3D;[^/]+ | Execute upgrade task
[**FetchUpgradeBundleFromUrl**](SystemAdministrationLifecycleApi.md#FetchUpgradeBundleFromUrl) | **Post** /upgrade/bundles | Fetch upgrade bundle from given url
[**FinishMigrationFinish**](SystemAdministrationLifecycleApi.md#FinishMigrationFinish) | **Post** /migration/plan?action&#x3D;finish | Mark completion of a migration cycle
[**GetAllPreUpgradeChecksInCsvFormatCsv**](SystemAdministrationLifecycleApi.md#GetAllPreUpgradeChecksInCsvFormatCsv) | **Get** /upgrade/pre-upgrade-checks?format&#x3D;csv | Returns pre-upgrade checks in csv format
[**GetBackupConfig**](SystemAdministrationLifecycleApi.md#GetBackupConfig) | **Get** /cluster/backups/config | Get backup configuration
[**GetBackupHistory**](SystemAdministrationLifecycleApi.md#GetBackupHistory) | **Get** /cluster/backups/history | Get backup history
[**GetBackupOverview**](SystemAdministrationLifecycleApi.md#GetBackupOverview) | **Get** /cluster/backups/overview | Get all backup related information for a site
[**GetBackupStatus**](SystemAdministrationLifecycleApi.md#GetBackupStatus) | **Get** /cluster/backups/status | Get backup status
[**GetCapacityThresholds**](SystemAdministrationLifecycleApi.md#GetCapacityThresholds) | **Get** /capacity/threshold | Returns warning threshold(s) set for NSX Objects.
[**GetCapacityUsage**](SystemAdministrationLifecycleApi.md#GetCapacityUsage) | **Get** /capacity/usage | Returns capacity usage data for NSX objects
[**GetDiscoveredSwitches**](SystemAdministrationLifecycleApi.md#GetDiscoveredSwitches) | **Get** /migration/discovered-switches | Get the list of discovered switches (DVS, VSS)
[**GetFeedbackRequests**](SystemAdministrationLifecycleApi.md#GetFeedbackRequests) | **Get** /migration/feedback-requests | NSX-V feedback details
[**GetFeedbackSummary**](SystemAdministrationLifecycleApi.md#GetFeedbackSummary) | **Get** /migration/feedback-summary | Feedback request summary
[**GetGroupedFeedbackRequests**](SystemAdministrationLifecycleApi.md#GetGroupedFeedbackRequests) | **Get** /migration/grouped-feedback-requests | NSX-V feedback details
[**GetLogicalConstructMigrationStats**](SystemAdministrationLifecycleApi.md#GetLogicalConstructMigrationStats) | **Get** /migration/logical-constructs/stats | Get migration stats for logical constructs phase
[**GetMigratedSwitches**](SystemAdministrationLifecycleApi.md#GetMigratedSwitches) | **Get** /migration/migrated-switches | Get the list of migrated switches (DVS, VSS)
[**GetMigrationNodes**](SystemAdministrationLifecycleApi.md#GetMigrationNodes) | **Get** /migration/nodes | Get list of nodes across all types
[**GetMigrationNodesSummary**](SystemAdministrationLifecycleApi.md#GetMigrationNodesSummary) | **Get** /migration/nodes-summary | Get summary of nodes
[**GetMigrationPlanSettings**](SystemAdministrationLifecycleApi.md#GetMigrationPlanSettings) | **Get** /migration/plan/{component_type}/settings | Get migration plan settings for the component
[**GetMigrationStatusSummary**](SystemAdministrationLifecycleApi.md#GetMigrationStatusSummary) | **Get** /migration/status-summary | Get migration status summary
[**GetMigrationSummary**](SystemAdministrationLifecycleApi.md#GetMigrationSummary) | **Get** /migration/summary | Get migration summary
[**GetMigrationSwitch**](SystemAdministrationLifecycleApi.md#GetMigrationSwitch) | **Get** /migration/switch | Get the switch set as current scope for migration
[**GetMigrationUnit**](SystemAdministrationLifecycleApi.md#GetMigrationUnit) | **Get** /migration/migration-units/{migration-unit-id} | Get a specific migration unit
[**GetMigrationUnitAggregateInfo**](SystemAdministrationLifecycleApi.md#GetMigrationUnitAggregateInfo) | **Get** /migration/migration-units/aggregate-info | Get migration units aggregate-info
[**GetMigrationUnitGroup**](SystemAdministrationLifecycleApi.md#GetMigrationUnitGroup) | **Get** /migration/migration-unit-groups/{group-id} | Return migration unit group information
[**GetMigrationUnitGroupAggregateInfo**](SystemAdministrationLifecycleApi.md#GetMigrationUnitGroupAggregateInfo) | **Get** /migration/migration-unit-groups/aggregate-info | Return aggregate information of all migration unit groups
[**GetMigrationUnitGroupStatus**](SystemAdministrationLifecycleApi.md#GetMigrationUnitGroupStatus) | **Get** /migration/migration-unit-groups/{group-id}/status | Get migration status for group
[**GetMigrationUnitGroups**](SystemAdministrationLifecycleApi.md#GetMigrationUnitGroups) | **Get** /migration/migration-unit-groups | Return information of all migration unit groups
[**GetMigrationUnitGroupsStatus**](SystemAdministrationLifecycleApi.md#GetMigrationUnitGroupsStatus) | **Get** /migration/migration-unit-groups-status | Get migration status for migration unit groups
[**GetMigrationUnits**](SystemAdministrationLifecycleApi.md#GetMigrationUnits) | **Get** /migration/migration-units | Get migration units
[**GetMigrationUnitsStats**](SystemAdministrationLifecycleApi.md#GetMigrationUnitsStats) | **Get** /migration/migration-units-stats | Get migration units stats
[**GetNodes**](SystemAdministrationLifecycleApi.md#GetNodes) | **Get** /upgrade/nodes | Get list of nodes across all types
[**GetNodesSummary**](SystemAdministrationLifecycleApi.md#GetNodesSummary) | **Get** /upgrade/nodes-summary | Get summary of nodes
[**GetNsxvSetupDetails**](SystemAdministrationLifecycleApi.md#GetNsxvSetupDetails) | **Get** /migration/setup | NSX-V setup details
[**GetPreUpgradeCheckFailures**](SystemAdministrationLifecycleApi.md#GetPreUpgradeCheckFailures) | **Get** /upgrade/pre-upgrade-checks/failures | Get Pre-upgrade Check Failures
[**GetRestoreConfig**](SystemAdministrationLifecycleApi.md#GetRestoreConfig) | **Get** /cluster/restore/config | Deprecated. Get Restore configuration
[**GetSshFingerprintOfServerRetrieveSshFingerprint**](SystemAdministrationLifecycleApi.md#GetSshFingerprintOfServerRetrieveSshFingerprint) | **Post** /cluster/backups?action&#x3D;retrieve_ssh_fingerprint | Get ssh fingerprint of remote(backup) server
[**GetUcFunctionalState**](SystemAdministrationLifecycleApi.md#GetUcFunctionalState) | **Get** /upgrade/functional-state | Get functional state of the upgrade coordinator
[**GetUcUpgradeStatus**](SystemAdministrationLifecycleApi.md#GetUcUpgradeStatus) | **Get** /upgrade/uc-upgrade-status | Get upgrade-coordinator upgrade status
[**GetUpgradeBundleInfo**](SystemAdministrationLifecycleApi.md#GetUpgradeBundleInfo) | **Get** /upgrade/bundles/{bundle-id} | Get uploaded upgrade bundle information
[**GetUpgradeBundleUploadStatus**](SystemAdministrationLifecycleApi.md#GetUpgradeBundleUploadStatus) | **Get** /upgrade/bundles/{bundle-id}/upload-status | Get uploaded upgrade bundle upload status
[**GetUpgradeChecksInfo**](SystemAdministrationLifecycleApi.md#GetUpgradeChecksInfo) | **Get** /upgrade/upgrade-checks-info | Returns information about upgrade checks
[**GetUpgradeEULAAcceptance**](SystemAdministrationLifecycleApi.md#GetUpgradeEULAAcceptance) | **Get** /upgrade/eula/acceptance | Return the acceptance status of end user license agreement 
[**GetUpgradeEULAContent**](SystemAdministrationLifecycleApi.md#GetUpgradeEULAContent) | **Get** /upgrade/eula/content | Return the content of end user license agreement 
[**GetUpgradeHistory**](SystemAdministrationLifecycleApi.md#GetUpgradeHistory) | **Get** /upgrade/history | Get upgrade history
[**GetUpgradePlanSettings**](SystemAdministrationLifecycleApi.md#GetUpgradePlanSettings) | **Get** /upgrade/plan/{component_type}/settings | Get upgrade plan settings for the component
[**GetUpgradeProgressStatus**](SystemAdministrationLifecycleApi.md#GetUpgradeProgressStatus) | **Get** /node/upgrade/progress-status | Get upgrade progress status
[**GetUpgradeStatusSummary**](SystemAdministrationLifecycleApi.md#GetUpgradeStatusSummary) | **Get** /upgrade/status-summary | Get upgrade status summary
[**GetUpgradeSummary**](SystemAdministrationLifecycleApi.md#GetUpgradeSummary) | **Get** /upgrade/summary | Get upgrade summary
[**GetUpgradeTaskStatus**](SystemAdministrationLifecycleApi.md#GetUpgradeTaskStatus) | **Get** /node/upgrade | Get upgrade task status
[**GetUpgradeUnit**](SystemAdministrationLifecycleApi.md#GetUpgradeUnit) | **Get** /upgrade/upgrade-units/{upgrade-unit-id} | Get a specific upgrade unit
[**GetUpgradeUnitAggregateInfo**](SystemAdministrationLifecycleApi.md#GetUpgradeUnitAggregateInfo) | **Get** /upgrade/upgrade-units/aggregate-info | Get upgrade units aggregate-info
[**GetUpgradeUnitGroup**](SystemAdministrationLifecycleApi.md#GetUpgradeUnitGroup) | **Get** /upgrade/upgrade-unit-groups/{group-id} | Return upgrade unit group information
[**GetUpgradeUnitGroupAggregateInfo**](SystemAdministrationLifecycleApi.md#GetUpgradeUnitGroupAggregateInfo) | **Get** /upgrade/upgrade-unit-groups/aggregate-info | Return aggregate information of all upgrade unit groups
[**GetUpgradeUnitGroupStatus**](SystemAdministrationLifecycleApi.md#GetUpgradeUnitGroupStatus) | **Get** /upgrade/upgrade-unit-groups/{group-id}/status | Get upgrade status for group
[**GetUpgradeUnitGroups**](SystemAdministrationLifecycleApi.md#GetUpgradeUnitGroups) | **Get** /upgrade/upgrade-unit-groups | Return information of all upgrade unit groups
[**GetUpgradeUnitGroupsStatus**](SystemAdministrationLifecycleApi.md#GetUpgradeUnitGroupsStatus) | **Get** /upgrade/upgrade-unit-groups-status | Get upgrade status for upgrade unit groups
[**GetUpgradeUnits**](SystemAdministrationLifecycleApi.md#GetUpgradeUnits) | **Get** /upgrade/upgrade-units | Get upgrade units
[**GetUpgradeUnitsStats**](SystemAdministrationLifecycleApi.md#GetUpgradeUnitsStats) | **Get** /upgrade/upgrade-units-stats | Get upgrade units stats
[**GetVersionWhitelist**](SystemAdministrationLifecycleApi.md#GetVersionWhitelist) | **Get** /upgrade/version-whitelist | Get the version whitelist
[**GetVersionWhitelistByComponent**](SystemAdministrationLifecycleApi.md#GetVersionWhitelistByComponent) | **Get** /upgrade/version-whitelist/{component_type} | Get the version whitelist for the specified component
[**GetView**](SystemAdministrationLifecycleApi.md#GetView) | **Get** /ui-views/{view-id} | Returns View Information
[**GetWidgetConfiguration**](SystemAdministrationLifecycleApi.md#GetWidgetConfiguration) | **Get** /ui-views/{view-id}/widgetconfigurations/{widgetconfiguration-id} | Returns Widget Configuration Information
[**InitiateClusterRestoreStart**](SystemAdministrationLifecycleApi.md#InitiateClusterRestoreStart) | **Post** /cluster/restore?action&#x3D;start | Initiate a restore operation
[**ListClusterBackupTimestamps**](SystemAdministrationLifecycleApi.md#ListClusterBackupTimestamps) | **Get** /cluster/restore/backuptimestamps | List timestamps of all available Cluster Backups.
[**ListFiles**](SystemAdministrationLifecycleApi.md#ListFiles) | **Get** /node/file-store | List node files
[**ListRestoreInstructionResources**](SystemAdministrationLifecycleApi.md#ListRestoreInstructionResources) | **Get** /cluster/restore/instruction-resources | List resources for a given instruction, to be shown to/executed by users. 
[**ListViews**](SystemAdministrationLifecycleApi.md#ListViews) | **Get** /ui-views | Returns the Views based on query criteria defined in ViewQueryParameters.
[**ListWidgetConfigurations**](SystemAdministrationLifecycleApi.md#ListWidgetConfigurations) | **Get** /ui-views/{view-id}/widgetconfigurations | Returns the Widget Configurations based on query criteria defined in WidgetQueryParameters.
[**PauseMigrationPause**](SystemAdministrationLifecycleApi.md#PauseMigrationPause) | **Post** /migration/plan?action&#x3D;pause | Pause migration
[**PauseUpgradePause**](SystemAdministrationLifecycleApi.md#PauseUpgradePause) | **Post** /upgrade/plan?action&#x3D;pause | Pause upgrade
[**QueryClusterRestoreStatus**](SystemAdministrationLifecycleApi.md#QueryClusterRestoreStatus) | **Get** /cluster/restore/status | Query Restore Request Status
[**ReadFile**](SystemAdministrationLifecycleApi.md#ReadFile) | **Get** /node/file-store/{file-name}/data | Read file contents
[**ReadFileProperties**](SystemAdministrationLifecycleApi.md#ReadFileProperties) | **Get** /node/file-store/{file-name} | Read file properties
[**ReadFileThumbprint**](SystemAdministrationLifecycleApi.md#ReadFileThumbprint) | **Get** /node/file-store/{file-name}/thumbprint | Read file thumbprint
[**ReorderMigrationUnitGroupReorder**](SystemAdministrationLifecycleApi.md#ReorderMigrationUnitGroupReorder) | **Post** /migration/migration-unit-groups/{group-id}?action&#x3D;reorder | Reorder migration unit group
[**ReorderMigrationUnitReorder**](SystemAdministrationLifecycleApi.md#ReorderMigrationUnitReorder) | **Post** /migration/migration-unit-groups/{group-id}/migration-unit/{migration-unit-id}?action&#x3D;reorder | Reorder an migration unit within the migration unit group
[**ReorderUpgradeUnitGroupReorder**](SystemAdministrationLifecycleApi.md#ReorderUpgradeUnitGroupReorder) | **Post** /upgrade/upgrade-unit-groups/{group-id}?action&#x3D;reorder | Reorder upgrade unit group
[**ReorderUpgradeUnitReorder**](SystemAdministrationLifecycleApi.md#ReorderUpgradeUnitReorder) | **Post** /upgrade/upgrade-unit-groups/{group-id}/upgrade-unit/{upgrade-unit-id}?action&#x3D;reorder | Reorder an upgrade unit within the upgrade unit group
[**RequestOnetimeBackupBackupToRemote**](SystemAdministrationLifecycleApi.md#RequestOnetimeBackupBackupToRemote) | **Post** /cluster?action&#x3D;backup_to_remote | Request one-time backup
[**RequestOnetimeInventorySummarySummarizeInventoryToRemote**](SystemAdministrationLifecycleApi.md#RequestOnetimeInventorySummarySummarizeInventoryToRemote) | **Post** /cluster?action&#x3D;summarize_inventory_to_remote | Request one-time inventory summary.
[**ResetMigrationPlanReset**](SystemAdministrationLifecycleApi.md#ResetMigrationPlanReset) | **Post** /migration/plan?action&#x3D;reset | Reset migration plan to default plan
[**ResetUpgradePlanReset**](SystemAdministrationLifecycleApi.md#ResetUpgradePlanReset) | **Post** /upgrade/plan?action&#x3D;reset | Reset upgrade plan to default plan
[**RetryClusterRestoreRetry**](SystemAdministrationLifecycleApi.md#RetryClusterRestoreRetry) | **Post** /cluster/restore?action&#x3D;retry | Retry any failed restore operation
[**SetMigrationSwitch**](SystemAdministrationLifecycleApi.md#SetMigrationSwitch) | **Put** /migration/switch | Set the switch as current scope for migration
[**StartMigrationStart**](SystemAdministrationLifecycleApi.md#StartMigrationStart) | **Post** /migration/plan?action&#x3D;start | Start migration
[**StartRollbackMigrationRollback**](SystemAdministrationLifecycleApi.md#StartRollbackMigrationRollback) | **Post** /migration/plan?action&#x3D;rollback | Rollbabck migration
[**StartUpgradeStart**](SystemAdministrationLifecycleApi.md#StartUpgradeStart) | **Post** /upgrade/plan?action&#x3D;start | Start upgrade
[**SuspendClusterRestoreSuspend**](SystemAdministrationLifecycleApi.md#SuspendClusterRestoreSuspend) | **Post** /cluster/restore?action&#x3D;suspend | Suspend any running restore operation
[**TriggerUcUpgradeUpgradeUc**](SystemAdministrationLifecycleApi.md#TriggerUcUpgradeUpgradeUc) | **Post** /upgrade?action&#x3D;upgrade_uc | Upgrade the upgrade coordinator.
[**UpdateCapacityThresholds**](SystemAdministrationLifecycleApi.md#UpdateCapacityThresholds) | **Put** /capacity/threshold | Updates the warning threshold(s) for NSX Objects.
[**UpdateFeedbackResponse**](SystemAdministrationLifecycleApi.md#UpdateFeedbackResponse) | **Put** /migration/feedback-response | Migration feedback response
[**UpdateFile**](SystemAdministrationLifecycleApi.md#UpdateFile) | **Put** /node/file-store/{file-name}/data | Replace file contents
[**UpdateMigrationPlanSettings**](SystemAdministrationLifecycleApi.md#UpdateMigrationPlanSettings) | **Put** /migration/plan/{component_type}/settings | Update migration plan settings for the component
[**UpdateMigrationUnitGroup**](SystemAdministrationLifecycleApi.md#UpdateMigrationUnitGroup) | **Put** /migration/migration-unit-groups/{group-id} | Update the migration unit group
[**UpdateNsxvSetupDetails**](SystemAdministrationLifecycleApi.md#UpdateNsxvSetupDetails) | **Put** /migration/setup | NSX-V setup details
[**UpdateUpgradePlanSettings**](SystemAdministrationLifecycleApi.md#UpdateUpgradePlanSettings) | **Put** /upgrade/plan/{component_type}/settings | Update upgrade plan settings for the component
[**UpdateUpgradeUnitGroup**](SystemAdministrationLifecycleApi.md#UpdateUpgradeUnitGroup) | **Put** /upgrade/upgrade-unit-groups/{group-id} | Update the upgrade unit group
[**UpdateVersionWhitelist**](SystemAdministrationLifecycleApi.md#UpdateVersionWhitelist) | **Put** /upgrade/version-whitelist/{component_type} | Update the version whitelist for the specified component type
[**UpdateView**](SystemAdministrationLifecycleApi.md#UpdateView) | **Put** /ui-views/{view-id} | Update View
[**UpdateWidgetConfiguration**](SystemAdministrationLifecycleApi.md#UpdateWidgetConfiguration) | **Put** /ui-views/{view-id}/widgetconfigurations/{widgetconfiguration-id} | Update Widget Configuration
[**UpgradeSelectedUnitsUpgradeSelectedUnits**](SystemAdministrationLifecycleApi.md#UpgradeSelectedUnitsUpgradeSelectedUnits) | **Post** /upgrade/plan?action&#x3D;upgrade_selected_units | Upgrade selected units
[**UploadUpgradeBundleAsyncUpload**](SystemAdministrationLifecycleApi.md#UploadUpgradeBundleAsyncUpload) | **Post** /upgrade/bundles?action&#x3D;upload | Upload upgrade bundle

# **AbortMigrationAbort**
> AbortMigrationAbort(ctx, )
Abort migration

Resets all migration steps done so far, so that migration can be restarted with new setup details. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AbortPreUpgradeChecksAbortPreUpgradeChecks**
> AbortPreUpgradeChecksAbortPreUpgradeChecks(ctx, )
Abort pre-upgrade checks

Aborts execution of pre-upgrade checks if already in progress. Halts the execution of checks awaiting execution at this point and makes best-effort attempts to stop checks already in execution. Returns without action if execution of pre-upgrade checks is not in progress. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AcceptRecommendedValueInFeedbackAcceptRecommended**
> AcceptRecommendedValueInFeedbackAcceptRecommended(ctx, )
Accept default action for feedbacks

Pick default resolution for all feedback items. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AcceptUpgradeEULA**
> AcceptUpgradeEULA(ctx, )
Accept end user license agreement 

Accept end user license agreement 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddMigrationUnitsToGroupAddMigrationUnits**
> MigrationUnitList AddMigrationUnitsToGroupAddMigrationUnits(ctx, body, groupId)
Add migration units to specified migration unit group

Add migration units to specified migration unit group. The migration units will be added at the end of the migration unit list. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MigrationUnitList**](MigrationUnitList.md)|  | 
  **groupId** | **string**|  | 

### Return type

[**MigrationUnitList**](MigrationUnitList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddUpgradeUnitsToGroupAddUpgradeUnits**
> UpgradeUnitList AddUpgradeUnitsToGroupAddUpgradeUnits(ctx, body, groupId)
Add upgrade units to specified upgrade unit group

Add upgrade units to specified upgrade unit group. The upgrade units will be added at the end of the upgrade unit list. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpgradeUnitList**](UpgradeUnitList.md)|  | 
  **groupId** | **string**|  | 

### Return type

[**UpgradeUnitList**](UpgradeUnitList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdvanceClusterRestoreAdvance**
> ClusterRestoreStatus AdvanceClusterRestoreAdvance(ctx, body)
Advance any suspended restore operation

Advance any currently suspended restore operation. The operation might have been suspended because (1) the user had suspended it previously, or (2) the operation is waiting for user input, to be provided as a part of the POST request body. This operation is only valid when a GET cluster/restore/status returns a status with value SUSPENDED. Otherwise, a 409 response is returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AdvanceClusterRestoreRequest**](AdvanceClusterRestoreRequest.md)|  | 

### Return type

[**ClusterRestoreStatus**](ClusterRestoreStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelClusterRestoreCancel**
> ClusterRestoreStatus CancelClusterRestoreCancel(ctx, )
Cancel any running restore operation

This operation is only valid when a restore is in suspended state. The UI user can cancel any restore operation when the restore is suspended either due to an error, or for a user input. The API user would need to monitor the progression of a restore by calling periodically \"/api/v1/cluster/restore/status\" API. The response object (ClusterRestoreStatus), contains a field \"endpoints\". The API user can cancel the restore process if 'cancel' action is shown in the endpoint field. This operation is only valid when a GET cluster/restore/status returns a status with value SUSPENDED. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterRestoreStatus**](ClusterRestoreStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelUpgradeBundleUploadCancelUpload**
> CancelUpgradeBundleUploadCancelUpload(ctx, bundleId)
Cancel upgrade bundle upload

Cancel upload of upgrade bundle. This API works only when bundle upload is in-progress and will not work during post-processing of upgrade bundle. If bundle upload is in-progress, then the API call returns http OK response after cancelling the upload and deleting partially uploaded bundle. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bundleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigureBackupConfig**
> BackupConfiguration ConfigureBackupConfig(ctx, body, optional)
Configure backup

Configure file server and timers for automated backup. If secret fields are omitted (password, passphrase) then use the previously set value. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BackupConfiguration**](BackupConfiguration.md)|  | 
 **optional** | ***SystemAdministrationLifecycleApiConfigureBackupConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiConfigureBackupConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **frameType** | **optional.**| Frame type | [default to LOCAL_LOCAL_MANAGER]
 **siteId** | **optional.**| Site ID | [default to localhost]

### Return type

[**BackupConfiguration**](BackupConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigureRestoreConfig**
> RestoreConfiguration ConfigureRestoreConfig(ctx, body)
Deprecated. Configure Restore SFTP server credentials

Deprecated. Please use API /cluster/backups/config, to configure remote file server(where backed-up files are stored) details during restore. In older versions - Configure file server where the backed-up files used for the Restore operation are available. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RestoreConfiguration**](RestoreConfiguration.md)|  | 

### Return type

[**RestoreConfiguration**](RestoreConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContinueMigrationContinue**
> ContinueMigrationContinue(ctx, optional)
Continue migration

Continue the migration. Resumes the migration from the point where it was paused. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiContinueMigrationContinueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiContinueMigrationContinueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **optional.Bool**| Skip to migration of next component. | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContinueUpgradeContinue**
> ContinueUpgradeContinue(ctx, optional)
Continue upgrade

Continue the upgrade. Resumes the upgrade from the point where it was paused. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiContinueUpgradeContinueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiContinueUpgradeContinueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentType** | **optional.String**| Component to upgrade. | 
 **skip** | **optional.Bool**| Skip to upgrade of next component. | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CopyFromRemoteFileCopyFromRemoteFile**
> FileProperties CopyFromRemoteFileCopyFromRemoteFile(ctx, body, fileName)
Copy a remote file to the file store

Copy a remote file to the file store. If you use scp or sftp, you must provide the remote server's SSH fingerprint. See the <i>NSX-T Administration Guide</i> for information and instructions about finding the SSH fingerprint. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CopyFromRemoteFileProperties**](CopyFromRemoteFileProperties.md)|  | 
  **fileName** | **string**| Destination filename | 

### Return type

[**FileProperties**](FileProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CopyToRemoteFileCopyToRemoteFile**
> CopyToRemoteFileCopyToRemoteFile(ctx, body, fileName)
Copy file in the file store to a remote file store

Copy a file in the file store to a remote server. If you use scp or sftp, you must provide the remote server's SSH fingerprint. See the <i>NSX-T Administration Guide</i> for information and instructions about finding the SSH fingerprint. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CopyToRemoteFileProperties**](CopyToRemoteFileProperties.md)|  | 
  **fileName** | **string**| Destination filename | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFile**
> FileProperties CreateFile(ctx, fileName)
Upload a file to the file store

When you issue this API, the client must specify: - HTTP header Content-Type:application/octet-stream. - Request body with the contents of the file in the filestore. In the CLI, you can view the filestore with the <em>get files</em> command. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileName** | **string**| Destination filename | 

### Return type

[**FileProperties**](FileProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMigrationUnitGroup**
> MigrationUnitGroup CreateMigrationUnitGroup(ctx, body)
Create a group

Create a group of migration units. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MigrationUnitGroup**](MigrationUnitGroup.md)|  | 

### Return type

[**MigrationUnitGroup**](MigrationUnitGroup.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRemoteDirectoryCreateRemoteDirectory**
> CreateRemoteDirectoryCreateRemoteDirectory(ctx, body)
Create directory in remote file server

Create a directory on the remote remote server. Supports only SFTP. You must provide the remote server's SSH fingerprint. See the <i>NSX Administration Guide</i> for information and instructions about finding the SSH fingerprint. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateRemoteDirectoryProperties**](CreateRemoteDirectoryProperties.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUpgradeUnitGroup**
> UpgradeUnitGroup CreateUpgradeUnitGroup(ctx, body)
Create a group

Create a group of upgrade units. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpgradeUnitGroup**](UpgradeUnitGroup.md)|  | 

### Return type

[**UpgradeUnitGroup**](UpgradeUnitGroup.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateView**
> View CreateView(ctx, body)
Creates a new View.

Creates a new View.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**View**](View.md)|  | 

### Return type

[**View**](View.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateWidgetConfiguration**
> WidgetConfiguration CreateWidgetConfiguration(ctx, body, viewId)
Creates a new Widget Configuration.

Creates a new Widget Configuration and adds it to the specified view. Supported resource_types are LabelValueConfiguration, DonutConfiguration, GridConfiguration, StatsConfiguration, MultiWidgetConfiguration, GraphConfiguration and ContainerConfiguration.  Note: Expressions should be given in a single line. If an expression spans   multiple lines, then form the expression in a single line. For label-value pairs, expressions are evaluated as follows:   a. First, render configurations are evaluated in their order of      appearance in the widget config. The 'field' is evaluated at the end.   b. Second, when render configuration is provided then the order of      evaluation is      1. If expressions provided in 'condition' and 'display value' are         well-formed and free of runtime-errors such as 'null pointers' and         evaluates to 'true'; Then remaining render configurations are not         evaluated, and the current render configuration's 'display value'         is taken as the final value.      2. If expression provided in 'condition' of render configuration is         false, then next render configuration is evaluated.      3. Finally, 'field' is evaluated only when every render configuration         evaluates to false and no error occurs during steps 1 and 2 above.  If an error occurs during evaluation of render configuration, then an   error message is shown. The display value corresponding to that label is   not shown and evaluation of the remaining render configurations continues   to collect and show all the error messages (marked with the 'Label' for   identification) as 'Error_Messages: {}'.  If during evaluation of expressions for any label-value pair an error   occurs, then it is marked with error. The errors are shown in the report,   along with the label value pairs that are error-free.  Important: For elements that take expressions, strings should be provided   by escaping them with a back-slash. These elements are - condition, field,   tooltip text and render_configuration's display_value. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WidgetConfiguration**](WidgetConfiguration.md)|  | 
  **viewId** | **string**|  | 

### Return type

[**WidgetConfiguration**](WidgetConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletView**
> DeletView(ctx, viewId)
Delete View

Delete View

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **viewId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFile**
> DeleteFile(ctx, fileName)
Delete file

Delete file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileName** | **string**| Name of the file to delete | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMigrationUnitGroup**
> DeleteMigrationUnitGroup(ctx, groupId)
Delete the migration unit group

Delete the specified group. NOTE - A group can be deleted only if it is empty. If user tries to delete a group containing one or more migration units, the operation will fail and an error will be returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUpgradeUnitGroup**
> DeleteUpgradeUnitGroup(ctx, groupId)
Delete the upgrade unit group

Delete the specified group. NOTE - A group can be deleted only if it is empty. If user tries to delete a group containing one or more upgrade units, the operation will fail and an error will be returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWidgetConfiguration**
> DeleteWidgetConfiguration(ctx, viewId, widgetconfigurationId)
Delete Widget Configuration

Detaches widget from a given view. If the widget is no longer part of any view, then it will be purged. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **viewId** | **string**|  | 
  **widgetconfigurationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecutePostUpgradeChecksExecutePostUpgradeChecks**
> ExecutePostUpgradeChecksExecutePostUpgradeChecks(ctx, componentType)
Execute post-upgrade checks

Run pre-defined checks to identify issues after upgrade of a component. The results of the checks are added to the respective upgrade units aggregate-info. The progress and status of post-upgrade checks is part of aggregate-info of individual upgrade unit groups. Returns HTTP status 500 with error code 30953 if execution of post-upgrade checks is already in progress. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **componentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecutePreUpgradeChecksExecutePreUpgradeChecks**
> ExecutePreUpgradeChecksExecutePreUpgradeChecks(ctx, optional)
Execute pre-upgrade checks

Run pre-defined checks to identify potential issues which can be encountered during an upgrade or can cause an upgrade to fail. The results of the checks are added to the respective upgrade units aggregate-info. The progress and status of operation is part of upgrade status summary of individual components. Returns HTTP status 500 with error code 30953 if execution of pre-upgrade checks is already in progress. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiExecutePreUpgradeChecksExecutePreUpgradeChecksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiExecutePreUpgradeChecksExecutePreUpgradeChecksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentType** | **optional.String**| Component type on which the action is performed or on which the results are filtered | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecuteUpgradeTask_**
> UpgradeTaskProperties ExecuteUpgradeTask_(ctx, body)
Execute upgrade task

Execute upgrade task. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpgradeTaskProperties**](UpgradeTaskProperties.md)|  | 

### Return type

[**UpgradeTaskProperties**](UpgradeTaskProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchUpgradeBundleFromUrl**
> UpgradeBundleId FetchUpgradeBundleFromUrl(ctx, body)
Fetch upgrade bundle from given url

Fetches the upgrade bundle from url. The call returns after fetch is initiated. Check status by periodically retrieving upgrade bundle upload status using GET /upgrade/bundles/<bundle-id>/upload-status. The upload is complete when the status is SUCCESS. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpgradeBundleFetchRequest**](UpgradeBundleFetchRequest.md)|  | 

### Return type

[**UpgradeBundleId**](UpgradeBundleId.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FinishMigrationFinish**
> FinishMigrationFinish(ctx, )
Mark completion of a migration cycle

This API marks the completion of one execution of migration workflow. This API resets internal  execution state and hence needs to be invoked before starting subsequent workflow run. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllPreUpgradeChecksInCsvFormatCsv**
> UpgradeCheckCsvListResult GetAllPreUpgradeChecksInCsvFormatCsv(ctx, )
Returns pre-upgrade checks in csv format

Returns pre-upgrade checks in csv format 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UpgradeCheckCsvListResult**](UpgradeCheckCsvListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupConfig**
> BackupConfiguration GetBackupConfig(ctx, )
Get backup configuration

Get a configuration of a file server and timers for automated backup. Fields that contain secrets (password, passphrase) are not returned. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**BackupConfiguration**](BackupConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupHistory**
> BackupOperationHistory GetBackupHistory(ctx, )
Get backup history

Get history of previous backup operations 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**BackupOperationHistory**](BackupOperationHistory.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupOverview**
> BackupOverview GetBackupOverview(ctx, optional)
Get all backup related information for a site

Get a configuration of a file server, timers for automated backup, latest backup status, backups list for a site. Fields that contain secrets (password, passphrase) are not returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetBackupOverviewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetBackupOverviewOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **frameType** | **optional.String**| Frame type | [default to LOCAL_LOCAL_MANAGER]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **showBackupsList** | **optional.Bool**| Need a list of backups | [default to true]
 **siteId** | **optional.String**| Site ID | [default to localhost]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**BackupOverview**](BackupOverview.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupStatus**
> CurrentBackupOperationStatus GetBackupStatus(ctx, )
Get backup status

Get status of active backup operations 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CurrentBackupOperationStatus**](CurrentBackupOperationStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCapacityThresholds**
> CapacityThresholdList GetCapacityThresholds(ctx, )
Returns warning threshold(s) set for NSX Objects.

Returns warning threshold(s) set for NSX Objects.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CapacityThresholdList**](CapacityThresholdList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCapacityUsage**
> CapacityUsageResponse GetCapacityUsage(ctx, optional)
Returns capacity usage data for NSX objects

Returns capacity usage data for NSX objects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetCapacityUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetCapacityUsageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **optional.String**|  | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **force** | **optional.Bool**|  | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**CapacityUsageResponse**](CapacityUsageResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDiscoveredSwitches**
> MigrationSwitchListResult GetDiscoveredSwitches(ctx, optional)
Get the list of discovered switches (DVS, VSS)

Get the list of discovered switches (DVS, VSS) for the selected VC. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetDiscoveredSwitchesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetDiscoveredSwitchesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**MigrationSwitchListResult**](MigrationSwitchListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeedbackRequests**
> MigrationFeedbackRequestListResult GetFeedbackRequests(ctx, optional)
NSX-V feedback details

Get feedback details of NSX-V to be migrated. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetFeedbackRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetFeedbackRequestsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **optional.String**| Category on which feedback request should be filtered | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **hash** | **optional.String**| Hash based on which feedback request should be filtered | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **state** | **optional.String**| Filter based on current state of the feedback request | [default to ALL]
 **subCategory** | **optional.String**| Sub category based on which feedback request should be filtered | 

### Return type

[**MigrationFeedbackRequestListResult**](MigrationFeedbackRequestListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeedbackSummary**
> MigrationFeedbackSummaryListResult GetFeedbackSummary(ctx, optional)
Feedback request summary

Get feedback summary of NSX-V to be migrated. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetFeedbackSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetFeedbackSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**MigrationFeedbackSummaryListResult**](MigrationFeedbackSummaryListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupedFeedbackRequests**
> GroupedMigrationFeedbackRequestListResult GetGroupedFeedbackRequests(ctx, optional)
NSX-V feedback details

Get feedback details of NSX-V to be migrated, grouped by feedback type. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetGroupedFeedbackRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetGroupedFeedbackRequestsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **optional.String**| Category on which feedback request should be filtered | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **hash** | **optional.String**| Hash based on which feedback request should be filtered | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **state** | **optional.String**| Filter based on current state of the feedback request | [default to ALL]
 **subCategory** | **optional.String**| Sub category based on which feedback request should be filtered | 

### Return type

[**GroupedMigrationFeedbackRequestListResult**](GroupedMigrationFeedbackRequestListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalConstructMigrationStats**
> LogicalConstructMigrationStatsListResult GetLogicalConstructMigrationStats(ctx, optional)
Get migration stats for logical constructs phase

Get migration stats for logical constructs phase. This API can be polled for getting runtime progress of the migration from source to target.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetLogicalConstructMigrationStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetLogicalConstructMigrationStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**LogicalConstructMigrationStatsListResult**](LogicalConstructMigrationStatsListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMigratedSwitches**
> MigrationSwitchListResult GetMigratedSwitches(ctx, optional)
Get the list of migrated switches (DVS, VSS)

Get the list of migrated switches (DVS, VSS) for the selected VC. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetMigratedSwitchesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetMigratedSwitchesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**MigrationSwitchListResult**](MigrationSwitchListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMigrationNodes**
> NodeInfoListResult GetMigrationNodes(ctx, optional)
Get list of nodes across all types

Get list of nodes. If request parameter component type is specified, then all nodes for that component will be returned. If request parameter component version is specified, then all nodes at that version will be returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetMigrationNodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetMigrationNodesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentType** | **optional.String**| Component type based on which nodes will be filtered | 
 **componentVersion** | **optional.String**| Component version based on which nodes will be filtered | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**NodeInfoListResult**](NodeInfoListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMigrationNodesSummary**
> NodeSummaryList GetMigrationNodesSummary(ctx, )
Get summary of nodes

Get summary of nodes, which includes node count for each type and component version.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeSummaryList**](NodeSummaryList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMigrationPlanSettings**
> MigrationPlanSettings GetMigrationPlanSettings(ctx, componentType)
Get migration plan settings for the component

Get the migration plan settings for the component. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **componentType** | **string**|  | 

### Return type

[**MigrationPlanSettings**](MigrationPlanSettings.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMigrationStatusSummary**
> MigrationStatus GetMigrationStatusSummary(ctx, optional)
Get migration status summary

Get migration status summary

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetMigrationStatusSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetMigrationStatusSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentType** | **optional.String**| Component type based on which migration units to be filtered | 

### Return type

[**MigrationStatus**](MigrationStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMigrationSummary**
> MigrationSummary GetMigrationSummary(ctx, )
Get migration summary

Get migration summary

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MigrationSummary**](MigrationSummary.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMigrationSwitch**
> MigrationSwitchInfo GetMigrationSwitch(ctx, )
Get the switch set as current scope for migration

The user is returned the switch (DVS/VSS) set as current scope of migration. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MigrationSwitchInfo**](MigrationSwitchInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMigrationUnit**
> MigrationUnit GetMigrationUnit(ctx, migrationUnitId)
Get a specific migration unit

Get a specific migration unit

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationUnitId** | **string**|  | 

### Return type

[**MigrationUnit**](MigrationUnit.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMigrationUnitAggregateInfo**
> MigrationUnitAggregateInfoListResult GetMigrationUnitAggregateInfo(ctx, optional)
Get migration units aggregate-info

Get migration units aggregate-info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetMigrationUnitAggregateInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetMigrationUnitAggregateInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentType** | **optional.String**| Component type based on which migration units to be filtered | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **groupId** | **optional.String**| Identifier of group based on which migration units to be filtered | 
 **hasErrors** | **optional.Bool**| Flag to indicate whether to return only migration units with errors | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **metadata** | **optional.String**| Metadata about migration unit to filter on | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **selectionStatus** | **optional.String**| Flag to indicate whether to return only selected, only deselected or both type of migration units | [default to ALL]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**MigrationUnitAggregateInfoListResult**](MigrationUnitAggregateInfoListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMigrationUnitGroup**
> MigrationUnitGroup GetMigrationUnitGroup(ctx, groupId, optional)
Return migration unit group information

Returns information about a specific migration unit group in the migration plan.  If request parameter summary is set to true, then only count of migration units will be returned, migration units list will be empty. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
 **optional** | ***SystemAdministrationLifecycleApiGetMigrationUnitGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetMigrationUnitGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **summary** | **optional.Bool**| Flag indicating whether to return the summary | [default to false]

### Return type

[**MigrationUnitGroup**](MigrationUnitGroup.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMigrationUnitGroupAggregateInfo**
> MigrationUnitGroupAggregateInfoListResult GetMigrationUnitGroupAggregateInfo(ctx, optional)
Return aggregate information of all migration unit groups

Return information of all migration unit groups in the migration plan.  If request parameter summary is set to true, then only count of migration units will be returned, migration units list will be empty. If request parameter component type is specified, then all migration unit groups for that component will be returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetMigrationUnitGroupAggregateInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetMigrationUnitGroupAggregateInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentType** | **optional.String**| Component type based on which migration unit groups to be filtered | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **summary** | **optional.Bool**| Flag indicating whether to return summary | [default to false]
 **sync** | **optional.Bool**| Synchronize before returning migration unit groups | [default to false]

### Return type

[**MigrationUnitGroupAggregateInfoListResult**](MigrationUnitGroupAggregateInfoListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMigrationUnitGroupStatus**
> MigrationUnitStatusListResult GetMigrationUnitGroupStatus(ctx, groupId, optional)
Get migration status for group

Get migration status for migration units in the specified group. User can specify whether to show only the migration units with errors. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
 **optional** | ***SystemAdministrationLifecycleApiGetMigrationUnitGroupStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetMigrationUnitGroupStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **hasErrors** | **optional.Bool**| Flag to indicate whether to return only migration units with errors | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**MigrationUnitStatusListResult**](MigrationUnitStatusListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMigrationUnitGroups**
> MigrationUnitGroupListResult GetMigrationUnitGroups(ctx, optional)
Return information of all migration unit groups

Return information of all migration unit groups in the migration plan.  If request parameter summary is set to true, then only count of migration units will be returned, migration units list will be empty. If request parameter component type is specified, then all migration unit groups for that component will be returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetMigrationUnitGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetMigrationUnitGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentType** | **optional.String**| Component type based on which migration unit groups to be filtered | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **summary** | **optional.Bool**| Flag indicating whether to return summary | [default to false]
 **sync** | **optional.Bool**| Synchronize before returning migration unit groups | [default to false]

### Return type

[**MigrationUnitGroupListResult**](MigrationUnitGroupListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMigrationUnitGroupsStatus**
> MigrationUnitGroupStatusListResult GetMigrationUnitGroupsStatus(ctx, optional)
Get migration status for migration unit groups

Get migration status for migration unit groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetMigrationUnitGroupsStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetMigrationUnitGroupsStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentType** | **optional.String**| Component type based on which migration unit groups to be filtered | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**MigrationUnitGroupStatusListResult**](MigrationUnitGroupStatusListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMigrationUnits**
> MigrationUnitListResult GetMigrationUnits(ctx, optional)
Get migration units

Get migration units

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetMigrationUnitsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetMigrationUnitsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentType** | **optional.String**| Component type based on which migration units to be filtered | 
 **currentVersion** | **optional.String**| Current version of migration unit based on which migration units to be filtered | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **groupId** | **optional.String**| UUID of group based on which migration units to be filtered | 
 **hasWarnings** | **optional.Bool**| Flag to indicate whether to return only migration units with warnings | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **metadata** | **optional.String**| Metadata about migration unit to filter on | 
 **migrationUnitType** | **optional.String**| Migration unit type based on which migration units to be filtered | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**MigrationUnitListResult**](MigrationUnitListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMigrationUnitsStats**
> MigrationUnitTypeStatsList GetMigrationUnitsStats(ctx, optional)
Get migration units stats

Get migration units stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetMigrationUnitsStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetMigrationUnitsStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **sync** | **optional.Bool**| Synchronize before returning migration unit stats | [default to false]

### Return type

[**MigrationUnitTypeStatsList**](MigrationUnitTypeStatsList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodes**
> NodeInfoListResult GetNodes(ctx, optional)
Get list of nodes across all types

Get list of nodes. If request parameter component type is specified, then all nodes for that component will be returned. If request parameter component version is specified, then all nodes at that version will be returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetNodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetNodesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentType** | **optional.String**| Component type based on which nodes will be filtered | 
 **componentVersion** | **optional.String**| Component version based on which nodes will be filtered | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**NodeInfoListResult**](NodeInfoListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodesSummary**
> NodeSummaryList GetNodesSummary(ctx, )
Get summary of nodes

Get summary of nodes, which includes node count for each type and component version.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeSummaryList**](NodeSummaryList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNsxvSetupDetails**
> MigrationSetupInfo GetNsxvSetupDetails(ctx, )
NSX-V setup details

Get setup details of NSX-V to be migrated. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MigrationSetupInfo**](MigrationSetupInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPreUpgradeCheckFailures**
> UpgradeCheckFailureListResult GetPreUpgradeCheckFailures(ctx, optional)
Get Pre-upgrade Check Failures

Get failures resulting from the last execution of pre-upgrade checks. If the execution of checks is in progress, the response has the list of failures observed so far. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetPreUpgradeCheckFailuresOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetPreUpgradeCheckFailuresOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentType** | **optional.String**| Component type | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **filterText** | **optional.String**| Filter text | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **type_** | **optional.String**| Status of the upgrade check | 

### Return type

[**UpgradeCheckFailureListResult**](UpgradeCheckFailureListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRestoreConfig**
> RestoreConfiguration GetRestoreConfig(ctx, )
Deprecated. Get Restore configuration

Deprecated. Please use API /cluster/backups/config, to get remote file server(where backuped-up files are stored) details durign restore. In older versions - Get configuration information for the file server used to store backed-up files. Fields that contain secrets (password, passphrase) are not returned. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RestoreConfiguration**](RestoreConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSshFingerprintOfServerRetrieveSshFingerprint**
> RemoteServerFingerprint GetSshFingerprintOfServerRetrieveSshFingerprint(ctx, body)
Get ssh fingerprint of remote(backup) server

Get SHA256 fingerprint of ECDSA key of remote server. The caller should independently verify that the key is trusted. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoteServerFingerprintRequest**](RemoteServerFingerprintRequest.md)|  | 

### Return type

[**RemoteServerFingerprint**](RemoteServerFingerprint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUcFunctionalState**
> UcFunctionalState GetUcFunctionalState(ctx, )
Get functional state of the upgrade coordinator

Get the functional state of the upgrade coordinator. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UcFunctionalState**](UcFunctionalState.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUcUpgradeStatus**
> UcUpgradeStatus GetUcUpgradeStatus(ctx, )
Get upgrade-coordinator upgrade status

Get upgrade-coordinator upgrade status 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UcUpgradeStatus**](UcUpgradeStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUpgradeBundleInfo**
> UpgradeBundleInfo GetUpgradeBundleInfo(ctx, bundleId)
Get uploaded upgrade bundle information

Get uploaded upgrade bundle information 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bundleId** | **string**|  | 

### Return type

[**UpgradeBundleInfo**](UpgradeBundleInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUpgradeBundleUploadStatus**
> UpgradeBundleUploadStatus GetUpgradeBundleUploadStatus(ctx, bundleId)
Get uploaded upgrade bundle upload status

Get uploaded upgrade bundle upload status 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bundleId** | **string**|  | 

### Return type

[**UpgradeBundleUploadStatus**](UpgradeBundleUploadStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUpgradeChecksInfo**
> ComponentUpgradeChecksInfoListResult GetUpgradeChecksInfo(ctx, optional)
Returns information about upgrade checks

Returns information of pre-upgrade and post-upgrade checks. If request parameter component type is specified, then returns information about all pre-upgrade and post-upgrade for the component. Otherwise returns information of checks across all component types. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetUpgradeChecksInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetUpgradeChecksInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentType** | **optional.String**| Component type based on which upgrade checks are to be filtered | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ComponentUpgradeChecksInfoListResult**](ComponentUpgradeChecksInfoListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUpgradeEULAAcceptance**
> EulaAcceptance GetUpgradeEULAAcceptance(ctx, )
Return the acceptance status of end user license agreement 

Return the acceptance status of end user license agreement 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EulaAcceptance**](EULAAcceptance.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUpgradeEULAContent**
> EulaContent GetUpgradeEULAContent(ctx, optional)
Return the content of end user license agreement 

Return the content of end user license agreement in the specified format. By default, it's pure string without line break 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetUpgradeEULAContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetUpgradeEULAContentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **valueFormat** | **optional.String**| End User License Agreement content output format | 

### Return type

[**EulaContent**](EULAContent.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUpgradeHistory**
> UpgradeHistoryList GetUpgradeHistory(ctx, )
Get upgrade history

Get upgrade history

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UpgradeHistoryList**](UpgradeHistoryList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUpgradePlanSettings**
> UpgradePlanSettings GetUpgradePlanSettings(ctx, componentType)
Get upgrade plan settings for the component

Get the upgrade plan settings for the component. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **componentType** | **string**|  | 

### Return type

[**UpgradePlanSettings**](UpgradePlanSettings.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUpgradeProgressStatus**
> UpgradeProgressStatus GetUpgradeProgressStatus(ctx, )
Get upgrade progress status

Get progress status of last upgrade step, if upgrade bundle is present. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UpgradeProgressStatus**](UpgradeProgressStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUpgradeStatusSummary**
> UpgradeStatus GetUpgradeStatusSummary(ctx, optional)
Get upgrade status summary

Get upgrade status summary

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetUpgradeStatusSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetUpgradeStatusSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentType** | **optional.String**| Component type based on which upgrade units to be filtered | 
 **selectionStatus** | **optional.String**| Flag to indicate whether to return status for only selected, only deselected or both type of upgrade units | [default to ALL]
 **showHistory** | **optional.Bool**| Get upgrade activity for a given component | 

### Return type

[**UpgradeStatus**](UpgradeStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUpgradeSummary**
> UpgradeSummary GetUpgradeSummary(ctx, )
Get upgrade summary

Get upgrade summary

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UpgradeSummary**](UpgradeSummary.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUpgradeTaskStatus**
> UpgradeTaskProperties GetUpgradeTaskStatus(ctx, optional)
Get upgrade task status

Get upgrade task status for the given task of the given bundle. Both bundle_name and task_id must be provided, otherwise you will receive a 404 NOT FOUND response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetUpgradeTaskStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetUpgradeTaskStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bundleName** | **optional.String**| Bundle Name | 
 **upgradeTaskId** | **optional.String**| Upgrade Task ID | 

### Return type

[**UpgradeTaskProperties**](UpgradeTaskProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUpgradeUnit**
> UpgradeUnit GetUpgradeUnit(ctx, upgradeUnitId)
Get a specific upgrade unit

Get a specific upgrade unit

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **upgradeUnitId** | **string**|  | 

### Return type

[**UpgradeUnit**](UpgradeUnit.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUpgradeUnitAggregateInfo**
> UpgradeUnitAggregateInfoListResult GetUpgradeUnitAggregateInfo(ctx, optional)
Get upgrade units aggregate-info

Get upgrade units aggregate-info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetUpgradeUnitAggregateInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetUpgradeUnitAggregateInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentType** | **optional.String**| Component type based on which upgrade units to be filtered | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **groupId** | **optional.String**| Identifier of group based on which upgrade units to be filtered | 
 **hasErrors** | **optional.Bool**| Flag to indicate whether to return only upgrade units with errors | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **metadata** | **optional.String**| Metadata about upgrade unit to filter on | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **selectionStatus** | **optional.String**| Flag to indicate whether to return only selected, only deselected or both type of upgrade units | [default to ALL]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **upgradeUnitDisplayName** | **optional.String**| Display name of upgrade unit | 

### Return type

[**UpgradeUnitAggregateInfoListResult**](UpgradeUnitAggregateInfoListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUpgradeUnitGroup**
> UpgradeUnitGroup GetUpgradeUnitGroup(ctx, groupId, optional)
Return upgrade unit group information

Returns information about a specific upgrade unit group in the upgrade plan.  If request parameter summary is set to true, then only count of upgrade units will be returned, upgrade units list will be empty. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
 **optional** | ***SystemAdministrationLifecycleApiGetUpgradeUnitGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetUpgradeUnitGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **summary** | **optional.Bool**| Flag indicating whether to return the summary | [default to false]

### Return type

[**UpgradeUnitGroup**](UpgradeUnitGroup.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUpgradeUnitGroupAggregateInfo**
> UpgradeUnitGroupAggregateInfoListResult GetUpgradeUnitGroupAggregateInfo(ctx, optional)
Return aggregate information of all upgrade unit groups

Return information of all upgrade unit groups in the upgrade plan.  If request parameter summary is set to true, then only count of upgrade units will be returned, upgrade units list will be empty. If request parameter component type is specified, then all upgrade unit groups for that component will be returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetUpgradeUnitGroupAggregateInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetUpgradeUnitGroupAggregateInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentType** | **optional.String**| Component type based on which upgrade unit groups to be filtered | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **summary** | **optional.Bool**| Flag indicating whether to return summary | [default to false]
 **sync** | **optional.Bool**| Synchronize before returning upgrade unit groups | [default to false]

### Return type

[**UpgradeUnitGroupAggregateInfoListResult**](UpgradeUnitGroupAggregateInfoListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUpgradeUnitGroupStatus**
> UpgradeUnitStatusListResult GetUpgradeUnitGroupStatus(ctx, groupId, optional)
Get upgrade status for group

Get upgrade status for upgrade units in the specified group. User can specify whether to show only the upgrade units with errors. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
 **optional** | ***SystemAdministrationLifecycleApiGetUpgradeUnitGroupStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetUpgradeUnitGroupStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **hasErrors** | **optional.Bool**| Flag to indicate whether to return only upgrade units with errors | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**UpgradeUnitStatusListResult**](UpgradeUnitStatusListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUpgradeUnitGroups**
> UpgradeUnitGroupListResult GetUpgradeUnitGroups(ctx, optional)
Return information of all upgrade unit groups

Return information of all upgrade unit groups in the upgrade plan.  If request parameter summary is set to true, then only count of upgrade units will be returned, upgrade units list will be empty. If request parameter component type is specified, then all upgrade unit groups for that component will be returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetUpgradeUnitGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetUpgradeUnitGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentType** | **optional.String**| Component type based on which upgrade unit groups to be filtered | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **summary** | **optional.Bool**| Flag indicating whether to return summary | [default to false]
 **sync** | **optional.Bool**| Synchronize before returning upgrade unit groups | [default to false]

### Return type

[**UpgradeUnitGroupListResult**](UpgradeUnitGroupListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUpgradeUnitGroupsStatus**
> UpgradeUnitGroupStatusListResult GetUpgradeUnitGroupsStatus(ctx, optional)
Get upgrade status for upgrade unit groups

Get upgrade status for upgrade unit groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetUpgradeUnitGroupsStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetUpgradeUnitGroupsStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentType** | **optional.String**| Component type on which the action is performed or on which the results are filtered | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**UpgradeUnitGroupStatusListResult**](UpgradeUnitGroupStatusListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUpgradeUnits**
> UpgradeUnitListResult GetUpgradeUnits(ctx, optional)
Get upgrade units

Get upgrade units

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetUpgradeUnitsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetUpgradeUnitsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentType** | **optional.String**| Component type based on which upgrade units to be filtered | 
 **currentVersion** | **optional.String**| Current version of upgrade unit based on which upgrade units to be filtered | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **groupId** | **optional.String**| UUID of group based on which upgrade units to be filtered | 
 **hasWarnings** | **optional.Bool**| Flag to indicate whether to return only upgrade units with warnings | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **metadata** | **optional.String**| Metadata about upgrade unit to filter on | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **upgradeUnitType** | **optional.String**| Upgrade unit type based on which upgrade units to be filtered | 

### Return type

[**UpgradeUnitListResult**](UpgradeUnitListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUpgradeUnitsStats**
> UpgradeUnitTypeStatsList GetUpgradeUnitsStats(ctx, optional)
Get upgrade units stats

Get upgrade units stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiGetUpgradeUnitsStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiGetUpgradeUnitsStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **sync** | **optional.Bool**| Synchronize before returning upgrade unit stats | [default to false]

### Return type

[**UpgradeUnitTypeStatsList**](UpgradeUnitTypeStatsList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVersionWhitelist**
> AcceptableComponentVersionList GetVersionWhitelist(ctx, )
Get the version whitelist

Get whitelist of versions for different components

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AcceptableComponentVersionList**](AcceptableComponentVersionList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVersionWhitelistByComponent**
> AcceptableComponentVersion GetVersionWhitelistByComponent(ctx, componentType)
Get the version whitelist for the specified component

Get whitelist of versions for a component. Component can include HOST, EDGE, CCP, MP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **componentType** | **string**|  | 

### Return type

[**AcceptableComponentVersion**](AcceptableComponentVersion.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetView**
> View GetView(ctx, viewId)
Returns View Information

Returns Information about a specific View. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **viewId** | **string**|  | 

### Return type

[**View**](View.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWidgetConfiguration**
> WidgetConfiguration GetWidgetConfiguration(ctx, viewId, widgetconfigurationId)
Returns Widget Configuration Information

Returns Information about a specific Widget Configuration. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **viewId** | **string**|  | 
  **widgetconfigurationId** | **string**|  | 

### Return type

[**WidgetConfiguration**](WidgetConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InitiateClusterRestoreStart**
> ClusterRestoreStatus InitiateClusterRestoreStart(ctx, body)
Initiate a restore operation

Start the restore of an NSX cluster, from some previously backed-up configuration. This operation is only valid when a GET cluster/restore/status returns a status with value NOT_STARTED. Otherwise, a 409 response is returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InitiateClusterRestoreRequest**](InitiateClusterRestoreRequest.md)|  | 

### Return type

[**ClusterRestoreStatus**](ClusterRestoreStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListClusterBackupTimestamps**
> ClusterBackupInfoListResult ListClusterBackupTimestamps(ctx, optional)
List timestamps of all available Cluster Backups.

Returns timestamps for all backup files that are available on the SFTP server. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiListClusterBackupTimestampsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiListClusterBackupTimestampsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ClusterBackupInfoListResult**](ClusterBackupInfoListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFiles**
> FilePropertiesListResult ListFiles(ctx, )
List node files

List node files

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FilePropertiesListResult**](FilePropertiesListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRestoreInstructionResources**
> ActionableResourceListResult ListRestoreInstructionResources(ctx, instructionId, optional)
List resources for a given instruction, to be shown to/executed by users. 

For restore operations requiring user input e.g. performing an action, accepting/rejecting an action, etc. the information to be conveyed to users is provided in this call. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instructionId** | **string**| Id of the instruction set whose instructions are to be returned | 
 **optional** | ***SystemAdministrationLifecycleApiListRestoreInstructionResourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiListRestoreInstructionResourcesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ActionableResourceListResult**](ActionableResourceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListViews**
> ViewList ListViews(ctx, optional)
Returns the Views based on query criteria defined in ViewQueryParameters.

If no query params are specified then all the views entitled for the user are returned. The views to which a user is entitled to include the views created by the user and the shared views. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiListViewsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiListViewsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | **optional.String**| The tag for which associated views to be queried. | 
 **viewIds** | **optional.String**| Ids of the Views | 
 **widgetId** | **optional.String**| Id of widget configuration | 

### Return type

[**ViewList**](ViewList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWidgetConfigurations**
> WidgetConfigurationList ListWidgetConfigurations(ctx, viewId, optional)
Returns the Widget Configurations based on query criteria defined in WidgetQueryParameters.

If no query params are specified then all the Widget Configurations of the specified view are returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **viewId** | **string**|  | 
 **optional** | ***SystemAdministrationLifecycleApiListWidgetConfigurationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiListWidgetConfigurationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **container** | **optional.String**| Id of the container | 
 **widgetIds** | **optional.String**| Ids of the WidgetConfigurations | 

### Return type

[**WidgetConfigurationList**](WidgetConfigurationList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PauseMigrationPause**
> PauseMigrationPause(ctx, )
Pause migration

Pause the migration. Migration will be paused after migration of all the nodes currently in progress is completed either successfully or with failure. User can make changes in the migration plan when the migration is paused. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PauseUpgradePause**
> PauseUpgradePause(ctx, )
Pause upgrade

Pause the upgrade. Upgrade will be paused after upgrade of all the nodes currently in progress is completed either successfully or with failure. User can make changes in the upgrade plan when the upgrade is paused. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryClusterRestoreStatus**
> ClusterRestoreStatus QueryClusterRestoreStatus(ctx, optional)
Query Restore Request Status

Returns status information for the specified NSX cluster restore request. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiQueryClusterRestoreStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiQueryClusterRestoreStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restoreComponent** | **optional.String**|  | [default to LOCAL_MANAGER]

### Return type

[**ClusterRestoreStatus**](ClusterRestoreStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadFile**
> ReadFile(ctx, fileName)
Read file contents

Read file contents

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileName** | **string**| Name of the file to read | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadFileProperties**
> FileProperties ReadFileProperties(ctx, fileName)
Read file properties

Read file properties

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileName** | **string**| Name of the file to retrieve information about | 

### Return type

[**FileProperties**](FileProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadFileThumbprint**
> FileThumbprint ReadFileThumbprint(ctx, fileName)
Read file thumbprint

Read file thumbprint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileName** | **string**| Name of the file for which thumbprint should be computed | 

### Return type

[**FileThumbprint**](FileThumbprint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReorderMigrationUnitGroupReorder**
> ReorderMigrationUnitGroupReorder(ctx, body, groupId)
Reorder migration unit group

Reorder an migration unit group by placing it before/after the specified migration unit group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ReorderMigrationRequest**](ReorderMigrationRequest.md)|  | 
  **groupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReorderMigrationUnitReorder**
> ReorderMigrationUnitReorder(ctx, body, groupId, migrationUnitId)
Reorder an migration unit within the migration unit group

Reorder an migration unit within the migration unit group by placing it before/after the specified migration unit 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ReorderMigrationRequest**](ReorderMigrationRequest.md)|  | 
  **groupId** | **string**|  | 
  **migrationUnitId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReorderUpgradeUnitGroupReorder**
> ReorderUpgradeUnitGroupReorder(ctx, body, groupId)
Reorder upgrade unit group

Reorder an upgrade unit group by placing it before/after the specified upgrade unit group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ReorderRequest**](ReorderRequest.md)|  | 
  **groupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReorderUpgradeUnitReorder**
> ReorderUpgradeUnitReorder(ctx, body, groupId, upgradeUnitId)
Reorder an upgrade unit within the upgrade unit group

Reorder an upgrade unit within the upgrade unit group by placing it before/after the specified upgrade unit 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ReorderRequest**](ReorderRequest.md)|  | 
  **groupId** | **string**|  | 
  **upgradeUnitId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestOnetimeBackupBackupToRemote**
> RequestOnetimeBackupBackupToRemote(ctx, optional)
Request one-time backup

Request one-time backup. The backup will be uploaded using the same server configuration as for automatic backup. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationLifecycleApiRequestOnetimeBackupBackupToRemoteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationLifecycleApiRequestOnetimeBackupBackupToRemoteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **frameType** | **optional.String**| Frame type | [default to LOCAL_LOCAL_MANAGER]
 **siteId** | **optional.String**| Site ID | [default to localhost]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestOnetimeInventorySummarySummarizeInventoryToRemote**
> RequestOnetimeInventorySummarySummarizeInventoryToRemote(ctx, )
Request one-time inventory summary.

Request one-time inventory summary. The backup will be uploaded using the same server configuration as for an automatic backup. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetMigrationPlanReset**
> ResetMigrationPlanReset(ctx, componentType)
Reset migration plan to default plan

Reset the migration plan to default plan. User has an option to change the default plan. But if after making changes, user wants to go back to the default plan, this is the way to do so. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **componentType** | **string**| Component type | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetUpgradePlanReset**
> ResetUpgradePlanReset(ctx, componentType)
Reset upgrade plan to default plan

Reset the upgrade plan to default plan. User has an option to change the default plan. But if after making changes, user wants to go back to the default plan, this is the way to do so. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **componentType** | **string**| Component type | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetryClusterRestoreRetry**
> ClusterRestoreStatus RetryClusterRestoreRetry(ctx, )
Retry any failed restore operation

Retry any currently in-progress, failed restore operation. Only the last step of the multi-step restore operation would have failed,and only that step is retried. This operation is only valid when a GET cluster/restore/status returns a status with value FAILED. Otherwise, a 409 response is returned. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterRestoreStatus**](ClusterRestoreStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetMigrationSwitch**
> MigrationSwitchInfo SetMigrationSwitch(ctx, body)
Set the switch as current scope for migration

The user specifies a DVS / VSS as the current scope of migration. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MigrationSwitchInfo**](MigrationSwitchInfo.md)|  | 

### Return type

[**MigrationSwitchInfo**](MigrationSwitchInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartMigrationStart**
> StartMigrationStart(ctx, )
Start migration

Start the migration. Migration will start as per the migration plan. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartRollbackMigrationRollback**
> StartRollbackMigrationRollback(ctx, )
Rollbabck migration

Roll back the migration. Changes applied to target NSX will be reverted. Use the migration status API to monitor progress of roll back. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartUpgradeStart**
> StartUpgradeStart(ctx, )
Start upgrade

Start the upgrade. Upgrade will start as per the upgrade plan. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SuspendClusterRestoreSuspend**
> ClusterRestoreStatus SuspendClusterRestoreSuspend(ctx, )
Suspend any running restore operation

Suspend any currently running restore operation. The restore operation is made up of a number of steps. When this call is issued, any currently running step is allowed to finish (successfully or with errors), and the next step (and therefore the entire restore operation) is suspended until a subsequent resume or cancel call is issued. This operation is only valid when a GET cluster/restore/status returns a status with value RUNNING. Otherwise, a 409 response is returned. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterRestoreStatus**](ClusterRestoreStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TriggerUcUpgradeUpgradeUc**
> TriggerUcUpgradeUpgradeUc(ctx, )
Upgrade the upgrade coordinator.

Upgrade the upgrade coordinator module itself. This call is invoked after user uploads an upgrade bundle. Once this call is invoked, upgrade coordinator stops and gets restarted and target version upgrade coordinator module comes up after restart. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCapacityThresholds**
> CapacityThresholdList UpdateCapacityThresholds(ctx, body)
Updates the warning threshold(s) for NSX Objects.

Updates the warning threshold(s) for NSX Objects specified, and returns new threshold(s). Threshold list in the request must contain value for GLOBAL_DEFAULT threshold_type which represents global thresholds. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CapacityThresholdList**](CapacityThresholdList.md)|  | 

### Return type

[**CapacityThresholdList**](CapacityThresholdList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFeedbackResponse**
> UpdateFeedbackResponse(ctx, body)
Migration feedback response

Provide response for feedback queries needed for migration. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MigrationFeedbackResponseList**](MigrationFeedbackResponseList.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFile**
> FileProperties UpdateFile(ctx, fileName)
Replace file contents

Replace file contents

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileName** | **string**| Name of the file to replace | 

### Return type

[**FileProperties**](FileProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMigrationPlanSettings**
> MigrationPlanSettings UpdateMigrationPlanSettings(ctx, body, componentType)
Update migration plan settings for the component

Update the migration plan settings for the component. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MigrationPlanSettings**](MigrationPlanSettings.md)|  | 
  **componentType** | **string**|  | 

### Return type

[**MigrationPlanSettings**](MigrationPlanSettings.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMigrationUnitGroup**
> MigrationUnitGroup UpdateMigrationUnitGroup(ctx, body, groupId)
Update the migration unit group

Update the specified migration unit group. Removal of migration units from the group using this is not allowed. An error will be returned in that case. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MigrationUnitGroup**](MigrationUnitGroup.md)|  | 
  **groupId** | **string**|  | 

### Return type

[**MigrationUnitGroup**](MigrationUnitGroup.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNsxvSetupDetails**
> MigrationSetupInfo UpdateNsxvSetupDetails(ctx, body)
NSX-V setup details

Provide setup details of NSX-V to be migrated. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MigrationSetupInfo**](MigrationSetupInfo.md)|  | 

### Return type

[**MigrationSetupInfo**](MigrationSetupInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUpgradePlanSettings**
> UpgradePlanSettings UpdateUpgradePlanSettings(ctx, body, componentType)
Update upgrade plan settings for the component

Update the upgrade plan settings for the component. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpgradePlanSettings**](UpgradePlanSettings.md)|  | 
  **componentType** | **string**|  | 

### Return type

[**UpgradePlanSettings**](UpgradePlanSettings.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUpgradeUnitGroup**
> UpgradeUnitGroup UpdateUpgradeUnitGroup(ctx, body, groupId)
Update the upgrade unit group

Update the specified upgrade unit group. Removal of upgrade units from the group using this is not allowed. An error will be returned in that case. Following extended_configuration is supported:  Key: upgrade_mode Supported values: maintenance_mode,in_place Default: maintenance_mode  Key: maintenance_mode_config_vsan_mode Supported values: evacuate_all_data, ensure_object_accessibility, no_action Default: ensure_object_accessibility  Key: maintenance_mode_config_evacuate_powered_off_vms Supported values: true, false Default: false  Key: rebootless_upgrade Supported values: true, false Default: true 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpgradeUnitGroup**](UpgradeUnitGroup.md)|  | 
  **groupId** | **string**|  | 

### Return type

[**UpgradeUnitGroup**](UpgradeUnitGroup.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVersionWhitelist**
> UpdateVersionWhitelist(ctx, body, componentType)
Update the version whitelist for the specified component type

Update the version whitelist for the specified component type (HOST, EDGE, CCP, MP).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VersionList**](VersionList.md)|  | 
  **componentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateView**
> View UpdateView(ctx, body, viewId)
Update View

Update View

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**View**](View.md)|  | 
  **viewId** | **string**|  | 

### Return type

[**View**](View.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWidgetConfiguration**
> WidgetConfiguration UpdateWidgetConfiguration(ctx, body, viewId, widgetconfigurationId)
Update Widget Configuration

Updates the widget at the given view. If the widget is referenced by other views, then the widget will be updated in all the views that it is part of. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WidgetConfiguration**](WidgetConfiguration.md)|  | 
  **viewId** | **string**|  | 
  **widgetconfigurationId** | **string**|  | 

### Return type

[**WidgetConfiguration**](WidgetConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpgradeSelectedUnitsUpgradeSelectedUnits**
> UpgradeSelectedUnitsUpgradeSelectedUnits(ctx, body)
Upgrade selected units

Upgrades, Resumes the upgrade of a selected set of units. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpgradeUnitList**](UpgradeUnitList.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadUpgradeBundleAsyncUpload**
> UpgradeBundleId UploadUpgradeBundleAsyncUpload(ctx, file)
Upload upgrade bundle

Upload the upgrade bundle. This call returns after upload is completed. You can check bundle processing status periodically by retrieving upgrade bundle upload status to find out if the upload and processing is completed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **file** | **string****string**|  | 

### Return type

[**UpgradeBundleId**](UpgradeBundleId.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

