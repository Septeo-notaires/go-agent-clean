package internal

type ITerminal interface {
	CleanAgent(path string) error
	ShutdownAgent(agentName string) error
	StartAgent(agentName string) error
}

type WindowsTerminal struct {
}

type LinuxTerminal struct {
}

// Windows Terminal

func (w *WindowsTerminal) CleanAgent(path string) error {
	return nil
}

func (w *WindowsTerminal) ShutdownAgent(agentName string) error {
	return nil
}

func (w *WindowsTerminal) StartAgent(agentName string) error {
	return nil
}

// Linux Terminal

func (w *LinuxTerminal) CleanAgent(path string) error {
	return nil
}

func (w *LinuxTerminal) ShutdownAgent(agentName string) error {
	return nil
}

func (w *LinuxTerminal) StartAgent(agentName string) error {
	return nil
}
