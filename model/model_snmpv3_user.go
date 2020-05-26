package nsxt

// This object contains properties for a SNMP v3 user that can be used to receive SNMP traps/notifications from NSX and/or poll NSX nodes over SNMP.
type Snmpv3User struct {
	// Privacy password used for SNMP v3 communication. This field is required when adding a user. When updating a user, do not include this field in the request. If this field is present in an update request, it will be considered as a new value for privacy password.
	PrivPassword string `json:"priv_password,omitempty"`
	// Access permissions for polling NSX nodes over SNMP v3.
	Access string `json:"access,omitempty"`
	// Authentication password used for SNMP v3 communication. This field is required when adding a user. When updating a user, do not include this field in the request. If this field is present in an update request, it will be considered as a new value for authentication password.
	AuthPassword string `json:"auth_password,omitempty"`
	// Unique SNMP v3 user id.
	UserId string `json:"user_id"`
	// Security level indicates whether SNMP communication involves authentication and privacy protocols for this user. Value \"AUTH_PRIV\" indicates both authentication and privacy protocols will be used for SNMP communication.
	SecurityLevel string `json:"security_level,omitempty"`
}
