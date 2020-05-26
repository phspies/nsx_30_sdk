package nsxt

type RepoSyncStatusReport struct {
	// Status of the repo sync operation on the single nsx-manager 
	Status string `json:"status"`
	// Describes the steps which repo sync operation is performing currently. 
	StatusMessage string `json:"status_message,omitempty"`
	// In case if repo sync fails due to some issue, an error message will be stored here. 
	FailureMessage string `json:"failure_message,omitempty"`
	// In case of repo sync related failure, the code for the error will be stored here. 
	FailureCode int64 `json:"failure_code,omitempty"`
}
