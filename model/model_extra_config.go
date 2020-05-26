package nsxt

// Extra config is intended for supporting vendor specific configuration on the data path, it can be set as key value string pairs on either logical switch or logical port. If it was set on logical switch, it will be inherited automatically by logical ports in it. Also logical port setting will override logical switch setting if specific key was dual set on both logical switch and logical port. 
type ExtraConfig struct {
	ConfigPair *KeyValuePair `json:"config_pair"`
}
