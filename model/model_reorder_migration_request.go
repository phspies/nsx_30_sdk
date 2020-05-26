package nsxt

type ReorderMigrationRequest struct {
	// flag indicating whether the migration unit group/migration unit is to be placed before or after the specified migration unit group/migration unit
	IsBefore bool `json:"is_before,omitempty"`
	// id of the migration unit group/migration unit before/after which the migration unit group/migration unit is to be placed
	Id string `json:"id"`
}
