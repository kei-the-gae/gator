package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/kei-the-gae/gator/internal/config"
	"github.com/kei-the-gae/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	s := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)

	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	cmdName := args[0]
	cmdArgs := args[1:]

	if err := cmds.run(s, command{Name: cmdName, Args: cmdArgs}); err != nil {
		log.Fatal(err)
	}

}
