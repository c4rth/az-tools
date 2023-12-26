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

	subscriptions, err := model.ReadSubscriptions("az-tools.yaml")
	if err != nil {
		panic(err)
	}

	rootDir := "subscriptions"
	root := tview.NewTreeNode(rootDir).SetSelectable(false)
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)
	for _, subscription := range subscriptions {
		subNode := tview.NewTreeNode(subscription.Subscription).
			SetSelectable(false)
		for _, aks := range subscription.Aks {
			aksNode := tview.NewTreeNode(aks.ResourceGroup + "/" + aks.Name).
				SetSelectable(true).
				SetReference(model.NodeReference{Subscription: subscription.Subscription, Aks: aks})
			subNode.AddChild(aksNode)
		}
		root.AddChild(subNode)
	}

	app := tview.NewApplication().SetRoot(tree, true)

	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference = node.GetReference().(model.NodeReference)
		selected = true
		app.Stop()
	}).
		SetBorder(true).
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
