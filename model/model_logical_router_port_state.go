package nsxt

// This holds the state of Logical Router Port. If there are errors in realizing LRP outside of MP, it gives details of the components and specific errors. 
type LogicalRouterPortState struct {
	// Request identifier of the API which modified the entity.
	PendingChangeList []string `json:"pending_change_list,omitempty"`
	// Array of DAD status which contains DAD information for IP addresses on the port. 
	Ipv6DadStatuses []IPv6DadStatus `json:"ipv6_dad_statuses,omitempty"`
}
