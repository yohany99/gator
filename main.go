package main

import (
	"fmt"
	"os"

	"github.com/yohany99/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	currState := state{
		cfg: &cfg,
	}
	currCommands := commands{
		commandMap: make(map[string]func(*state, command) error),
	}
	currCommands.register("login", handlerLogin)
	if len(os.Args) < 2 {
		fmt.Println("command is required")
		os.Exit(0)
	}
	currCommand := command{
		name: os.Args[1],
		args: os.Args[2:],
	}
	err = currCommands.run(&currState, currCommand)
	if err != nil {
		fmt.Println(err)
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
