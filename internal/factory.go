package internal

type (
	OsType int

	TerminalFactory struct {
		osType OsType
	}
)

const (
	Linux OsType = iota
	Windows
)

func NewTerminalFactory(osType OsType) TerminalFactory {
	return TerminalFactory{
		osType: osType,
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
