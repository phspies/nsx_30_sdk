package nsxt

// Abstract base type for specification of IPs to be used with host switch virtual tunnel endpoints
type IpAssignmentSpec struct {
	ResourceType string `json:"resource_type"`
}
