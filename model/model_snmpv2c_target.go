package nsxt

// This object contains SNMP v2c target/receiver where SNMP traps/notifications will be sent.
type Snmpv2cTarget struct {
	// Unique non-sensitive community name to identify community.
	CommunityName string `json:"community_name"`
	// Community string (shared secret). This field is required when adding a community target. When updating a community target, do not include this field in the request. If this field is present in an update request, it will be considered as a new value for community string.
	CommunityString string `json:"community_string,omitempty"`
	// SNMP v2c target server's port number.
	Port int64 `json:"port,omitempty"`
	// SNMP v2c target server's IP or FQDN.
	Server string `json:"server"`
}
