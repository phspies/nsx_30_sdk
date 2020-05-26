package nsxt

// IPSec VPN tunnel profile is a reusable profile that captures phase two negotiation parameters and tunnel properties. Any changes affects all IPSec VPN sessions consuming this profile.
type IpSecVpnTunnelProfile struct {
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
	// Encapsulation Mode to be used for encryption of packet. Tunnel mode protects internal routing information by encrypting IP header of original packet.
	EncapsulationMode string `json:"encapsulation_mode,omitempty"`
	// IPSec transform specifies IPSec security protocol.
	TransformProtocol string `json:"transform_protocol,omitempty"`
	// Algorithm to be used for message digest. Default digest algorithm is implicitly covered by default encryption algorithm \"AES_GCM_128\".
	DigestAlgorithms []string `json:"digest_algorithms,omitempty"`
	// Encryption algorithm to encrypt/decrypt the messages exchanged between IPSec VPN initiator and responder during tunnel negotiation. Default is AES_GCM_128.
	EncryptionAlgorithms []string `json:"encryption_algorithms,omitempty"`
	// If true, perfect forward secrecy (PFS) is enabled.
	EnablePerfectForwardSecrecy bool `json:"enable_perfect_forward_secrecy,omitempty"`
	// Diffie-Hellman group to be used if PFS is enabled. Default is GROUP14.
	DhGroups []string `json:"dh_groups,omitempty"`
	// Defragmentation policy helps to handle defragmentation bit present in the inner packet. COPY copies the defragmentation bit from the inner IP packet into the outer packet. CLEAR ignores the defragmentation bit present in the inner packet.
	DfPolicy string `json:"df_policy,omitempty"`
	// SA life time specifies the expiry time of security association. Default is 3600 seconds. 
	SaLifeTime int64 `json:"sa_life_time,omitempty"`
}
