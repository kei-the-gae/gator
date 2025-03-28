package main

import (
	"log"
	"os"

	"github.com/kei-the-gae/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	s := &state{
		cfg: &cfg,
	}
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	args := os.Args[1:]
	if len(args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	cmdName := args[0]
	cmdArgs := args[1:]

	if err := cmds.run(s, command{Name: cmdName, Args: cmdArgs}); err != nil {
		log.Fatal(err)
	}

}
