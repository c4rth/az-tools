package cli

import (
	"github.com/TwiN/go-color"
	"k8switch/internal/types"
	"os/exec"
	"strings"
)

type ExecCommand struct {
	Text string
}

func (command ExecCommand) Exec(args ...string) error {
	println(color.OverGreen(command.Text))
	cmd := exec.Command(args[0], args[1:]...)
	println(color.InGreen(strings.Join(cmd.Args, " ")))
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

func execRancherSwitch(node types.NodeReference) error {
	cmd := ExecCommand{Text: "Set Rancher context to " + node.Cluster}
	return cmd.Exec("kubectl", "config", "use-context", node.Cluster)
}

func execAzSwitch(node types.NodeReference) error {
	cmdSubsc := ExecCommand{Text: "Set active subscription to " + node.Subscription}
	err := cmdSubsc.Exec("az", "account", "set", "--subscription", node.Subscription)
	if err != nil {
		return err
	}
	cmdAks := ExecCommand{Text: "Set AKS to " + node.Cluster + " in " + node.ResourceGroup}
	return cmdAks.Exec("az", "aks", "get-credentials", "-n", node.Cluster, "-g", node.ResourceGroup, "--overwrite")
}

func ExecCommands(node types.NodeReference) error {
	switch node.Kind {
	case types.KindRancher:
		return execRancherSwitch(node)
	case types.KindAks:
		return execAzSwitch(node)
	}
	return nil
}
