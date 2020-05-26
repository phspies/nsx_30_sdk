package nsxt

// This object contains SNMP v3 target/receiver where SNMP traps/notifications will be sent.
type Snmpv3Target struct {
	// Security level indicates whether SNMP communication involves authentication and privacy protocols for this user. Value \"AUTH_PRIV\" indicates both authentication and privacy protocols will be used for SNMP communication.
	SecurityLevel string `json:"security_level,omitempty"`
	// SNMP v3 user id used to notify target server. This SNMP v3 user should already be added in this profile.
	UserId string `json:"user_id"`
	// SNMP v3 target server's port.
	Port int64 `json:"port,omitempty"`
	// SNMP v3 target server's IP or FQDN.
	Server string `json:"server"`
}
