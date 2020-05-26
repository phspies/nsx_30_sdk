package nsxt

type DirectoryAdGroup struct {
	// Domain sync node under which this directory group is located. We currently sync only from Root node and hence this attribute doesn't have a specific value set.
	DomainSyncNodeId string `json:"domain_sync_node_id,omitempty"`
	// Directory group distinguished name
	DistinguishedName string `json:"distinguished_name"`
	// Domain ID this directory group belongs to.
	DomainId string `json:"domain_id"`
	// Directory group resource type comes from multiple sub-classes extending this base class. For example, DirectoryAdGroup is one accepted resource_type. If there are more sub-classes defined, they will also be accepted resource_type.
	ResourceType string `json:"resource_type"`
	// Each active directory domain has a domain naming context (NC), which contains domain-specific data. The root of this naming context is represented by a domain's distinguished name (DN) and is typically referred to as the NC head.
	DomainName string `json:"domain_name"`
	// GUID is a 128-bit value that is unique not only in the enterprise but also across the world. GUIDs are assigned to every object created by Active Directory, not just User and Group objects.
	ObjectGuid string `json:"object_guid"`
	// A security identifier (SID) is a unique value of variable length used to identify a trustee. A SID consists of the following components - The revision level of the SID structure; A 48-bit identifier authority value that identifies the authority that issued the SID; A variable number of subauthority or relative identifier (RID) values that uniquely identify the trustee relative to the authority that issued the SID.
	SecureId string `json:"secure_id"`
}
