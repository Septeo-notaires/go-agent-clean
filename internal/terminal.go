package internal

import (
	"errors"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

var excludePath []string = []string{
	"_temp",
	"_tool",
	"_tasks",
	"ReleaseRootMapping",
	"SourceRootMapping",
}

type (
	ITerminal interface {
		CleanAgent(path string) error
		ShutdownAgent(agentName string) error
		StartAgent(agentName string) error
	}

	WindowsTerminal struct {
		m *mgr.Mgr
	}

	LinuxTerminal struct {
	}
)

// Windows Terminal

func NewWindowTerminal() WindowsTerminal {
	m, err := mgr.Connect()
	if err != nil {
		panic(err)
	}

	return WindowsTerminal{
		m: m,
	}
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

func (w *WindowsTerminal) ShutdownAgent(agentName string) error {
	log.Info().Msg("Shutdown : " + agentName)
	service, err := w.m.OpenService(agentName)
	if err != nil {
		return err
	}

	_, err = service.Control(svc.Stop)
	time.Sleep(5 * time.Second)
	if err != nil {
		return err
	}
	return nil
}

func (w *WindowsTerminal) StartAgent(agentName string) error {
	log.Info().Msg("Start : " + agentName)
	service, err := w.m.OpenService(agentName)
	if err != nil {
		return err
	}

	err = service.Start()
	if err != nil {
		return err
	}
	return nil
}

// Linux Terminal
func NewLinuxTerminal() LinuxTerminal {
	return LinuxTerminal{}
}

func (w *LinuxTerminal) CleanAgent(path string) error {
	return nil
}

func (w *LinuxTerminal) ShutdownAgent(agentName string) error {
	return nil
}

func (w *LinuxTerminal) StartAgent(agentName string) error {
	return nil
}

//Common Function

func isExcludedName(dirName string) bool {
	for _, value := range excludePath {
		if value == dirName {
			return true
		}
	}
	return false
}
