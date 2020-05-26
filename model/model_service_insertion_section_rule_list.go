package nsxt

// List of ServiceInsertion Rules.
type ServiceInsertionSectionRuleList struct {
	// Ensures that a three way TCP handshake is done before the data packets are sent if the value is set to be true. tcp_strict=true is supported only for stateful sections.
	TcpStrict bool `json:"tcp_strict,omitempty"`
	// List of Service Insertion rules in the section. Only homogeneous rules are supported.
	Rules []ServiceInsertionRule `json:"rules"`
}
