package ui

import (
	"strings"
	"testing"

	"k8switch/internal/types"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func TestTreeRespectsTerminalDefaults(t *testing.T) {
	screen := newTestScreen(t)
	tree := buildTree(types.Config{
		Ranchers: []types.Rancher{{Name: "[red]rancher"}},
		Subscriptions: []types.Subscription{{
			Name: "subscription",
			ResourceGroups: []types.ResourceGroup{{
				Name: "group",
				Aks:  []types.Aks{{Name: "cluster"}},
			}},
		}},
	})
	tree.SetRect(0, 0, 100, 25)
	selected := tree.GetRoot().GetChildren()[0].GetChildren()[0]
	tree.SetCurrentNode(selected)
	tree.Draw(screen)

	text, reverse, headings := inspectScreen(t, screen, true)
	for _, want := range []string{"K8S", "Rancher", "[red]rancher", "Azure", "subscription", "group / cluster"} {
		if !strings.Contains(text, want) {
			t.Errorf("rendered tree does not contain %q", want)
		}
	}
	if !strings.Contains(reverse, "rancher") {
		t.Errorf("selected node is not rendered in reverse video: %q", reverse)
	}
	for _, want := range []string{"K8S", "Rancher", "Azure", "subscription"} {
		if !strings.Contains(headings, want) {
			t.Errorf("heading %q does not use white on green: %q", want, headings)
		}
	}
	if strings.Contains(headings, "rancher") || strings.Contains(headings, "group / cluster") {
		t.Errorf("selectable items should use terminal defaults: %q", headings)
	}
}

func TestDummyModalRespectsTerminalDefaults(t *testing.T) {
	screen := newTestScreen(t)
	pages := tview.NewPages()
	modal := buildDummyConfigModal(pages)
	tview.NewApplication().SetFocus(modal)
	modal.Draw(screen)

	text, reverse, _ := inspectScreen(t, screen, false)
	if !strings.Contains(text, "dummy configuration") {
		t.Error("dummy configuration message was not rendered")
	}
	if !strings.Contains(reverse, "OK") {
		t.Errorf("focused button is not rendered in reverse video: %q", reverse)
	}
}

func newTestScreen(t *testing.T) tcell.SimulationScreen {
	t.Helper()
	screen := tcell.NewSimulationScreen("UTF-8")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(screen.Fini)
	screen.SetSize(100, 25)
	return screen
}

func inspectScreen(t *testing.T, screen tcell.Screen, allowHeadings bool) (string, string, string) {
	t.Helper()
	var text, reverse, headings strings.Builder
	width, height := screen.Size()
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			char, _, style, _ := screen.GetContent(x, y)
			fg, bg, attr := style.Decompose()
			if allowHeadings && fg == tcell.ColorWhite && bg == tcell.ColorGreen {
				headings.WriteRune(char)
			} else if fg != tcell.ColorDefault || bg != tcell.ColorDefault {
				t.Errorf("cell (%d,%d) %q uses foreground %v and background %v", x, y, char, fg, bg)
			}
			text.WriteRune(char)
			if attr&tcell.AttrReverse != 0 {
				reverse.WriteRune(char)
			}
		}
		text.WriteByte('\n')
	}
	return text.String(), reverse.String(), headings.String()
}
