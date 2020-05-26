package nsxt

// Protocol on which a particular ServiceInsertion Rule should apply to.
type ServiceInsertionService struct {
	Service *NsServiceElement `json:"service,omitempty"`
}
