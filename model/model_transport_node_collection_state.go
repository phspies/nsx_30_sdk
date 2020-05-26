package nsxt

// Realization state of attaching or detaching Transport node profile on compute collection.
type TransportNodeCollectionState struct {
	// Transport node profile(TNP) will not be applied to a discovered node(DN) if some validations are not passed. In this case transport node is not created or existing transport node is not updated with TNP configurations.
	ValidationErrors []ValidationError `json:"validation_errors,omitempty"`
	// If the host preparation or transport node creation is going on for any host then state will be \"IN_PROGRESS\".  If setting desired state of the transport node failed for any of the host then state will be \"FAILED_TO_CREATE\"  If realization of transport node failed for any of the host then state will be \"FAILED_TO_REALIZE\"  If Transport node is successfully created for all of the hosts in compute collection then state will be \"SUCCESS\"  You can override the configuration for one or more hosts in the compute collection by update TN(transport node) request on individual TN. If TN is successfully created for all hosts in compute collection and one or more hosts have overridden configuration then transport node collection state will be \"PROFILE_MISMATCH\". 
	State string `json:"state,omitempty"`
	// Average of all transport node deployment progress in a cluster. Applicable only if transport node profile is applied on a cluster.
	AggregateProgressPercentage int64 `json:"aggregate_progress_percentage,omitempty"`
}
