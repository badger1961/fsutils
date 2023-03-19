package main

import (
	"errors"
	"flag"
	"os"

	"gitflic.com/aag031/test_player/internal/fsulog"
	"gitflic.com/aag031/test_player/internal/fsutils"
)

const VERSION = "0.0.1"

var LOG = fsulog.GetLogger()

func main() {
	LOG = fsulog.GetLogger()
	commandLineOptions, err := parseCommandLine()
	if err != nil {
		os.Exit(1)
	}

	if commandLineOptions.IsVersionPrint {
		LOG.Infof("Current version is : %s ", VERSION)
		os.Exit(1)
	}

	if !commandLineOptions.IsBackUpMode {
		LOG.Errorf(("Hmm .... No actual mode.So EXIT"))
		os.Exit(1)
	}

	if len(commandLineOptions.FolderName) == 0 {
		LOG.Errorln("Backupmode is true, but target folder is not set")
		os.Exit(1)
	}

	fsutils.BackupFolderWithTimeStamp(commandLineOptions.FolderName)
}

type CommandLineOptions struct {
	FolderName     string
	IsBackUpMode   bool
	IsVersionPrint bool
}

func parseCommandLine() (*CommandLineOptions, error) {
	var folderName string
	var isBackUpMode bool
	var isVersionPrint bool

	flag.StringVar(&folderName, "dir", "", "--folder name for processing")
	flag.StringVar(&folderName, "d", "", "--folder name for processing")
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
	commandLineOptions.IsBackUpMode = isBackUpMode
	return commandLineOptions, nil
}

func initCommandLineOptions() *CommandLineOptions {
	var commandLineOptions = new(CommandLineOptions)
	commandLineOptions.FolderName = ""
	commandLineOptions.IsBackUpMode = false
	commandLineOptions.IsVersionPrint = false
	return commandLineOptions
}
