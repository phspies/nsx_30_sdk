package nsxt

// Past backup operation details
type BackupOperationHistory struct {
	// Statuses of previous inventory backups
	InventoryBackupStatuses []BackupOperationStatus `json:"inventory_backup_statuses,omitempty"`
	// Statuses of previous cluser backups
	ClusterBackupStatuses []BackupOperationStatus `json:"cluster_backup_statuses,omitempty"`
	// Statuses of previous node backups
	NodeBackupStatuses []BackupOperationStatus `json:"node_backup_statuses,omitempty"`
}
