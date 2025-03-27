package main

import (
	"bufio"
	"os"
	"runtime"
	"sync"

	"github.com/go-agent-clean/internal/terminal"
	"github.com/go-agent-clean/internal/toml"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var conf toml.Config

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if file, err := os.OpenFile("./config.toml", os.O_RDONLY, os.ModeAppend); err != nil {
		panic(err)
	} else {
		defer file.Close()
		reader := bufio.NewReader(file)
		blob, _ := reader.ReadString(0)
		conf = toml.DecodeFile(blob)
	}
}

func main() {
	var operatingSystem terminal.OsType

	if runtime.GOOS == "windows" {
		operatingSystem = terminal.Windows
	} else if runtime.GOOS == "linux" {
		operatingSystem = terminal.Linux
	} else {
		log.Error().Msg("Operating System not compatible")
		os.Exit(-1)
	}

	factory := terminal.NewTerminalFactory(operatingSystem)
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
