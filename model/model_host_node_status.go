package nsxt

// Host Node Status
type HostNodeStatus struct {
	// Gives details of state of desired configuration. This property is available only if Tranport Node exists for the host. Following are the supported values pending  - Transport Node configuration status is pending in_progress - Transport Node configuration status is in progress success - Transport Node configuration status is successful failed - Transport Node configuration status is failed partial_success - Transport Node configuration status is partial success orphaned - Transport Node configuration status is orphaned unknown - Transport Node configuration status is unknown error - Error occured during Transport Node configuration
	ConfigStatus string `json:"config_status,omitempty"`
	// Unique Id of the host node
	NodeId string `json:"node_id,omitempty"`
	// This specifies the current nsx install status for host node. Following are the supported values INSTALL_IN_PROGRESS - NSX installation is in progress on the host INSTALL_FAILED - NSX installation failed on the host INSTALL_SUCCESSFUL - NSX installation successful on the host UNINSTALL_IN_PROGRESS - NSX uninstallation in progress on the host UNINSTALL_FAILED - NSX uninstallation failed on the host UNINSTALL_SUCCESSFUL - NSX uninstallation successful on the host UNINSTALL_SCHEDULED - NSX uninstallation is scheduled on the host UPGRADE_IN_PROGRESS - NSX upgrade is in progress on the host UPGRADE_FAILED - NSX upgrade failed on the host DEPLOYMENT_QUEUED - Deployment is queued on the DEPLOYMENT_IN_PROGRESS - Deployment is in progress DEPLOYMENT_FAILED - Deployment is failed DEPLOYMENT_SUCCESSFUL - Deployment is successful UNDEPLOYMENT_QUEUED - Undeployment is queued UNDEPLOYMENT_IN_PROGRESS - Undeployment is in progress UNDEPLOYMENT_FAILED - Undeployment failed UNDEPLOYMENT_SUCCESSFUL - Undeployment is successful UPGRADE_QUEUED - Upgrade is queued HOST_DISCONNECTED - Host is disconnected POWERED_OFF - Host is powered off
	DeploymentStatus string `json:"deployment_status,omitempty"`
}
