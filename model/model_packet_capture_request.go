package nsxt

// This type is used to create packet request on give node. Need to specify related parameters according to the capture point.
type PacketCaptureRequest struct {
	// Define the transport node to capture data.
	Node string `json:"node,omitempty"`
	// Define the capture direction. Support three types INPUT/OUTPUT/DUAL.
	Direction string `json:"direction,omitempty"`
	// Define the packet capture duration time. After the capture duration time, the capture process will stop working.
	Capduration int32 `json:"capduration,omitempty"`
	// Define the packet capture amount size.
	Capamount int32 `json:"capamount,omitempty"`
	// This type is used to differenite the incoming request from CLI/UI.
	Capsource string `json:"capsource"`
	// Define the transport node to capture data.
	NodeIp string `json:"node_ip,omitempty"`
	// Define the capture value of given capture point.
	Capvalue string `json:"capvalue,omitempty"`
	// Define the capture filter type. Support PRE/POST mode.
	Filtertype string `json:"filtertype,omitempty"`
	// Define the point to capture data.
	Cappoint string `json:"cappoint"`
	// Define the packet capture file size limit.
	Capfilesize int32 `json:"capfilesize,omitempty"`
	Options *PacketCaptureOptionList `json:"options,omitempty"`
	// Set the stream port to receive the capture packet. The STREAM mode is based on GRE-in-UDP Encapsulation(RFC8086). Packets are sent to UDP port 4754.
	Streamport int32 `json:"streamport,omitempty"`
	// Define the rate of packet capture process.
	Caprate int32 `json:"caprate,omitempty"`
	// The CPU core id on Edge node.
	Capcore int32 `json:"capcore,omitempty"`
	// Limit the number of bytes captured from each packet.
	Capsnaplen int32 `json:"capsnaplen,omitempty"`
	// Set the stream address to receive the capture packet.
	Streamaddress string `json:"streamaddress,omitempty"`
	// Define the capture streaming mode. The STREAM mode will send the data to given stream address and port. And the STANDALONE mode will save the capture file in local folder.
	Capmode string `json:"capmode,omitempty"`
}
