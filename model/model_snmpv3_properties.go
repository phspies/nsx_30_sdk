package nsxt

// This object contains list of SNMP v3 users used to poll NSX nodes over SNMP and list of SNMP v3 targets used to receive SNMP traps/notifications from NSX nodes. Users specified in a SNMP v3 target must exist in the list of SNMP v3 users.
type Snmpv3Properties struct {
	// Authentication protocol used for SNMP v3 communication.
	AuthProtocol string `json:"auth_protocol,omitempty"`
	// Privacy protocol used for SNMP v3 communication.
	PrivProtocol string `json:"priv_protocol,omitempty"`
	// List of SNMP v3 users allowed to poll NSX nodes over SNMP. Also, users specified in a SNMP v3 target must exist in this list.
	Users []Snmpv3User `json:"users,omitempty"`
	// List of SNMP v3 targets/receivers where SNMP v3 traps/notifications will be sent from NSX nodes.
	Targets []Snmpv3Target `json:"targets,omitempty"`
}
