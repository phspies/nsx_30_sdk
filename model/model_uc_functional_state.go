package nsxt

// Upgrade coordinator Uc functional State.
type UcFunctionalState struct {
	// function state of the upgrade coordinator
	State string `json:"state,omitempty"`
	// error message that explains why UC is on standby mode.
	ErrorMessage string `json:"error_message,omitempty"`
}
