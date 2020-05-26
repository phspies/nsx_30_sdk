package nsxt

// Packet capture session information.
type PacketCaptureSession struct {
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
	// Packet capture session id.
	Sessionid string `json:"sessionid"`
	// Packet capture file location.
	Filelocation string `json:"filelocation,omitempty"`
	// Packet capture file Size in bytes.
	Filesize int32 `json:"filesize,omitempty"`
	// Packet capture session name.
	Sessionname string `json:"sessionname,omitempty"`
	// Error messasge in capture.
	Errormsg string `json:"errormsg,omitempty"`
	// Timestamp when session was stopped in epoch millisecond.
	Endtime int64 `json:"endtime,omitempty"`
	Request *PacketCaptureRequest `json:"request"`
	// Timestamp when session was created in epoch millisecond.
	Starttime int64 `json:"starttime,omitempty"`
	// Packet capture session status.
	Sessionstatus string `json:"sessionstatus"`
}