package nsxt

// Configuration for taking manual/automated backup
type BackupConfiguration struct {
	RemoteFileServer *RemoteFileServer `json:"remote_file_server"`
	// true if automated backup is enabled
	BackupEnabled bool `json:"backup_enabled,omitempty"`
	// Passphrase used to encrypt backup files. The passphrase specified must be at least 8 characters in length and must contain at least one lowercase, one uppercase, one numeric character and one special character (any other non-space character). 
	Passphrase string `json:"passphrase,omitempty"`
	BackupSchedule *BackupSchedule `json:"backup_schedule,omitempty"`
	// A number of seconds after a last backup, that needs to pass, before a topology change will trigger a generation of a new cluster/node backups. If parameter is not provided, then changes in a topology will not trigger a generation of cluster/node backups.
	AfterInventoryUpdateInterval int64 `json:"after_inventory_update_interval,omitempty"`
	// The minimum number of seconds between each upload of the inventory summary to backup server.
	InventorySummaryInterval int64 `json:"inventory_summary_interval,omitempty"`
}
