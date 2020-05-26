package nsxt

type FirewallSectionLock struct {
	// Comments for section lock/unlock.
	Comments string `json:"comments"`
	// Revision of the section.
	SectionRevision int64 `json:"section_revision"`
}
