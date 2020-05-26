package nsxt

type PendingChange struct {
	// Request identifier of the API which modified the entity.
	RequestId string `json:"request_id,omitempty"`
}
