package nsxt

// OpenID Connect end-point specifying where to fetch the JWKS document used to validate JWT tokens for TokenBasedPrincipalIdentities. 
type OidcEndPoint struct {
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
	// Issuer of the JWT tokens for the given type. This field is fetched from the meta-data located at the oidc_uri. 
	Issuer string `json:"issuer,omitempty"`
	// The URI where the JWKS document is located that has the key used to validate the JWT signature. 
	JwksUri string `json:"jwks_uri,omitempty"`
	// Type used to distinguish the OIDC end-points by IDP.
	OidcType string `json:"oidc_type,omitempty"`
	// URI of the OpenID Connect end-point.
	OidcUri string `json:"oidc_uri"`
	// Thumbprint in SHA-256 format used to verify the server certificate at the URI. 
	Thumbprint string `json:"thumbprint"`
}
