package nsxt

type LogicalPortMacTableCsvEntry struct {
	// The type of the MAC address
	MacType string `json:"mac_type"`
	// The MAC address
	MacAddress string `json:"mac_address"`
}
