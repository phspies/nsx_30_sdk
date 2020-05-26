package nsxt

// Domain synchronization settings
type DirectoryDomainSyncSettings struct {
	// Sync delay after Directory domain has been successfully created. if delay is -1, initial full sync will not be triggered. 
	SyncDelayInSec int32 `json:"sync_delay_in_sec,omitempty"`
	// Directory domain full synchronization schedule using cron expression. For example, cron expression \"0 0 12 ? * SUN *\" means full sync is scheduled every Sunday midnight. If this object is null, it means there is no background cron job running for full sync.
	FullSyncCronExpr string `json:"full_sync_cron_expr,omitempty"`
	// Directory domain delta synchronization interval time between two delta sync in minutes.
	DeltaSyncInterval int64 `json:"delta_sync_interval,omitempty"`
}
