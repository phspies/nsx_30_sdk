package nsxt

// List of entities where Service Insertion will not be enforced. Exclusion List can contain NSGroup(s) or LogicalPort(s) or LogicalSwitch(es).
type SiExcludeList struct {
	// Total number of members present in Exclude List.
	MemberCount int64 `json:"member_count,omitempty"`
	// List of members in Exclusion List
	Members []ResourceReference `json:"members"`
}
