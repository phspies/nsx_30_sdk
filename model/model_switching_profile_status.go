package nsxt

type SwitchingProfileStatus struct {
	// Number of logical ports using a switching profile
	NumLogicalPorts int64 `json:"num_logical_ports,omitempty"`
	// Identifier for the switching profile
	SwitchingProfileId string `json:"switching_profile_id,omitempty"`
	// Number of logical switches using a switching profile
	NumLogicalSwitches int64 `json:"num_logical_switches,omitempty"`
}
