package main

import (
	"fmt"
	"os"

	"github.com/Fryuni/world-anvil-cli/pkg/config"
)

func main() {
	cfg := config.Load()

	if len(os.Args) < 2 {
		fmt.Println("World Anvil CLI")
		fmt.Println("Usage: world-anvil-cli <command>")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "version":
		fmt.Println("world-anvil-cli v0.1.0")
	case "config":
		fmt.Printf("Configuration loaded: %+v\n", cfg)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
