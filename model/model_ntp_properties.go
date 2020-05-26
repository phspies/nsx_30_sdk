package nsxt

// This object contains the list of NTP servers used by NSX nodes.
type NtpProperties struct {
	// List of NTP servers.
	Servers []string `json:"servers"`
}
