package nsxt

// Node properties
type NodeProperties struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Current time expressed in milliseconds since epoch
	SystemTime int64 `json:"system_time,omitempty"`
	// Node Unique Identifier
	NodeUuid string `json:"node_uuid,omitempty"`
	// Message of the day to display when users login to node using the NSX CLI
	Motd *interface{} `json:"motd,omitempty"`
	// NSX CLI inactivity timeout, set to 0 to configure no timeout
	CliTimeout int64 `json:"cli_timeout,omitempty"`
	// Kernel version
	KernelVersion string `json:"kernel_version,omitempty"`
	// Export restrictions in effect, if any
	ExportType string `json:"export_type,omitempty"`
	// Host name or fully qualified domain name of node
	Hostname string `json:"hostname,omitempty"`
	// Product version
	ProductVersion string `json:"product_version,omitempty"`
	// Node version
	NodeVersion string `json:"node_version,omitempty"`
	// System date time in UTC
	SystemDatetime string `json:"system_datetime,omitempty"`
	// Fully qualified domain name
	FullyQualifiedDomainName string `json:"fully_qualified_domain_name,omitempty"`
	// Timezone
	Timezone string `json:"timezone,omitempty"`
}
