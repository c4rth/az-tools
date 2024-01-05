package main

import (
	"az-tools/azcommand"
	"az-tools/model"
	"github.com/rivo/tview"
	"os"
)

func main() {

	reference := model.NodeReference{}
	selected := false

	//subscriptions, err := model.ReadSubscriptions("az-tools.yaml")
	config, err := model.ReadConfig()
	if err != nil {
		panic(err)
	}

	rootDir := "[:bold]subscriptions"
	root := tview.NewTreeNode(rootDir).SetSelectable(false)
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)
	for _, subscription := range config.Subscriptions {
		subscriptionNode := tview.NewTreeNode("[:bold]" + subscription.Name).
			SetSelectable(false)
		for _, resourceGroup := range subscription.ResourceGroups {
			for _, aks := range resourceGroup.Aks {
				aksNode := tview.NewTreeNode(resourceGroup.Name + " / " + aks.Name).
					SetSelectable(true).
					SetReference(model.NodeReference{Subscription: subscription.Name, ResourceGroup: resourceGroup.Name, Aks: aks.Name})
				subscriptionNode.AddChild(aksNode)
			}
		}

		root.AddChild(subscriptionNode)
	}

	app := tview.NewApplication().SetRoot(tree, true)

	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference = node.GetReference().(model.NodeReference)
		selected = true
		app.Stop()
	}).
		SetBorder(true).
		SetBorderPadding(0, 0, 1, 1).
		SetTitle("Az tools").
		SetTitleAlign(tview.AlignLeft)

	if err := app.Run(); err != nil {
		panic(err)
	}

	if selected {
		err = azcommand.ExecCommands(reference)
		if err != nil {
			os.Exit(1)
		}
		os.Exit(0)
	}
}
