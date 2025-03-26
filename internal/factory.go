package internal

type OsType int

const (
	Linux OsType = iota
	Windows
)

type TerminalFactory struct {
	osType          OsType
	interpreterName string
}

func NewTerminalFactory(interpreterName string, osType OsType) TerminalFactory {
	return TerminalFactory{
		osType:          osType,
		interpreterName: interpreterName,
	}
}

func (t *TerminalFactory) NewTerminal() ITerminal {
	switch t.osType {
	case Windows:
		term := NewWindowTerminal()
		return &term
	case Linux:
		return &LinuxTerminal{}
	default:
		return nil
	}
}
