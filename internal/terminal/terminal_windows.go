//go:build windows
// +build windows

package terminal

import (
	"time"

	"github.com/rs/zerolog/log"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

type WindowsTerminal struct {
	m *mgr.Mgr
}

type LinuxTerminal struct {
}

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

func (w *LinuxTerminal) ShutdownAgent(agentName string) error {
	return nil
}

func (w *LinuxTerminal) StartAgent(agentName string) error {
	return nil
}
