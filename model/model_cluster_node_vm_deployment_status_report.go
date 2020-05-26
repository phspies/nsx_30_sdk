package nsxt

// Contains up-to-date information relating to an auto-deployed VM, including its status and (potentially) an error message. 
type ClusterNodeVmDeploymentStatusReport struct {
	// Status of the addition or deletion of an auto-deployed cluster node VM. 
	Status string `json:"status"`
	DeploymentProgressState *VmDeploymentProgressState `json:"deployment_progress_state,omitempty"`
	// In case of auto-deployment-related failure, an error message will be stored here. 
	FailureMessage string `json:"failure_message,omitempty"`
	// In case of auto-deployment-related failure, the code for the error will be stored here. 
	FailureCode int64 `json:"failure_code,omitempty"`
}
