package nsxt

type UpgradeCheckListResults struct {
	FailedChecks *UpgradeCheckListResult `json:"failed_checks,omitempty"`
	ChecksWithWarnings *UpgradeCheckListResult `json:"checks_with_warnings,omitempty"`
	SuccessfulChecks *UpgradeCheckListResult `json:"successful_checks,omitempty"`
}
