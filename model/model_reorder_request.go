package nsxt

type ReorderRequest struct {
	// flag indicating whether the upgrade unit group/upgrade unit is to be placed before or after the specified upgrade unit group/upgrade unit
	IsBefore bool `json:"is_before,omitempty"`
	// id of the upgrade unit group/upgrade unit before/after which the upgrade unit group/upgrade unit is to be placed
	Id string `json:"id"`
}
