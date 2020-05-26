package nsxt

// Deployment progress state of node VM. This Object contains name of current deployment step and overall progress percentage.
type VmDeploymentProgressState struct {
	// Overall progress percentage of deployment completed
	Progress int64 `json:"progress,omitempty"`
	// Name of the current running step of deployment
	CurrentStepTitle string `json:"current_step_title,omitempty"`
}
