package nsxt

type BackupUiFramesInfoList struct {
	// List of backup frames(and metadata) to be displayed in UI
	BackupFramesList []BackupUiFramesInfo `json:"backup_frames_list,omitempty"`
}
