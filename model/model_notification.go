package nsxt

type Notification struct {
	// A string identifying feature_name.notification_name to indicate a notification watcher is interested in receiving notifications for the URI identified by the feature_name.notification_name.
	NotificationId string `json:"notification_id,omitempty"`
	// Optional list of URIs
	UriFilters []string `json:"uri_filters,omitempty"`
}
