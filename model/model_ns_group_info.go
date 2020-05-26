package nsxt

// NSGroupInfo contains information about a particular NSGroup used in a SI Rule. It also contains information about policy path used to create this NSGroup.
type NsGroupInfo struct {
	// Relative Policy path of a particular NSGroup.
	NsgroupPolicyPath string `json:"nsgroup_policy_path,omitempty"`
	Nsgroup *ResourceReference `json:"nsgroup,omitempty"`
}
