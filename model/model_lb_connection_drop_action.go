package nsxt

// This action is used to drop the connections. There is no extra propery in this action. If there is no match condition specified, the connection will be always dropped. This action can be specified at HTTP_ACCESS or HTTP_FORWARDING pahse. 
type LbConnectionDropAction struct {
	// The property identifies the load balancer rule action type. 
	Type_ string `json:"type"`
}
