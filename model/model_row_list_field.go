package nsxt

// Root of the api result set for forming rows.
type RowListField struct {
	// Short name or alias of row list field, if any. If unspecified, the row list field can be referenced by its index in the array of row list fields as $<index> (for example, $0).
	Alias string `json:"alias,omitempty"`
	// JSON path to the root of the api result set for forming rows.
	Path string `json:"path"`
}
