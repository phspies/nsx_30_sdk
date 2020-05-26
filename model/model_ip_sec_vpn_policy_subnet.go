package nsxt

// Used to specify the local/peer subnets in IPSec VPN Policy rule.
type IpSecVpnPolicySubnet struct {
	// Subnet used in policy rule.
	Subnet string `json:"subnet"`
}
