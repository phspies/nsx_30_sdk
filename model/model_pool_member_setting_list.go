package nsxt

// List of PoolMemberSetting
type PoolMemberSettingList struct {
	// List of pool member settings to be passed to add, update and remove APIs 
	Members []PoolMemberSetting `json:"members"`
}
