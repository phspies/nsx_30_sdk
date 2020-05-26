package nsxt

// This object contains list of SNMP v2c communities used to poll NSX nodes over SNMP and list of SNMP v2c targets used to receive SNMP traps/notifications from NSX nodes.
type Snmpv2cProperties struct {
	// List of SNMP v2c communities allowed to poll NSX nodes over SNMP v2c.
	Communities []Snmpv2cCommunity `json:"communities,omitempty"`
	// List of SNMP v2c targets/receivers where SNMP v2c traps/notifications will be sent from NSX nodes.
	Targets []Snmpv2cTarget `json:"targets,omitempty"`
}
