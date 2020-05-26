package nsxt

// Details about an existing NSX manager to be migrated
type SourceNsxApiEndpoint struct {
	// VC port that will be used to fetch details.
	VcPort int32 `json:"vc_port,omitempty"`
	// Username for connecting to VC.
	VcUsername string `json:"vc_username"`
	// IP address or host name of VC.
	VcIp string `json:"vc_ip"`
	// IP address or hostname of a source NSX API endpoint. This field is not applicable in case of vSphere network migration.
	Ip string `json:"ip,omitempty"`
	// Auth token used to make REST calls to source NSX API endpoint. This field is not applicable in case of vSphere network migration.
	AuthToken string `json:"auth_token,omitempty"`
	// Signifies Universal Sync role status (STANDALONE, PRIMARY, SECONDARY) of a source NSX API endpoint.
	NsxSyncrole string `json:"nsx_syncrole,omitempty"`
	// Build version of VC.
	VcVersion string `json:"vc_version,omitempty"`
	// Username for connecting to NSX manager. This field is not applicable in case of vSphere network migration.
	NsxUsername string `json:"nsx_username,omitempty"`
	// Build version (major, minor, patch) of a source NSX API endpoint.
	NsxVersion string `json:"nsx_version,omitempty"`
	// Password for connecting to NSX manager. This field is not applicable in case of vSphere network migration.
	NsxPassword string `json:"nsx_password,omitempty"`
	// Password for connecting to VC.
	VcPassword string `json:"vc_password,omitempty"`
}
