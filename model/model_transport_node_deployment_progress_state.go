package nsxt

// Deployment progress state of transport node. Object has current deployment step title and progress in percentage.
type TransportNodeDeploymentProgressState struct {
	// Percentage of deployment completed
	Progress int64 `json:"progress,omitempty"`
	// Deployment step title
	CurrentStepTitle string `json:"current_step_title,omitempty"`
}
