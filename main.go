package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/yohany99/gator/internal/config"
	"github.com/yohany99/gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
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
	currCommands.register("agg", handlerAgg)
	currCommands.register("addfeed", handlerAddFeed)
	currCommands.register("feeds", handlerFeeds)
	currCommands.register("follow", handlerFollow)
	currCommands.register("following", handlerFollowing)
	if len(os.Args) < 2 {
		log.Fatal("usage: cli <command> [args...]")
	}
	currCommand := command{
		name: os.Args[1],
		args: os.Args[2:],
	}
	err = currCommands.run(&currState, currCommand)
	if err != nil {
		log.Fatal(err)
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
