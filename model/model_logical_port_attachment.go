package nsxt

// Logical port attachment
type LogicalPortAttachment struct {
	// Indicates the type of logical port attachment. By default it is Virtual Machine interface (VIF)
	AttachmentType string `json:"attachment_type,omitempty"`
	Context *AttachmentContext `json:"context,omitempty"`
	// Identifier of the interface attached to the logical port
	Id string `json:"id"`
}
