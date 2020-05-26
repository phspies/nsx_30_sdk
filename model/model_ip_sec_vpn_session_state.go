package nsxt

// This holds the state of IPSec VPN Session. If there are errors in realizing session outside of MP, it gives details of the components and specific errors. 
type IpSecVpnSessionState struct {
	// Request identifier of the API which modified the entity.
	PendingChangeList []string `json:"pending_change_list,omitempty"`
}
