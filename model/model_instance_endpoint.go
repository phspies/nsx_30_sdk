package nsxt

// An InstanceEndpoint belongs to one ServiceInstance and represents a redirection target for a Rule. For Example - It can be an L3 Destination. Service Attachments is required for a InstanceEndpoint of type LOGICAL, and deployed_to if its a VIRTUAL InstanceEndpoint.
type InstanceEndpoint struct {
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
	// Id(s) of the Service Attachment where this enndpoint is connected to. Service Attachment is mandatory for LOGICAL Instance Endpoint.
	ServiceAttachments []ResourceReference `json:"service_attachments,omitempty"`
	// Target IPs on an interface of the Service Instance.
	TargetIps []IpInfo `json:"target_ips"`
	// LOGICAL - It must be created with a ServiceAttachment and identifies a destination connected to the Service Port of the ServiceAttachment, through the ServiceAttachment's Logical Switch. VIRTUAL - It represents a L3 destination the router can route to but does not provide any further information about its location in the network. Virtual InstanceEndpoints are used for redirection targets that are not connected to Service Ports, such as the next-hop routers on the Edge uplinks.
	EndpointType string `json:"endpoint_type,omitempty"`
	// The Service instancee with which the instance endpoint is associated.
	ServiceInstanceId string `json:"service_instance_id,omitempty"`
	// Link Ids are mandatory for VIRTUAL Instance Endpoint. Even though VIRTUAL, the Instance Endpoint should be connected/accessible through an NSX object. The link id is this NSX object id. Example - For North-South Service Insertion, this is the LogicalRouter Id through which the targetIp/L3 destination accessible.
	LinkIds []ResourceReference `json:"link_ids,omitempty"`
}
