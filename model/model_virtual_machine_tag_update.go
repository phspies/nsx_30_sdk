package nsxt

type VirtualMachineTagUpdate struct {
	// External id of the virtual machine to which tags are to be applied
	ExternalId string `json:"external_id"`
	// List of tags to be applied to the virtual machine
	Tags []Tag `json:"tags"`
}
