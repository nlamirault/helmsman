// Copyright (C) 2016, 2017 Nicolas Lamirault <nicolas.lamirault@gmail.com>

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
	// "log"
	"strings"

	"github.com/golang/glog"
	"github.com/jroimartin/gocui"

	"github.com/nlamirault/helmsman/k8s"
	"github.com/nlamirault/helmsman/version"
)

var (
	cmdBuffer = []string{}
	cmdIdx    = 0
)

func registerKeybindings(g *gocui.Gui, k8sclient *k8s.Client) error {
	glog.V(2).Info("Register keybindings")
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
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, nextViewHandler); err != nil {
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
		return kubernetesMenuDispatcher(g, v, k8sclient)
	}); err != nil {
		return err
	}
	if err := g.SetKeybinding(mainViewName, gocui.KeyArrowDown, gocui.ModNone, cursorDownHandler); err != nil {
		return err
	}
	if err := g.SetKeybinding(mainViewName, gocui.KeyArrowUp, gocui.ModNone, cursorUpHandler); err != nil {
		return err
	}

	// Details
	if err := g.SetKeybinding(mainViewName, gocui.KeyCtrlD, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return kubernetesDescriptionDispatcher(g, v, k8sclient)
	}); err != nil {
		return err
	}
	if err := g.SetKeybinding(detailViewName, gocui.KeyCtrlW, gocui.ModNone, closeDetailsViewHandler); err != nil {
		return err
	}

	// Arrow up/down scrolls cmd history
	// if err := g.SetKeybinding("input", gocui.KeyArrowUp, gocui.ModNone,
	// 	func(g *gocui.Gui, v *gocui.View) error {
	// 		scrollHistory(v, -1)
	// 		return nil
	// 	}); err != nil {
	// 	return err
	// }
	// if err := g.SetKeybinding("input", gocui.KeyArrowDown, gocui.ModNone,
	// 	func(g *gocui.Gui, v *gocui.View) error {
	// 		scrollHistory(v, 1)
	// 		return nil
	// 	}); err != nil {
	// 	return err
	// }
	return nil
}

func quitHandler(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func closeDetailsViewHandler(g *gocui.Gui, v *gocui.View) error {
	// if _, err := setCurrentViewOnTop(g, mainViewName); err != nil {
	// 	return err
	// }
	// return nil
	glog.V(2).Infof("View: %s", v.Name())
	return closeView(g, v)
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
	clearView(mainView)
	printHelp(mainView)
	return nil
}

func nextViewHandler(g *gocui.Gui, v *gocui.View) error {
	if v == nil {
		if _, err := g.SetCurrentView(mainViewName); err != nil {
			return err
		}
		return nil
	}
	switch v.Name() {
	case sideViewName:
		if _, err := g.SetCurrentView(mainViewName); err != nil {
			return err
		}
		return nil
	case mainViewName:
		if _, err := g.SetCurrentView(sideViewName); err != nil {
			return err
		}
		return nil
	case detailViewName:
		if _, err := g.SetCurrentView(sideViewName); err != nil {
			return err
		}
		return nil
	}
	return nil
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

func kubernetesMenuDispatcher(g *gocui.Gui, v *gocui.View, client *k8s.Client) error {
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
		clearView(view)
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
		// case k8sJobs:
		// 	printK8SJobs(view, client)
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
		// glog.V(2).Infof("Active view: =========> %s\n", view.Name())
		if _, err := g.SetCurrentView(mainViewName); err != nil {
			return err
		}
	}
	return nil
}

func kubernetesDescriptionDispatcher(g *gocui.Gui, view *gocui.View, client *k8s.Client) error {
	if view != nil {
		_, cy := view.Cursor()
		line, err := view.Line(cy)
		if err != nil {
			fmt.Fprintf(view, "\033[31;01mInput error:\n%s\033[0m", err.Error())
			return nil
		}
		if len(line) == 0 {
			return nil
		}
		// view, err := g.View(detailViewName)
		// if err != nil {
		// 	return err
		// }
		view.Clear()
		// view.Editable = true
		view.Highlight = true
		// view.Title = "Details"
		clearView(view)
		data := strings.Split(line, " ")
		if len(data) < 2 {
			fmt.Fprintf(view, "\033[31;01mCan't extract Kubernetes information: \n%s\033[0m", line)
		}
		// glog.Infof("----> %s ==========> %s ========> %s\n", v.Name(), v.Title, line)
		switch view.Title {
		case k8sNamespaces:
			fmt.Fprintf(view, "1111111&")
		case k8sNodes:
			printK8SNodeDescription(view, data[1], client)
		case k8sPersistentVolumes:
			fmt.Fprintf(view, "3333")
		case k8sDeployments:
			fmt.Fprintf(view, "44444444444")
		case k8sReplicaSets:
			fmt.Fprintf(view, "55555555")
		case k8sReplicationControllers:
			fmt.Fprintf(view, "666666666666-")
		case k8sDaemonSets:
			fmt.Fprintf(view, "77777777777")
		case k8sJobs:
			fmt.Fprintf(view, "888888888888")
		case k8sPods:
			fmt.Fprintf(view, "99999999999")
		case k8sServices:
			fmt.Fprintf(view, "111111100000000")
		case k8sIngress:
			fmt.Fprintf(view, "11111111111111111111111111111111")
		case k8sPersistentVolumeClaims:
			fmt.Fprintf(view, "12222222222222222222é")
		case k8sSecrets:
			fmt.Fprintf(view, "13333333333333")
		case k8sConfigMaps:
			fmt.Fprintf(view, "14444444444444")
		}
		if _, err := setCurrentViewOnTop(g, view.Name()); err != nil {
			return err
		}
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
	fmt.Fprintf(v, " ^d       : Describe Kubernetes entity\n")
	fmt.Fprintf(v, " ^q       : Quit\n")
	fmt.Fprintf(v, " ^h       : Show help message\n")
}
