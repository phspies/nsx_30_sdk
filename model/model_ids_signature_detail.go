package nsxt

type IdsSignatureDetail struct {
	// Product affected by the signature.
	AffectedProduct string `json:"affected_product,omitempty"`
	// Source-destination direction.
	Direction string `json:"direction,omitempty"`
	// Protocol used in the packet analysis.
	Protocol string `json:"protocol,omitempty"`
	// Class type of the signature.
	ClassType string `json:"class_type,omitempty"`
	// Signature enabled.
	Enabled bool `json:"enabled,omitempty"`
	// Packet analysis action
	Action string `json:"action,omitempty"`
	// Vendor assigned classification tag.
	Tag []string `json:"tag,omitempty"`
	// Family of the malware tracked in the signature.
	MalwareFamily string `json:"malware_family,omitempty"`
	// Name of the signature.
	Name string `json:"name,omitempty"`
	// VMware defined signature category.
	Category []string `json:"category,omitempty"`
	// Signature CVSSV3 score.
	Cvssv3 string `json:"cvssv3,omitempty"`
	// Signature CVSSV2 score.
	Cvssv2 string `json:"cvssv2,omitempty"`
	// VMware defined signature severity.
	Severity string `json:"severity,omitempty"`
	// The revision of the signature
	SignatureRevision int64 `json:"signature_revision,omitempty"`
	// Performance impact of the signature.
	PerformanceImpact string `json:"performance_impact,omitempty"`
	// Flow established from server, from client etc.
	Flow string `json:"flow,omitempty"`
	// Signature vendor set severity of the signature rule.
	SignatureSeverity string `json:"signature_severity,omitempty"`
	// List of mitre attack URLs pertaining to signature.
	Urls []string `json:"urls,omitempty"`
	// Signature policy.
	Policy []string `json:"policy,omitempty"`
	// Target of the attack tracked in the signature.
	AttackTarget string `json:"attack_target,omitempty"`
	// Unique ID of the signature rule.
	SignatureId int64 `json:"signature_id,omitempty"`
	// CVE of the signature.
	Cves []string `json:"cves,omitempty"`
	// Signature type.
	Type_ []string `json:"type,omitempty"`
	// IDSSignatureDetail resource type.
	ResourceType string `json:"resource_type,omitempty"`
}
