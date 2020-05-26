package nsxt

// Relative to the form factor pre-defined reservation value. To reduce reservation of a VM to 50 percent, a user may specify 50 instead of the absolute number relevant for the edge form factor. 
type MemoryReservation struct {
	// Memory reserved relative to the default reservation of 100 percent. For example, take an edge virtual machine of medium form factor. By default, an edge of medium form factor is configured with 8 GB of memory and with reservation of 100 percent. So, 8 GB of memory is reserved. If you specify reservation_percentage value as 50 percent, then 4 GB of memory will be reserved. 
	ReservationPercentage int32 `json:"reservation_percentage,omitempty"`
}
