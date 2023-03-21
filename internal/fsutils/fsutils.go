package fsutils

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gitflic.com/aag031/test_player/internal/fsulog"
)

var LOG_INFO = fsulog.GetInfoLogger()
var LOG_ERROR = fsulog.GetErrorLogger()

func BackupFolderWithTimeStamp(folderName string, archiveName string) error {
	LOG_INFO.Printf("Backup folder %s to %s\n", folderName, archiveName)
	err := createZipArchive(folderName, archiveName)
	if err != nil {
		LOG_ERROR.Printf("Hmm ... Could not create ZIP file %s\n", err.Error())
	}
	t := time.Now()
	suffix := t.Format("20060102150405")

	arcPath, arcFileName := filepath.Split(archiveName)
	archiveExt := filepath.Ext(arcFileName)
	arcFileBaseName := strings.TrimSuffix(arcFileName, archiveExt)
	updateArcName := filepath.Join(arcPath, arcFileBaseName+"-"+suffix+archiveExt)
	os.Rename(archiveName, updateArcName)
	return nil
}

func createZipArchive(source, target string) error {
	file, err := os.Create(target)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := zip.NewWriter(file)
	defer w.Close()

	walker := func(path string, info os.FileInfo, err error) error {
		LOG_INFO.Printf("Processing: %#v\n", path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		relativePath, err := filepath.Rel(filepath.Dir(source), path)
	        if err != nil {
                   return err
                }

		f, err := w.Create(relativePath)
		if err != nil {
			return err
		}

		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}

		return nil
	}
	err = filepath.Walk(source, walker)
	if err != nil {
		panic(err)
	}

	return nil
}
