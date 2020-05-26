package nsxt

type MigrationPlanSettings struct {
	// Flag to indicate whether to pause the migration after migration of each group is completed
	PauseAfterEachGroup bool `json:"pause_after_each_group,omitempty"`
	// Flag to indicate whether to pause the migration plan execution when an error occurs
	PauseOnError bool `json:"pause_on_error,omitempty"`
	// Migration Method to specify whether the migration is to be performed serially or in parallel
	Parallel bool `json:"parallel,omitempty"`
}
