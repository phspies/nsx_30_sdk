package nsxt

type IPv6Profiles struct {
	// Identifier of Duplicate Address Detection profile. DAD profile has various configurations related to duplicate address detection. If no profile is associated manually to the router, then the system defined default DAD profile will be automatically applied. 
	DadProfileId string `json:"dad_profile_id,omitempty"`
	// Identifier of Neighbor Discovery Router Advertisement profile. NDRA profile has various configurations required for router advertisement. If no profile is associated manually to the router, then the system defined default NDRA profile will be automatically applied. 
	NdraProfileId string `json:"ndra_profile_id,omitempty"`
}
