package nsxt

type FirewallSectionRuleList struct {
	// Category from policy framework.
	Category string `json:"category,omitempty"`
	// This attribute represents enforcement point of firewall section. For example, firewall section enforced on logical port with attachment type bridge endpoint will have 'BRIDGEENDPOINT' value, firewall section enforced on logical router will have 'LOGICALROUTER' value and rest have 'VIF' value.
	EnforcedOn string `json:"enforced_on,omitempty"`
	// Section is locked/unlocked.
	Locked bool `json:"locked,omitempty"`
	// If TCP strict is enabled on a section and a packet matches rule in it, the following check will be performed. If the packet does not belong to an existing session, the kernel will check to see if the SYN flag of the packet is set. If it is not, then it will drop the packet.
	TcpStrict bool `json:"tcp_strict,omitempty"`
	// ID of the user who last modified the lock for the section.
	LockModifiedBy string `json:"lock_modified_by,omitempty"`
	// Section locked/unlocked time in epoch milliseconds.
	LockModifiedTime int64 `json:"lock_modified_time,omitempty"`
	// Priority of current section with respect to other sections. In case the field is empty, the list section api should be used to get section priority.
	Priority int64 `json:"priority,omitempty"`
	FirewallSchedule *ResourceReference `json:"firewall_schedule,omitempty"`
	// Comments for section lock/unlock.
	Comments string `json:"comments,omitempty"`
	// This flag indicates whether it is an auto-plumbed section that is associated to a LogicalRouter. Auto-plumbed sections are system owned and cannot be updated via the API.
	Autoplumbed bool `json:"autoplumbed,omitempty"`
	// List of firewall rules in the section. Only homogenous rules are supported.
	Rules []FirewallRule `json:"rules"`
}
