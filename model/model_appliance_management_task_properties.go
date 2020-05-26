package nsxt

// Appliance management task properties
type ApplianceManagementTaskProperties struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Current status of the task
	Status string `json:"status,omitempty"`
	// True if response for asynchronous request is available
	AsyncResponseAvailable bool `json:"async_response_available,omitempty"`
	// Description of the task
	Description string `json:"description,omitempty"`
	// The start time of the task in epoch milliseconds
	StartTime int64 `json:"start_time,omitempty"`
	// Details about the task if known
	Details *interface{} `json:"details,omitempty"`
	// True if this task can be canceled
	Cancelable bool `json:"cancelable,omitempty"`
	// HTTP request method
	RequestMethod string `json:"request_method,omitempty"`
	// The end time of the task in epoch milliseconds
	EndTime int64 `json:"end_time,omitempty"`
	// Task progress if known, from 0 to 100
	Progress int64 `json:"progress,omitempty"`
	// A message describing the disposition of the task
	Message string `json:"message,omitempty"`
	// Name of the user who created this task
	User string `json:"user,omitempty"`
	// Identifier for this task
	Id string `json:"id,omitempty"`
	// URI of the method invocation that spawned this task
	RequestUri string `json:"request_uri,omitempty"`
}
