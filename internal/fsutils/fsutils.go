package fsutils

import (
	"gitflic.com/aag031/test_player/internal/fsulog"
)

var LOG = fsulog.GetLogger()

func BackupFolderWithTimeStamp(folderName string) {
	LOG.Infof("Backup folder %s\n", folderName)
}
