package nsxt

type AddClusterNodeSpec struct {
	// External identifier of the node
	ExternalId string `json:"external_id,omitempty"`
	ControllerRoleConfig *AddControllerNodeSpec `json:"controller_role_config,omitempty"`
	// Display name for the node
	DisplayName string `json:"display_name,omitempty"`
	MgrRoleConfig *AddManagementNodeSpec `json:"mgr_role_config,omitempty"`
}
