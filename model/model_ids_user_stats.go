package nsxt

// List of Users logged into VMs where intrusions of a given signature were detected. 
type IdsUserStats struct {
	// Number of unique users logged into VMs on which a particular signature was detected.
	Count int64 `json:"count,omitempty"`
	// List of users logged into VMs on which a particular signature was detected.
	UserList []string `json:"user_list,omitempty"`
}
