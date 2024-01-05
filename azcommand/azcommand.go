package azcommand

import (
	"az-tools/model"
	"github.com/TwiN/go-color"
	"os/exec"
	"strings"
)

type AzCommand struct {
	Text string
}

func (command AzCommand) Exec(args ...string) error {
	println(color.OverGreen(command.Text))
	cmd := exec.Command("az", args...)
	println(color.InGreen(cmd))
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

func execAzAks(resourceGroup string, aks string) error {
	cmd := AzCommand{Text: "Set AKS to " + aks + " in " + resourceGroup}
	return cmd.Exec("aks", "get-credentials", "-n", aks, "-g", resourceGroup)
}

func ExecCommands(reference model.NodeReference) error {
	err := execAzSubscription(reference.Subscription)
	if err != nil {
		return err
	}
	err = execAzAks(reference.ResourceGroup, reference.Aks)
	if err != nil {
		return err
	}
	return nil
}
