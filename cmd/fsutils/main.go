package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"gitflic.com/aag031/test_player/internal/fsulog"
	"gitflic.com/aag031/test_player/internal/fsutils"
)

const VERSION = "0.0.3"

var LOG_DEBUG = fsulog.GetDebugLogger()
var LOG_ERROR = fsulog.GetErrorLogger()
var LOG_INFO = fsulog.GetInfoLogger()
var LOG_WARN = fsulog.GetWarnLogger()

func main() {

	commandLineOptions, err := parseCommandLine()
	if err != nil {
		os.Exit(1)
	}

	if commandLineOptions.IsVersionPrint {
		LOG_INFO.Printf("Current version is : %s ", VERSION)
		os.Exit(1)
	}

	if !commandLineOptions.IsBackUpMode {
		LOG_ERROR.Println(("Hmm .... No actual mode.So EXIT"))
		os.Exit(1)
	}

	if len(commandLineOptions.FolderName) == 0 {
		LOG_ERROR.Println("Backupmode is true, but target folder is not set")
		os.Exit(1)
	}

	err = fsutils.BackupFolderWithTimeStamp(commandLineOptions.FolderName, commandLineOptions.outputArchiveName)
	if err != nil {
		LOG_ERROR.Printf("Finished with error %s", err.Error())
		os.Exit(1)
	}
}

type CommandLineOptions struct {
	FolderName        string
	outputFolder      string
	outputArchiveName string
	IsBackUpMode      bool
	IsVersionPrint    bool
}

func parseCommandLine() (*CommandLineOptions, error) {
	var folderName string
	var outputFolder string
	var isBackUpMode bool
	var isVersionPrint bool

	flag.Usage = func() {
		_, baseName := filepath.Split(os.Args[0])
		fmt.Printf("Usage of %s:\n", baseName)
		flag.PrintDefaults()
	}

	flag.StringVar(&folderName, "dir", "", "folder name which will be archived")
	flag.StringVar(&folderName, "d", "", "folder name which will be archived")
	flag.StringVar(&outputFolder, "out", "", "folder name where new archive will be placed")
	flag.StringVar(&outputFolder, "o", "", "folder name where new archive will be placed")
	flag.BoolVar(&isBackUpMode, "backup", true, "setup the backup mode. zip archive will be created and  timestamp add to archive name")
	flag.BoolVar(&isBackUpMode, "b", true, "setup the backup mode. zip archive will be created and  timestamp add to archive name")
	flag.BoolVar(&isVersionPrint, "version", false, "print version numbert and exit")
	flag.BoolVar(&isVersionPrint, "v", false, "print version number and exit")

	if len(os.Args) == 1 {
		flag.Usage()
	}

	flag.Parse()
	if !flag.Parsed() {
		err := errors.New("Hmm could not parse command line")
		return nil, err
	}
	commandLineOptions := initCommandLineOptions()

	if isVersionPrint {
		commandLineOptions.IsVersionPrint = true
		return commandLineOptions, nil
	}

	commandLineOptions.FolderName = folderName
	commandLineOptions.outputFolder = outputFolder
	commandLineOptions.IsBackUpMode = isBackUpMode
	fileInfo, err := os.Stat(outputFolder)
	if err != nil {
		LOG_ERROR.Printf("Hmm ... could not verify output folder %s\n", outputFolder)
		return nil, err
	}

	if !fileInfo.IsDir() {
		LOG_ERROR.Printf("Output %s should be a folder\n", outputFolder)
		return nil, errors.New("Output point should be a folder")
	}

	baseNameArchive := filepath.Base(folderName)
	baseNameArchive = baseNameArchive + ".zip"
	commandLineOptions.outputArchiveName = filepath.Join(outputFolder, baseNameArchive)

	return commandLineOptions, nil
}

func initCommandLineOptions() *CommandLineOptions {
	var commandLineOptions = new(CommandLineOptions)
	commandLineOptions.FolderName = ""
	commandLineOptions.outputFolder = ""
	commandLineOptions.outputArchiveName = ""
	commandLineOptions.IsBackUpMode = false
	commandLineOptions.IsVersionPrint = false
	return commandLineOptions
}
