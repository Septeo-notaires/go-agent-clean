//go:build linux
// +build linux

package terminal

type WindowsTerminal struct {
}

type LinuxTerminal struct {
}

func NewWindowTerminal() WindowsTerminal {
	return WindowsTerminal{}
}

func (w *WindowsTerminal) ShutdownAgent(agentName string) error {
	return nil
}

func (w *WindowsTerminal) StartAgent(agentName string) error {
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
