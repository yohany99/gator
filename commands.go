package main

import (
	"errors"
	"fmt"
	"time"

	"context"

	"github.com/google/uuid"
	"github.com/yohany99/gator/internal/config"
	"github.com/yohany99/gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

type command struct {
	name string
	args []string
}

type commands struct {
	commandMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if function, ok := c.commandMap[cmd.name]; ok {
		err := function(s, cmd)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("command not found")
	}
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandMap[name] = f
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("username is missing")
	}
	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}
	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Println("user has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("name is required")
	}
	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}
	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return err
	}
	s.cfg.CurrentUserName = userParams.Name
	fmt.Printf("user created: %+v\n", user)
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetDB(context.Background())
	if err != nil {
		return err
	}
	return nil
}
