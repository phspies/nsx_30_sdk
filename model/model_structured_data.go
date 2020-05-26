package nsxt

// Structured data entry in RFC5424 log format
type StructuredData struct {
	// Audit flag of the log
	Audit string `json:"audit"`
	// Username value of the log
	Username string `json:"username,omitempty"`
	// External request Id value of the log
	EreqId string `json:"ereq_id,omitempty"`
	// Level value of the log
	Level string `json:"level,omitempty"`
	// Component value of the log
	Comp string `json:"comp"`
	// Error Code value of the log
	ErrorCode string `json:"error_code,omitempty"`
	// Sub-subcomponent value of the log
	S2comp string `json:"s2comp,omitempty"`
	// Request Id value of the log
	ReqId string `json:"req_id,omitempty"`
	// Entity Id value of the log
	EntId string `json:"ent_id,omitempty"`
	// Security flag of the log
	Security string `json:"security,omitempty"`
	// Subcomponent value of the log
	Subcomp string `json:"subcomp"`
}
