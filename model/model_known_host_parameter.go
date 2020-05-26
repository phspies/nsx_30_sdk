package nsxt

type KnownHostParameter struct {
	// Known host hostname or IP address
	Host string `json:"host"`
	// Known host port
	Port int64 `json:"port,omitempty"`
}
