package nsxt

// Dead peer detection (DPD) is a method that allows detection of unreachable internet key excahnge (IKE) peers. Any changes affects all IPSec VPN sessions consuming this profile.
type IpSecVpndpdProfile struct {
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
	// Maximum number of DPD messages retry attempts. This value is applicable for both dpd probe modes, periodic and on-demand.
	RetryCount int64 `json:"retry_count,omitempty"`
	// If true, enable dead peer detection.
	Enabled bool `json:"enabled,omitempty"`
	// DPD probe mode is used to query the liveliness of the peer. Two modes are possible - PERIODIC - is used to query the liveliness of the peer at regular intervals (dpd_probe_interval). It does not take into consideration traffic coming from the peer. The benefit of this mode over the on-demand mode is earlier detection of dead peers. However, use of periodic DPD incurs extra overhead. When communicating to large numbers of peers, please consider using on-demand DPD instead. ON_DEMAND - is used to query the liveliness of the peer by instructing the local endpoint to send DPD message to a peer if there is traffic to send to the peer AND the peer was idle for dpd_probe_interval seconds (i.e. there was no traffic from the peer for dpd_probe_interval seconds) 
	DpdProbeMode string `json:"dpd_probe_mode,omitempty"`
	// When the DPD probe mode is periodic, this interval is the number of seconds between DPD messages. When the DPD probe mode is on-demand, this interval is the number of seconds during which traffic is not received from the peer before DPD retry messages are sent if there is IPSec traffic to send. For PERIODIC Mode:    Minimum: 3    Maximum: 360    Default: 60 For ON_DEMAND Mode:    Minimum: 1    Maximum: 10    Default: 3 
	DpdProbeInterval int64 `json:"dpd_probe_interval,omitempty"`
}
