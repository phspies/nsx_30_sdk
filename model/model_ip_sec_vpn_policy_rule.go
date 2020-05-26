package nsxt

// For policy-based IPsec VPNs, a security policy specifies as its action the VPN tunnel to be used for transit traffic that meets the policy's match criteria.
type IpSecVpnPolicyRule struct {
	Owner *OwnerResourceLink `json:"_owner,omitempty"`
	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`
	// Unique policy id.
	Id string `json:"id,omitempty"`
	// The type of this resource.
	ResourceType string `json:"resource_type,omitempty"`
	// Description of this resource
	Description string `json:"description,omitempty"`
	// List of local subnets.
	Sources []IpSecVpnPolicySubnet `json:"sources,omitempty"`
	// PROTECT - Protect rules are defined per policy based IPSec VPN session. BYPASS - Bypass rules are defined per IPSec VPN service and affects all policy based IPSec VPN sessions. Bypass rules are prioritized over protect rules. 
	Action string `json:"action,omitempty"`
	// A flag to enable/disable the policy rule.
	Enabled bool `json:"enabled,omitempty"`
	// A flag to enable/disable the logging for the policy rule.
	Logged bool `json:"logged,omitempty"`
	// List of peer subnets.
	Destinations []IpSecVpnPolicySubnet `json:"destinations,omitempty"`
}
