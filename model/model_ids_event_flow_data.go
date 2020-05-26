package nsxt

// IDS event flow data specific to each IDS event. The data includes source ip, source port, destination ip, destination port, and protocol. 
type IdsEventFlowData struct {
	// Port on the destination VM where the traffic was sent to.
	DestinationPort int64 `json:"destination_port,omitempty"`
	// IP address of the VM that initiated the communication.
	ClientIp string `json:"client_ip,omitempty"`
	// Traffic protocol pertaining to the detected intrusion, could be TCP/UDP etc.
	Protocol string `json:"protocol,omitempty"`
	// Source port through which traffic was initiated that caused the intrusion to be detected.
	SourcePort int64 `json:"source_port,omitempty"`
	// IP address of VM on the host where IDS engine is running.
	LocalVmIp string `json:"local_vm_ip,omitempty"`
	// IP address of the source VM on the intrusion flow.
	SourceIp string `json:"source_ip,omitempty"`
	// IP address of the destination VM on the intrusion flow.
	DestinationIp string `json:"destination_ip,omitempty"`
}
