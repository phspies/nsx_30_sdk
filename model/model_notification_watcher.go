package nsxt

type NotificationWatcher struct {
	// Indicates system owned resource
	SystemOwned bool `json:"_system_owned,omitempty"`
	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`
	// Optional description that can be associated with this NotificationWatcher.
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
	// System generated identifier to identify a notification watcher uniquely. 
	Id string `json:"id,omitempty"`
	// The type of this resource.
	ResourceType string `json:"resource_type,omitempty"`
	// Optional time duration (in seconds) to specify request timeout to notification watcher. If the send reaches the timeout, will try to send refresh_needed as true in the next time interval. The default value is 30 seconds.
	SendTimeout int64 `json:"send_timeout,omitempty"`
	// URI notification requests should be made on the specified server.
	Uri string `json:"uri"`
	// Contains the hex-encoded SHA256 thumbprint of the HTTPS certificate. It must be specified if use_https is set to true.
	CertificateSha256Thumbprint string `json:"certificate_sha256_thumbprint,omitempty"`
	// Type of method notification requests should be made on the specified server. The value must be set to POST.
	Method string `json:"method"`
	// Optional time interval (in seconds) for which notification URIs will be accumulated. At the end of the time interval the accumulated notification URIs will be sent to this NotificationWatcher in the form of zero (nothing accumulated) or more notification requests as soon as possible. If it is not specified, the NotificationWatcher should expected to receive notifications at any time.
	SendInterval int64 `json:"send_interval,omitempty"`
	// If the number of notification URIs accumulated in specified send_interval exceeds max_send_uri_count, then multiple notification requests (each with max_send_uri_count or less number of notification URIs) will be sent to this NotificationWatcher. The default value is 5000.
	MaxSendUriCount int64 `json:"max_send_uri_count,omitempty"`
	AuthenticationScheme *NotificationAuthenticationScheme `json:"authentication_scheme"`
	// IP address or fully qualified domain name of the partner/customer watcher.
	Server string `json:"server"`
	// Optional integer port value to specify a non-standard HTTP or HTTPS port.
	Port int64 `json:"port,omitempty"`
	// Optional field, when set to true indicates REST API server should use HTTPS.
	UseHttps bool `json:"use_https,omitempty"`
}