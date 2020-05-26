package nsxt

// Uplink Teaming Policy with a name that can be referenced by logical switches
type NamedTeamingPolicy struct {
	// Teaming policy
	Policy string `json:"policy"`
	// List of Uplinks used in standby list
	StandbyList []Uplink `json:"standby_list,omitempty"`
	// List of Uplinks used in active list
	ActiveList []Uplink `json:"active_list"`
	// An uplink teaming policy of a given name defined in UplinkHostSwitchProfile. The names of all NamedTeamingPolicies in an UplinkHostSwitchProfile must be different, but a name can be shared by different UplinkHostSwitchProfiles. Different TransportNodes can use different NamedTeamingPolicies having the same name in different UplinkHostSwitchProfiles to realize an uplink teaming policy on a logical switch. An uplink teaming policy on a logical switch can be any policy defined by a user; it does not have to be a single type of FAILOVER or LOADBALANCE. It can be a combination of types, for instance, a user can define a policy with name \"MyHybridTeamingPolicy\" as \"FAILOVER on all ESX TransportNodes and LOADBALANCE on all KVM TransportNodes\". The name is the key of the teaming policy and can not be changed once assigned.
	Name string `json:"name"`
}
