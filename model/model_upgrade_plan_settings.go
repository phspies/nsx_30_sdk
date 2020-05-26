package nsxt

type UpgradePlanSettings struct {
	// Flag to indicate whether to pause the upgrade after upgrade of each group is completed
	PauseAfterEachGroup bool `json:"pause_after_each_group,omitempty"`
	// Flag to indicate whether to pause the upgrade plan execution when an error occurs
	PauseOnError bool `json:"pause_on_error,omitempty"`
	// Upgrade Method to specify whether the upgrade is to be performed serially or in parallel
	Parallel bool `json:"parallel,omitempty"`
}
