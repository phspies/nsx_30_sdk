# BackupConfiguration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteFileServer** | [***RemoteFileServer**](RemoteFileServer.md) |  | [default to null]
**BackupEnabled** | **bool** | true if automated backup is enabled | [optional] [default to false]
**Passphrase** | **string** | Passphrase used to encrypt backup files. The passphrase specified must be at least 8 characters in length and must contain at least one lowercase, one uppercase, one numeric character and one special character (any other non-space character).  | [optional] [default to null]
**BackupSchedule** | [***BackupSchedule**](BackupSchedule.md) |  | [optional] [default to null]
**AfterInventoryUpdateInterval** | **int64** | A number of seconds after a last backup, that needs to pass, before a topology change will trigger a generation of a new cluster/node backups. If parameter is not provided, then changes in a topology will not trigger a generation of cluster/node backups. | [optional] [default to null]
**InventorySummaryInterval** | **int64** | The minimum number of seconds between each upload of the inventory summary to backup server. | [optional] [default to 240]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

