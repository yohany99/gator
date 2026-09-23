package main

import (
	"context"
	"database/sql"
	"fmt"
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
	currCommands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	currCommands.register("feeds", handlerFeeds)
	currCommands.register("follow", middlewareLoggedIn(handlerFollow))
	currCommands.register("following", middlewareLoggedIn(handlerFollowing))
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
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := getLoggedInUser(s)
		if err != nil {
			return err
		}
		return handler(s, cmd, *user)
	}
}

func getLoggedInUser(s *state) (*database.User, error) {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return nil, fmt.Errorf("couldn't get user: %w", err)
	}
	return &user, nil
}
