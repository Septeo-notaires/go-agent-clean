package internal

import (
	"errors"
	"os"

	"github.com/rs/zerolog/log"
)

var excludePath []string = []string{
	"_temp",
	"_tool",
	"_tasks",
	"ReleaseRootMapping",
	"SourceRootMapping",
}

type ITerminal interface {
	CleanAgent(path string) error
	ShutdownAgent(agentName string) error
	StartAgent(agentName string) error
}

// Common Function
func isExcludedName(dirName string) bool {
	for _, value := range excludePath {
		if value == dirName {
			return true
		}
	}
	return false
}

func (w *WindowsTerminal) CleanAgent(path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !fi.IsDir() {
		return errors.New("The path should be a directory")
	}

	dirEntry, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, value := range dirEntry {
		name := value.Name()
		dirPath := path + "\\" + name
		log.Info().Msg("Remove : " + dirPath)
		if !isExcludedName(name) {
			err = os.RemoveAll(dirPath)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (w *LinuxTerminal) CleanAgent(path string) error {
	return nil
}
