package nsxt

type NodeUserSettings struct {
	// To configure username, you must provide this property together with <b>cli_password</b>. Username must contain ASCII characters only. 
	CliUsername string `json:"cli_username,omitempty"`
	// The default username is \"audit\". To configure username, you must provide this property together with <b>audit_password</b>. Username must contain ASCII characters only. 
	AuditUsername string `json:"audit_username,omitempty"`
	// Password for the node root user. For deployment, this property is required. After deployment, this property is ignored, and the node cli must be used to change the password. The password specified must be at least 12 characters in length and must contain at least one lowercase, one uppercase, one numeric character and  one special character (except quotes). Passwords based on dictionary words and palindromes are invalid. 
	RootPassword string `json:"root_password,omitempty"`
	// Password for the node cli user. For deployment, this property is required. After deployment, this property is ignored, and the node cli must be used to change the password. The password specified must be at least 12 characters in length and must contain at least one lowercase, one uppercase, one numeric character and one special character (except quotes). Passwords based on dictionary words and palindromes are invalid. 
	CliPassword string `json:"cli_password,omitempty"`
	// Password for the node audit user. For deployment, this property is required. After deployment, this property is ignored, and the node cli must be used to change the password. The password specified must be at least 12 characters in length and must contain at least one lowercase, one uppercase, one numeric character and one special character (except quotes). Passwords based on dictionary words and palindromes are invalid. 
	AuditPassword string `json:"audit_password,omitempty"`
}
