package nsxt

// Mapping of all vmk interfaces to destination networks
type VmknicNetwork struct {
	// When migrating vmks to N-VDS/logical switches, the id is the logical switch id. When migrating out of N-VDS/logical switches, the id is the vSphere Switch portgroup name in a single vSphere Standard Switch (VSS), or distributed virtual portgroup name in a single distributed virtual switch (DVS).
	DestinationNetwork string `json:"destination_network"`
	// The vmk interface name, e.g., vmk0, vmk1; the id assigned by vCenter.
	DeviceName string `json:"device_name"`
}
