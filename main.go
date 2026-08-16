package main

import (
	"az-tools/azcommand"
	"az-tools/model"
	"az-tools/ui"

	"fmt"
	"os"
)

func run() error {
	config, isDummy, err := model.ReadConfig()
	if err != nil {
		return err
	}

	selectedReference, selected, err := ui.RunTUI(config, isDummy)
	if err != nil {
		return err
	}
	if !selected {
		return nil
	}

	return azcommand.ExecCommands(selectedReference)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
