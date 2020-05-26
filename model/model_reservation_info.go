package nsxt

// CPU and Memory resource configuration is defined per Edge VM form factor. These resources are reserved 100 percent by default with Normal VM importance. Resource reservation tuning provides a means to optimize resource utilization and workaround hard resource limits. This solution should be used as a temporary workaround. It is recommended to add more resources to the compute cluster and change the reservation back to 100 percent for optimal performance. 
type ReservationInfo struct {
	CpuReservation *CpuReservation `json:"cpu_reservation,omitempty"`
	MemoryReservation *MemoryReservation `json:"memory_reservation,omitempty"`
}
