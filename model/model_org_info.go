package nsxt

// Organization ID and role, predefined for a particular type of VMware support.
type OrgInfo struct {
	// Organization ID, connected to a predefined role of a VMware support.
	OrgId string `json:"org_id"`
	// Predefined role of a VMware support.
	OrgRole string `json:"org_role"`
}
