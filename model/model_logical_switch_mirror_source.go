package nsxt

type LogicalSwitchMirrorSource struct {
	// Resource types of mirror source
	ResourceType string `json:"resource_type"`
	// Please note as logical port attached with vmk interface is unsupported as mirror source, traffic from those ports on source logical switch will not be mirrored. 
	SwitchId string `json:"switch_id"`
}
