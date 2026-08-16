package ui

import (
	"az-tools/model"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)


func RunTUI(config model.Config, isDummy bool) (model.NodeReference, bool, error) {
	app := tview.NewApplication()
	selectedReference := model.NodeReference{}
	selected := false

	tree := buildTree(config)
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		selectedReference = node.GetReference().(model.NodeReference)
		selected = true
		app.Stop()
	})

	pages := tview.NewPages().AddPage("tree", tree, true, true)
	if isDummy {
		pages.AddPage("modal", buildDummyConfigModal(pages), true, true)
	}

	if err := app.SetRoot(pages, true).Run(); err != nil {
		return model.NodeReference{}, false, err
	}

	return selectedReference, selected, nil
}

func buildTree(config model.Config) *tview.TreeView {
	root := tview.NewTreeNode("[white:green:b]subscriptions").SetSelectable(false)
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)
	tree.SetBackgroundColor(tcell.ColorBlack).
		SetBorder(true).
		SetBorderPadding(0, 0, 1, 1).
		SetTitle("Az tools").
		SetTitleAlign(tview.AlignLeft)

	for _, subscription := range config.Subscriptions {
		subscriptionNode := tview.NewTreeNode("[white:green:b]" + subscription.Name).
			SetSelectable(false).
			SetColor(tview.Styles.SecondaryTextColor)
		for _, resourceGroup := range subscription.ResourceGroups {
			for _, aks := range resourceGroup.Aks {
				aksNode := tview.NewTreeNode(resourceGroup.Name + " / " + aks.Name).
					SetSelectable(true).
					SetReference(model.NodeReference{
						Subscription:  subscription.Name,
						ResourceGroup: resourceGroup.Name,
						Aks:           aks.Name,
					})
				subscriptionNode.AddChild(aksNode)
			}
		}
		root.AddChild(subscriptionNode)
	}

	return tree
}

func buildDummyConfigModal(pages *tview.Pages) *tview.Modal {
	return tview.NewModal().
		SetText("A dummy configuration is created at: \n $HOME/.az-tools/config.yaml").
		AddButtons([]string{"OK"}).
		SetBackgroundColor(tcell.ColorGreen).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.SwitchToPage("tree")
		})
}
