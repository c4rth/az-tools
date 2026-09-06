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
	root := tview.NewTreeNode("[white:green:b]K8S").SetSelectable(false)
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)
	tree.SetBackgroundColor(tcell.ColorBlack).
		SetBorder(true).
		SetBorderPadding(0, 0, 1, 1).
		SetTitle("K8switch").
		SetTitleAlign(tview.AlignLeft)
	if len(config.Ranchers) != 0 {
		ranchersNode := tview.NewTreeNode("[white:green:b]Rancher").SetSelectable(false)
		for _, rancher := range config.Ranchers {
			rancherNode := tview.NewTreeNode(rancher.Name).
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
		azureNode := tview.NewTreeNode("[white:green:b]Azure").SetSelectable(false)
		for _, subscription := range config.Subscriptions {
			subscriptionNode := tview.NewTreeNode("[white:green:b]" + subscription.Name).
				SetSelectable(false).
				SetColor(tview.Styles.SecondaryTextColor)
			for _, resourceGroup := range subscription.ResourceGroups {
				for _, aks := range resourceGroup.Aks {
					aksNode := tview.NewTreeNode(resourceGroup.Name + " / " + aks.Name).
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

	return tree
}

func buildDummyConfigModal(pages *tview.Pages) *tview.Modal {
	loc := fmt.Sprintf("A dummy configuration is created at: \n $HOME%s%s", types.ConfigFileDir, types.ConfigFilename)
	return tview.NewModal().
		SetText(loc).
		AddButtons([]string{"OK"}).
		SetBackgroundColor(tcell.ColorGreen).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.SwitchToPage("tree")
		})
}
