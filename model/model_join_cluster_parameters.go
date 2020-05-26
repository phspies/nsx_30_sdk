package nsxt

// To join a new node to the NSX cluster, issue a JoinCluster API on the new node. The JoinCluster API takes this object as a parameter. Provide the ID of the NSX cluster you want the new node to join and the IP address of one of the nodes already in that cluster. The Cluster Boot Manager running on the new node will then add the new node to the NSX cluster by making a AttachClusterNode REST API call on the node that is already part of the cluster. In order to make a REST API call to the node in the cluster, the Cluster Boot Manager will need username and password of a priviledged user on the node in the cluster. In place of a username and password, Cluster Boot Manager could also use a OAuth token provided. The Cluster Boot Manager needs either the username and password or the OAuth token to make the REST call but not both.
type JoinClusterParameters struct {
	// Username on the cluster node
	Username string `json:"username,omitempty"`
	// SHA256 Thumbprint of the API certificate of the cluster node
	CertficateSha256Thumbprint string `json:"certficate_sha256_thumbprint"`
	// Limited time OAuth token instead of the username/password
	Token string `json:"token,omitempty"`
	// UUID of the cluster to join
	ClusterId string `json:"cluster_id"`
	// Password of the user on the cluster node
	Password string `json:"password,omitempty"`
	// IP address of a node already part of the cluster to join
	IpAddress string `json:"ip_address"`
	// API port on the cluster node
	Port int64 `json:"port,omitempty"`
}
