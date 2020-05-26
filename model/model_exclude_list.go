package nsxt

// List of entities where Distributed Firewall will not be enforced. Exclusion List can contain NSGroup(s) or LogicalPort(s) or LogicalSwitch(es) to exclude Distributed Firewall enforcement.
type ExcludeList struct {
	// Total number of members present in Exclude List.
	MemberCount int64 `json:"member_count,omitempty"`
	// List of members in Exclusion List
	Members []ResourceReference `json:"members"`
}
