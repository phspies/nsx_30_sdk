package nsxt

// Type to define services associated with every rule
type PbrService struct {
	Service *NsServiceElement `json:"service,omitempty"`
}
