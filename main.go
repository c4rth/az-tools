package main

import (
	"az-tools/azcommand"
	"az-tools/model"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"os"
)

func main() {

	reference := model.NodeReference{}
	selected := false

	//subscriptions, err := model.ReadSubscriptions("config.yaml")
	config, isDummy, err := model.ReadConfig()
	if err != nil {
		panic(err)
	}

	rootDir := "[white:green:b]subscriptions"
	root := tview.NewTreeNode(rootDir).SetSelectable(false)
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)
	for _, subscription := range config.Subscriptions {
		subscriptionNode := tview.NewTreeNode("[white:green:b]" + subscription.Name).SetSelectable(false).SetColor(tview.Styles.SecondaryTextColor)
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

	app := tview.NewApplication()

	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference = node.GetReference().(model.NodeReference)
		selected = true
		app.Stop()
	}).
		SetBackgroundColor(tcell.ColorBlack).
		SetBorder(true).
		SetBorderPadding(0, 0, 1, 1).
		SetTitle("Az tools").
		SetTitleAlign(tview.AlignLeft)

	pages := tview.NewPages().
		AddPage("tree", tree, true, true)

	if isDummy {
		modal := tview.NewModal().
			SetText("A dummy configuration is created at: \n $HOME/.az-tools/config.yaml").
			AddButtons([]string{"OK"}).
			SetBackgroundColor(tcell.ColorGreen).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				pages.SwitchToPage("tree")
			})
		pages.AddPage("modal", modal, true, true)
	}

	if err := app.SetRoot(pages, true).Run(); err != nil {
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
