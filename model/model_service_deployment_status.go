package nsxt

type ServiceDeploymentStatus struct {
	// List of issue and detailed description of the issue in case of deployment failure.
	DeploymentIssues []ServiceDeploymentIssue `json:"deployment_issues,omitempty"`
	// Timestamp when the data was last updated; unset if data source has never updated the data.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// Deployment status of NXGI Partner Service-VM on a compute collection. It shows the latest status during the process of deployment, redeploy, upgrade, and un-deployment on a compute collection such as VC cluster.
	DeploymentStatus string `json:"deployment_status,omitempty"`
	// Currently deployed Service Virtual Appliance version.
	SvaCurrentVersion string `json:"sva_current_version,omitempty"`
	// Id of service deployment.
	ServiceDeploymentId string `json:"service_deployment_id,omitempty"`
	// Max available SVA version for upgrade
	SvaMaxAvailableVersion string `json:"sva_max_available_version,omitempty"`
}
