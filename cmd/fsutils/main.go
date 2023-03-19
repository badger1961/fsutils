package main

import (
	"errors"
	"flag"
	"os"

	"gitflic.com/aag031/test_player/internal/fsulog"
	"gitflic.com/aag031/test_player/internal/fsutils"
)

const VERSION = "0.0.1"

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

	err = fsutils.BackupFolderWithTimeStamp(commandLineOptions.FolderName, commandLineOptions.ArchiveFileName)
	if err != nil {
		LOG_ERROR.Printf("Finished with error %s", err.Error())
		os.Exit(1)
	}
}

type CommandLineOptions struct {
	FolderName      string
	ArchiveFileName string
	IsBackUpMode    bool
	IsVersionPrint  bool
}

func parseCommandLine() (*CommandLineOptions, error) {
	var folderName string
	var archiveFileName string
	var isBackUpMode bool
	var isVersionPrint bool

	flag.StringVar(&folderName, "dir", "", "--folder name for processing")
	flag.StringVar(&folderName, "d", "", "--folder name for processing")
	flag.StringVar(&archiveFileName, "out", "", "--folder name for processing")
	flag.StringVar(&archiveFileName, "o", "", "--folder name for processing")
	flag.BoolVar(&isBackUpMode, "backup", true, "backup mode zip target folder and setup timestamp for archive")
	flag.BoolVar(&isBackUpMode, "b", true, "backup mode zip target folder and setup timestamp for archive")
	flag.BoolVar(&isVersionPrint, "version", false, "--version print version numbert and exit")
	flag.BoolVar(&isVersionPrint, "v", false, "-v print version number and exit")

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
	commandLineOptions.ArchiveFileName = archiveFileName
	commandLineOptions.IsBackUpMode = isBackUpMode
	return commandLineOptions, nil
}

func initCommandLineOptions() *CommandLineOptions {
	var commandLineOptions = new(CommandLineOptions)
	commandLineOptions.FolderName = ""
	commandLineOptions.ArchiveFileName = ""
	commandLineOptions.IsBackUpMode = false
	commandLineOptions.IsVersionPrint = false
	return commandLineOptions
}
