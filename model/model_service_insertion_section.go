package nsxt

// A ServiceInsertion section composed of ServiceInsertion Rules.
type ServiceInsertionSection struct {
	// Stateful or Stateless nature of distributed service section is enforced on all rules inside the section. Layer3 sections can be stateful or stateless. Layer2 sections can only be stateless.
	Stateful bool `json:"stateful"`
	// It is a boolean flag which reflects whether a distributed service section is default section or not. Each Layer 3 and Layer 2 section will have at least and at most one default section.
	IsDefault bool `json:"is_default,omitempty"`
	// List of objects where the rules in this section will be enforced. This will take precedence over rule level appliedTo.
	AppliedTos []ResourceReference `json:"applied_tos,omitempty"`
	// Number of rules in this section.
	RuleCount int64 `json:"rule_count,omitempty"`
	// Type of the rules which a section can contain. Only homogeneous sections are supported.
	SectionType string `json:"section_type"`
	// Ensures that a three way TCP handshake is done before the data packets are sent if the value is set to be true. tcp_strict=true is supported only for stateful sections.
	TcpStrict bool `json:"tcp_strict,omitempty"`
}
