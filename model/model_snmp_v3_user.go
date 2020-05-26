package nsxt

// SNMP v3 user properties
type SnmpV3User struct {
	// SNMP v3 user private password
	PrivPassword string `json:"priv_password,omitempty"`
	// SNMP v3 user auth password
	AuthPassword string `json:"auth_password,omitempty"`
	// SNMP v3 user ID
	UserId string `json:"user_id"`
}
