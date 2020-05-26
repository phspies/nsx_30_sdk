package nsxt

// Intrusions that are detected, grouped by signature. It contains the signature id, severity, name, the number of intrusions of that type and the first occurence. 
type IdsEventsBySignature struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Number of times this particular signature was detected.
	Count int64 `json:"count,omitempty"`
	// First occurence of the intrusion, in epoch milliseconds.
	FirstOccurence int64 `json:"first_occurence,omitempty"`
	// Severity of the threat covered by the signature, can be Critical, High, Medium, or Low.
	Severity string `json:"severity,omitempty"`
	// Name of the signature pertaining to the detected intrusion.
	SignatureName string `json:"signature_name,omitempty"`
	// Flag indicating an ongoing intrusion.
	IsOngoing bool `json:"is_ongoing,omitempty"`
	// Signature ID pertaining to the detected intrusion.
	SignatureId int64 `json:"signature_id,omitempty"`
	// IDSEvent resource type.
	ResourceType string `json:"resource_type,omitempty"`
}
