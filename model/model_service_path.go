package nsxt

// An instance of service chain that consists of forward and reverse service paths.
type ServicePath struct {
	ReversePath *UnidirectionalServicePath `json:"reverse_path,omitempty"`
	// Uuid of a service chain.
	ServiceChainUuid string `json:"service_chain_uuid,omitempty"`
	ForwardPath *UnidirectionalServicePath `json:"forward_path,omitempty"`
	// Unique identifier of a service path.
	ServicePathId int64 `json:"service_path_id,omitempty"`
	// A unique id of a service chain.
	ServiceChainId int64 `json:"service_chain_id,omitempty"`
}
