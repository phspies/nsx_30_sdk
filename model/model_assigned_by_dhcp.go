package nsxt

// This type can be specified in ip assignment spec of host switch if DHCP based IP assignment is desired for host switch virtual tunnel endpoints.
type AssignedByDhcp struct {
	ResourceType string `json:"resource_type"`
}
