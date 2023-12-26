package azcommand

import (
	"az-tools/model"
	"github.com/TwiN/go-color"
	"os"
	"os/exec"
	"strings"
)

type AzCommand struct {
	Text string
}

func (c AzCommand) Exec(args ...string) error {
	println(color.OverGreen(c.Text))
	println(color.OverGreen(args))
	cmd := exec.Command("az", args...)
	var stdOut strings.Builder
	var stdErr strings.Builder
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		println(color.InRed(stdErr.String()))
	} else {
		println(stdOut.String())
	}
	return err
}

func execAzSubscription(subscription string) error {
	cmd := AzCommand{Text: "Set subscription to " + subscription}
	return cmd.Exec("account", "set", "--subscription", subscription)
}

func execAzAks(aks model.Aks) error {
	cmd := AzCommand{Text: "Set AKS to " + aks.Name + " in " + aks.ResourceGroup}
	return cmd.Exec("aks", "get-credentials", "-n", aks.Name, "-g", aks.ResourceGroup)
}

func ExecCommands(reference model.NodeReference) error {
	err := execAzSubscription(reference.Subscription)
	if err != nil {
		os.Exit(1)
	}
	err = execAzAks(reference.Aks)
	if err != nil {
		os.Exit(1)
	}
	return nil
}
