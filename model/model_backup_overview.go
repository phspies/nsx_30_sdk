package nsxt

// Data for a single backup/restore card
type BackupOverview struct {
	// List of timestamps of backed-up cluster files
	Results []ClusterBackupInfo `json:"results,omitempty"`
	CurrentBackupOperationStatus *CurrentBackupOperationStatus `json:"current_backup_operation_status"`
	BackupOperationHistory *BackupOperationHistory `json:"backup_operation_history"`
	BackupConfig *BackupConfiguration `json:"backup_config"`
	RestoreStatus *ClusterRestoreStatus `json:"restore_status"`
}
