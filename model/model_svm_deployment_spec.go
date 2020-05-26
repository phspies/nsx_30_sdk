package nsxt

// Deployment Specs holds information required to deploy the Service-VMs.i.e. OVF url where the partner Service-VM OVF is hosted. The host type on which the OVF(Open Virtualization Format) can be deployed, Form factor to name a few.
type SvmDeploymentSpec struct {
	// Location of the partner VM OVF to be deployed.
	OvfUrl string `json:"ovf_url"`
	// Deployment Spec name for ease of use, since multiple DeploymentSpec can be specified.
	Name string `json:"name,omitempty"`
	// Minimum host version supported by this ovf. If a host in the deployment cluster is having version less than this, then service deployment will not happen on that host.
	MinHostVersion string `json:"min_host_version,omitempty"`
	// Supported ServiceInsertion Form Factor for the OVF deployment. The default FormFactor is Medium.
	ServiceFormFactor string `json:"service_form_factor,omitempty"`
	// Host Type on which the specified OVF can be deployed.
	HostType string `json:"host_type"`
	// Partner needs to specify the Service VM version which will get deployed.
	SvmVersion string `json:"svm_version,omitempty"`
}
