// Copyright (C) 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ui

import (
	"fmt"
	"log"
	"strings"

	"github.com/jroimartin/gocui"

	"github.com/nlamirault/helmsman/k8s"
	"github.com/nlamirault/helmsman/version"
)

var (
	cmdBuffer = []string{}
	cmdIdx    = 0
)

func registerKeybindings(g *gocui.Gui, k8sclient *k8s.Client) error {
	log.Printf("[DEBUG] Register keybindings")
	// Set quit
	if err := g.SetKeybinding("", gocui.KeyCtrlQ, gocui.ModNone, quitHandler); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlH, gocui.ModNone, helpHandler); err != nil {
		return err
	}

	// Submit a line
	// if err := g.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, inputLineHandler); err != nil {
	// 	return err
	// }

	// Tab
	if err := g.SetKeybinding(sideViewName, gocui.KeyTab, gocui.ModNone, nextViewHandler); err != nil {
		return err
	}
	if err := g.SetKeybinding(mainViewName, gocui.KeyTab, gocui.ModNone, nextViewHandler); err != nil {
		return err
	}

	// Cursors
	if err := g.SetKeybinding(sideViewName, gocui.KeyArrowDown, gocui.ModNone, cursorDownHandler); err != nil {
		return err
	}
	if err := g.SetKeybinding(sideViewName, gocui.KeyArrowUp, gocui.ModNone, cursorUpHandler); err != nil {
		return err
	}
	if err := g.SetKeybinding(sideViewName, gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return kubernetesDispatcher(g, v, k8sclient)
	}); err != nil {
		return err
	}
	if err := g.SetKeybinding(mainViewName, gocui.KeyArrowDown, gocui.ModNone, cursorDownHandler); err != nil {
		return err
	}
	if err := g.SetKeybinding(mainViewName, gocui.KeyArrowUp, gocui.ModNone, cursorUpHandler); err != nil {
		return err
	}

	// Arrow up/down scrolls cmd history
	if err := g.SetKeybinding("input", gocui.KeyArrowUp, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			scrollHistory(v, -1)
			return nil
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("input", gocui.KeyArrowDown, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			scrollHistory(v, 1)
			return nil
		}); err != nil {
		return err
	}
	return nil
}

func quitHandler(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func helpHandler(g *gocui.Gui, v *gocui.View) error {
	mainView, err := g.View(mainViewName)
	if err != nil {
		return err
	}
	mainView.Clear()
	// mainView.Editable = true
	mainView.Highlight = true
	mainView.Title = "Help"
	mainView.SetCursor(0, 0)
	mainView.SetOrigin(0, 0)
	printHelp(mainView)
	return nil
}

func nextViewHandler(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() == sideViewName {
		return g.SetCurrentView(mainViewName)
		// } else if v.Name() == "main" {
		// 	return g.SetCurrentView("input")
		// } else if v.Name() == "input" {
		// 	return g.SetCurrentView("side")
	}
	return g.SetCurrentView(sideViewName)
}

func cursorDownHandler(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func cursorUpHandler(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func kubernetesDispatcher(g *gocui.Gui, v *gocui.View, client *k8s.Client) error {
	// log.Printf("[DEBUG] Select cursor")
	if v != nil {
		_, cy := v.Cursor()
		l, err := v.Line(cy)
		if err != nil {
			l = ""
		}

		view, err := g.View(mainViewName)
		if err != nil {
			return err
		}
		view.Clear()
		// view.Editable = true
		view.Highlight = true
		view.Title = l
		view.SetCursor(0, 0)
		view.SetOrigin(0, 0)
		// fmt.Fprintf(mainView, "----> %s\n", l)
		switch l {
		case k8sNamespaces:
			printK8SNamespaces(view, client)
		case k8sNodes:
			printK8SNodes(view, client)
		case k8sPersistentVolumes:
			printK8SPersistentVolumes(view, client)
		case k8sDeployments:
			printK8SDeployments(view, client)
		case k8sReplicaSets:
			printK8SReplicaSets(view, client)
		case k8sReplicationControllers:
			printK8SReplicationControllers(view, client)
		case k8sDaemonSets:
			printK8SDaemonSets(view, client)
		case k8sJobs:
			printK8SJobs(view, client)
		case k8sPods:
			printK8SPods(view, client)
		case k8sServices:
			printK8SServices(view, client)
		case k8sIngress:
			printK8SIngresses(view, client)
		case k8sPersistentVolumeClaims:
			printK8SPersistentVolumeClaims(view, client)
		case k8sSecrets:
			printK8SSecrets(view, client)
		case k8sConfigMaps:
			printK8SConfigMaps(view, client)
		}
		return g.SetCurrentView(mainViewName)
	}
	return nil
}

func inputLineHandler(g *gocui.Gui, v *gocui.View) error {
	line := strings.TrimSpace(v.Buffer())
	if line != "" {
		cmdBuffer = append(cmdBuffer, line)
		cmdIdx = len(cmdBuffer)
	} else {
		// it's an empty line, return nil and be done with it
		return nil
	}

	// Get a main view obvject to print output to
	ov, _ := g.View(mainViewName)

	// Parse if it is an internal command or otherwise
	if len(line) > 0 && string(line[0]) == "/" {
		// parse internal command
		s, args := strings.Split(line, " "), []string{}
		cmd := s[0]
		if len(s) > 1 {
			args = append(args, s[1:]...)
		} else {
			args = append(args, "")
		}

		switch cmd {
		case "/quit":
			return gocui.ErrQuit
		case "/clear":
			ov.Clear()
			// case "/printInputBuffer":
			// 	fmt.Fprintln(ov, "INPUT BUFFER:")
			// 	fmt.Fprintln(ov, cmdBuffer)
			// case "/clearInputBuffer":
			// 	cmdBuffer = nil
		}
	} else {
		// print to output and do whatever with it
		fmt.Fprintln(ov, "cmd:", line)
	}

	// Clear the input buffer now that the line has been dealt with
	v.Clear()
	return nil
}

func scrollHistory(v *gocui.View, dy int) {
	if v != nil {
		if i := cmdIdx + dy; i >= 0 && i < len(cmdBuffer) {
			cmdIdx = i
			v.Clear()
			fmt.Fprintf(v, "%v", cmdBuffer[cmdIdx])
			v.SetOrigin(0, 0)
		}
	}
}

func printHelp(v *gocui.View) {
	v.Clear()
	v.Highlight = true
	v.Title = "Help"
	v.SetCursor(0, 0)
	v.SetOrigin(0, 0)
	fmt.Fprintf(v, "\n\nWelcome to Helmsman v%s\n", version.Version)
	fmt.Fprintf(v, "Keybindings:\n\n")
	fmt.Fprintf(v, " TAB      : Next view\n")
	fmt.Fprintf(v, " ← ↑ → ↓  : Move cursor\n")
	fmt.Fprintf(v, " ^q       : Quit\n")
	fmt.Fprintf(v, " ^h       : Show help message\n")
}
