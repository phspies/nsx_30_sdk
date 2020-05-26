package nsxt

// VM or vmknic entity attached to LogicalPort
type PortAttacher struct {
	// TransportNode on which the attacher resides
	Host string `json:"host"`
	// This is a vmknic name if the attacher is vmknic. Otherwise, it is full path of the attached VM's config file 
	Entity string `json:"entity"`
}
