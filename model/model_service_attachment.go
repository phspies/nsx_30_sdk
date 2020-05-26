package nsxt

// A ServiceAttachment represents a point on NSX entity (Example - Edge Router) to which ServiceInstance can be connected through an InstanceEndpoint. Example - In VMWare Hybric Cloud Extention (HCX) use case, HCX appliances connect to this Service Attachment Point. We do not handle the lifecycle of these appliance/s.
type ServiceAttachment struct {
	// Indicates system owned resource
	SystemOwned bool `json:"_system_owned,omitempty"`
	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`
	// Description of this resource
	Description string `json:"description,omitempty"`
	// Opaque identifiers meaningful to the API user
	Tags []Tag `json:"tags,omitempty"`
	// ID of the user who created this resource
	CreateUser string `json:"_create_user,omitempty"`
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed             to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed                 to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super                    user and can modify it, but only when providing                    the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this           entity. 
	Protection string `json:"_protection,omitempty"`
	// Timestamp of resource creation
	CreateTime int64 `json:"_create_time,omitempty"`
	// Timestamp of last modification
	LastModifiedTime int64 `json:"_last_modified_time,omitempty"`
	// ID of the user who last modified this resource
	LastModifiedUser string `json:"_last_modified_user,omitempty"`
	// Unique identifier of this resource
	Id string `json:"id,omitempty"`
	// The type of this resource.
	ResourceType string `json:"resource_type,omitempty"`
	// List of LogicalRouters to be connected to the ServicePlane logical switch via a ServiceLink.
	LogicalRouters []ResourceReference `json:"logical_routers,omitempty"`
	LogicalSwitch *ResourceReference `json:"logical_switch,omitempty"`
	// Local IPs associated with this Service Attachment.
	LocalIps []IpInfo `json:"local_ips,omitempty"`
	ServicePort *ResourceReference `json:"service_port,omitempty"`
	DeployedTo *ResourceReference `json:"deployed_to"`
	// UP - A Service Attachment will have its Service Port - UP and with a configured IP address. DOWN - An Inactive ServiceAttachment has its Service Port - DOWN. It can be used to connect set of appliances that do not need to exchange traffic to/from/through the Edge node.
	AttachmentStatus string `json:"attachment_status,omitempty"`
}
