package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/yohany99/gator/internal/config"
	"github.com/yohany99/gator/internal/database"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	dbQueries := database.New(db)
	currState := state{
		cfg: &cfg,
		db:  dbQueries,
	}
	currCommands := commands{
		commandMap: make(map[string]func(*state, command) error),
	}
	currCommands.register("login", handlerLogin)
	currCommands.register("register", handlerRegister)
	currCommands.register("reset", handlerReset)
	currCommands.register("users", handlerUsers)
	if len(os.Args) < 2 {
		fmt.Println("command is required")
		os.Exit(1)
	}
	currCommand := command{
		name: os.Args[1],
		args: os.Args[2:],
	}
	err = currCommands.run(&currState, currCommand)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// cfg, err := config.Read()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// cfg.SetUser("Yohan")
	// cfg, err = config.Read()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(cfg)
}
