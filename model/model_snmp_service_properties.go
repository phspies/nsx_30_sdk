package nsxt

// SNMP Service properties
type SnmpServiceProperties struct {
	// SNMP v3 auth protocol
	V3AuthProtocol string `json:"v3_auth_protocol"`
	// SNMP v1, v2c community strings
	Communities []string `json:"communities,omitempty"`
	// SNMP v3 is configured or not
	V3Configured bool `json:"v3_configured,omitempty"`
	// SNMP v3 private protocol
	V3PrivProtocol string `json:"v3_priv_protocol"`
	// SNMP v3 users information
	V3Users []SnmpV3User `json:"v3_users,omitempty"`
	// SNMP v2 is configured or not
	V2Configured bool `json:"v2_configured,omitempty"`
	// Start when system boots
	StartOnBoot bool `json:"start_on_boot"`
}
