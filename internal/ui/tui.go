package ui

import (
	"fmt"
	"k8switch/internal/types"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func RunTUI(config types.Config, isDummy bool) (types.NodeReference, bool, error) {
	app := tview.NewApplication()
	var selected types.NodeReference
	selectedFlag := false

	tree := buildTree(config)
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		ref := node.GetReference().(types.NodeReference)
		selected = ref
		selectedFlag = true
		app.Stop()
	})

	pages := tview.NewPages().AddPage("tree", tree, true, true)
	if isDummy {
		pages.AddPage("modal", buildDummyConfigModal(pages), true, true)
	}

	if err := app.SetRoot(pages, true).Run(); err != nil {
		return types.NodeReference{}, false, err
	}

	return selected, selectedFlag, nil
}

func buildTree(config types.Config) *tview.TreeView {
	heading := func(text string) *tview.TreeNode {
		style := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorGreen).Bold(true)
		return tview.NewTreeNode(text).SetSelectable(false).SetTextStyle(style)
	}
	root := heading("K8S")
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)
	tree.SetGraphicsColor(tcell.ColorDefault)
	tree.SetBackgroundColor(tcell.ColorDefault).
		SetBorderStyle(tcell.StyleDefault).
		SetTitleColor(tcell.ColorDefault).
		SetBorder(true).
		SetBorderPadding(0, 0, 1, 1).
		SetTitle("K8switch").
		SetTitleAlign(tview.AlignLeft)
	if len(config.Ranchers) != 0 {
		ranchersNode := heading("Rancher")
		for _, rancher := range config.Ranchers {
			rancherNode := tview.NewTreeNode(tview.Escape(rancher.Name)).
				SetTextStyle(tcell.StyleDefault).
				SetReference(types.NodeReference{
					Subscription:  "",
					ResourceGroup: "",
					Cluster:       rancher.Name,
					Kind:          types.KindRancher,
				})
			ranchersNode.AddChild(rancherNode)
		}
		root.AddChild(ranchersNode)
	}
	if len(config.Subscriptions) != 0 {
		azureNode := heading("Azure")
		for _, subscription := range config.Subscriptions {
			subscriptionNode := heading(tview.Escape(subscription.Name))
			for _, resourceGroup := range subscription.ResourceGroups {
				for _, aks := range resourceGroup.Aks {
					aksNode := tview.NewTreeNode(tview.Escape(resourceGroup.Name + " / " + aks.Name)).
						SetTextStyle(tcell.StyleDefault).
						SetSelectable(true).
						SetReference(types.NodeReference{
							Subscription:  subscription.Name,
							ResourceGroup: resourceGroup.Name,
							Cluster:       aks.Name,
							Kind:          types.KindAks,
						})
					subscriptionNode.AddChild(aksNode)
				}
			}
			azureNode.AddChild(subscriptionNode)
		}
		root.AddChild(azureNode)
	}

	root.Walk(func(node, parent *tview.TreeNode) bool {
		node.SetSelectedTextStyle(node.GetTextStyle().Reverse(true))
		return true
	})

	return tree
}

func buildDummyConfigModal(pages *tview.Pages) *tview.Modal {
	loc := fmt.Sprintf("A dummy configuration is created at: \n $HOME%s%s", types.ConfigFileDir, types.ConfigFilename)
	modal := tview.NewModal().
		SetText(loc).
		SetTextColor(tcell.ColorDefault).
		SetBackgroundColor(tcell.ColorDefault).
		SetButtonStyle(tcell.StyleDefault).
		SetButtonActivatedStyle(tcell.StyleDefault.Reverse(true)).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.SwitchToPage("tree")
		})
	modal.Box.SetBackgroundColor(tcell.ColorDefault).SetBorderStyle(tcell.StyleDefault)
	return modal
}
