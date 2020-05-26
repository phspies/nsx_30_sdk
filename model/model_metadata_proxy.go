package nsxt

type MetadataProxy struct {
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
	// secret to access metadata server
	Secret string `json:"secret,omitempty"`
	// The CAs referenced here must be uploaded to the truststore using the API POST /api/v1/trust-management/certificates?action=import. User needs to ensure a correct CA for this metedata server is used. The REST API can not detect a wrong CA which was used to verify a different server. If the Metadata Proxy reports an ERROR or NO_BACKUP status, user can check the metadata proxy log at transport node for a possible CA issue. 
	MetadataServerCaIds []string `json:"metadata_server_ca_ids,omitempty"`
	// If none is provided, the NSX will auto-select two edge-nodes from the given edge cluster. If user provides only one edge node, there will be no HA support. 
	EdgeClusterMemberIndexes []int64 `json:"edge_cluster_member_indexes,omitempty"`
	// The cryptographic protocols listed here are supported by the metadata proxy. The TLSv1.1 and TLSv1.2 are supported by default. 
	CryptoProtocols []string `json:"crypto_protocols,omitempty"`
	// The URL in format scheme://host:port/path. Please note, the scheme supports only http and https as of now, port supports range 3000 - 9000, inclusive. 
	MetadataServerUrl string `json:"metadata_server_url"`
	// id of attached logical port
	AttachedLogicalPortId string `json:"attached_logical_port_id,omitempty"`
	// Flag to enable the auto-relocation of standby Metadata Proxy in case of edge node failure. Only tier 1 and auto placed Metadata Proxy are considered for the relocation. 
	EnableStandbyRelocation bool `json:"enable_standby_relocation,omitempty"`
	// edge cluster uuid
	EdgeClusterId string `json:"edge_cluster_id"`
}
