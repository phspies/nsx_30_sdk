package nsxt

// VIF attachment state of a logical port
type LogicalPortAttachmentState struct {
	// A logicalPort must be in one of following state. FREE - If there are no active attachers. The LogicalPort may or may not have an attachment ID configured on it. This state is applicable only to LogialPort of static type. ATTACHED - LogicalPort has exactly one active attacher and no further configuration is pending. ATTACHED_PENDING_CONF - LogicalPort has exactly one attacher, however it may not have been configured completely. Additional configuration will be provided by other nsx components. ATTACHED_IN_MOTION - LogicalPort has multiple active attachers. This state represents a scenario where VM is moving from one location (host or storage) to another (e.g. vmotion, vSphere HA) DETACHED - A temporary state after all LogialPort attachers have been detached. This state is applicable only to LogicalPort of ephemeral type and the LogicalPort will soon be deleted. 
	State string `json:"state,omitempty"`
	// VM or vmknic entities that are attached to the LogicalPort
	Attachers []PortAttacher `json:"attachers,omitempty"`
	// VIF ID
	Id string `json:"id,omitempty"`
}
