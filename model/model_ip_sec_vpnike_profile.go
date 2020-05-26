package nsxt

// IKE Profile is a reusable profile that captures IKE phase one negotiation parameters. Any changes affects all IPSec VPN sessions consuming this profile.
type IpSecVpnikeProfile struct {
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
	// Algorithm to be used for message digest during Internet Key Exchange(IKE) negotiation. Default is SHA2_256.
	DigestAlgorithms []string `json:"digest_algorithms,omitempty"`
	// Encryption algorithm is used during Internet Key Exchange(IKE) negotiation. Default is AES_128.
	EncryptionAlgorithms []string `json:"encryption_algorithms,omitempty"`
	// Diffie-Hellman group to be used if PFS is enabled. Default is GROUP14.
	DhGroups []string `json:"dh_groups,omitempty"`
	// Life time for security association. Default is 86400 seconds (1 day).
	SaLifeTime int64 `json:"sa_life_time,omitempty"`
	// IKE protocol version to be used. IKE-Flex will initiate IKE-V2 and responds to both IKE-V1 and IKE-V2.
	IkeVersion string `json:"ike_version,omitempty"`
}
