package main

import (
	"bufio"
	"os"
	"sync"

	"github.com/go-agent-clean/internal"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var conf internal.Config

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if file, err := os.OpenFile("./config.toml", os.O_RDONLY, os.ModeAppend); err != nil {
		panic(err)
	} else {
		defer file.Close()
		reader := bufio.NewReader(file)
		blob, _ := reader.ReadString(0)
		conf = internal.DecodeFile(blob)
	}
}

func main() {
	factory := internal.NewTerminalFactory(internal.Windows)
	term := factory.NewTerminal()

	wg := sync.WaitGroup{}

	for _, agent := range conf.Agents {
		wg.Add(1)
		go func() {
			defer wg.Done()

			if err := term.ShutdownAgent(agent.Service); err != nil {
				log.Err(err).Stack().Msg("shutdown " + agent.Name)
				return
			}

			if err := term.CleanAgent(agent.Path); err != nil {
				log.Err(err).Stack().Msg("clear " + agent.Name)
				return
			}

			if err := term.StartAgent(agent.Service); err != nil {
				log.Err(err).Stack().Msg("start " + agent.Name)
				return
			}
		}()
	}

	wg.Wait()
}
