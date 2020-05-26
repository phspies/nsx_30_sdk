package nsxt

type DhcpHeader struct {
	// This is used to specify the general type of message. A client sending request to a server uses an op code of BOOTREQUEST, while a server replying uses an op code of BOOTREPLY.
	OpCode string `json:"op_code,omitempty"`
}
