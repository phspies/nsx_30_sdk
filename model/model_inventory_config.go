package nsxt

type InventoryConfig struct {
	// Soft limit on number of compute managers, which can be added, beyond which, addition of compute managers will result in warning getting logged 
	ComputeManagersSoftLimit int32 `json:"compute_managers_soft_limit"`
}
