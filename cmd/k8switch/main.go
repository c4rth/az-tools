package main

import (
	"k8switch/internal/cli"
	"k8switch/internal/config"
	"k8switch/internal/ui"

	"fmt"
	"os"
)

func run() error {
	config, isDummy, err := config.ReadConfig()
	if err != nil {
		return err
	}

	selectedNode, selected, err := ui.RunTUI(config, isDummy)
	if err != nil {
		return err
	}
	if !selected {
		return nil
	}

	return cli.ExecCommands(selectedNode)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
