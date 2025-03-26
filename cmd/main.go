package main

import (
	"github.com/go-agent-clean/internal"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func main() {
	log.Print("Clean Agent Hebdo")

	factory := internal.NewTerminalFactory("", internal.Windows)
	term := factory.NewTerminal()
	if err := term.ShutdownAgent("vstsagent.Septeo-GenApi.ComptaBuildAgents.Agent1PcDamienVv"); err != nil {
		log.Err(err).Stack().Msg("shutdown")
	}

	if err := term.StartAgent("vstsagent.Septeo-GenApi.ComptaBuildAgents.Agent1PcDamienVv"); err != nil {
		log.Err(err).Stack().Msg("start")
	}
}
