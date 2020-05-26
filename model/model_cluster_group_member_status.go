package nsxt

// This type contains the attributes and status of a group member.
type ClusterGroupMemberStatus struct {
	// IP of the group member
	MemberIp string `json:"member_ip,omitempty"`
	// FQDN of the group member
	MemberFqdn string `json:"member_fqdn,omitempty"`
	// Status of the group member
	MemberStatus string `json:"member_status,omitempty"`
	// UUID of the group member
	MemberUuid string `json:"member_uuid,omitempty"`
}
