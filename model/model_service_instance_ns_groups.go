package nsxt

// ServiceInstanceNSGroups contains list of NS Groups referenced in North-South Service Insertion Rules for a particular Service Instance.
type ServiceInstanceNsGroups struct {
	// List of NSGroups Used in ServiceInsertion Rules.
	Nsroups []NsGroupInfo `json:"nsroups,omitempty"`
}
