package nsxt

// This object contains SNMP v2c and SNMP v3 properties.
type SnmpProperties struct {
	V2c *Snmpv2cProperties `json:"v2c,omitempty"`
	V3 *Snmpv3Properties `json:"v3,omitempty"`
}
