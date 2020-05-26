package nsxt

// ServiceDeployment Spec consists of information required to deploy and configure the partner appliances. viz. Deployment template, deployment spec and NIC metatdata.
type ServiceDeploymentSpec struct {
	// Deployment Specs holds information required to deploy the Service-VMs. i.e. OVF url where the partner Service-VM OVF is hosted. The host type on which the OVF can be deployed, Form factor to name a few.
	DeploymentSpecs []SvmDeploymentSpec `json:"deployment_specs,omitempty"`
	// NIC metadata associated with the deployment spec.
	NicMetadataList []NicMetadata `json:"nic_metadata_list,omitempty"`
	// Deployment Template holds the attributes specific to partner for which the service is created. These attributes are opaque to NSX Manager.
	DeploymentTemplate []DeploymentTemplate `json:"deployment_template"`
	// Partner needs to specify the Service VM version which will get deployed.
	SvmVersion string `json:"svm_version,omitempty"`
}
