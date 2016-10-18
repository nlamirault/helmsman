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
	// Submit a line
	// if err := g.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, inputLineHandler); err != nil {
	// 	return err
	// }

	// Tab
	if err := g.SetKeybinding("side", gocui.KeyTab, gocui.ModNone, nextViewHandler); err != nil {
		return err
	}
	if err := g.SetKeybinding("main", gocui.KeyTab, gocui.ModNone, nextViewHandler); err != nil {
		return err
	}

	// Cursors
	if err := g.SetKeybinding("side", gocui.KeyArrowDown, gocui.ModNone, cursorDownHandler); err != nil {
		return err
	}
	if err := g.SetKeybinding("side", gocui.KeyArrowUp, gocui.ModNone, cursorUpHandler); err != nil {
		return err
	}
	if err := g.SetKeybinding("side", gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return kubernetesDispatcher(g, v, k8sclient)
	}); err != nil {
		return err
	}
	if err := g.SetKeybinding("main", gocui.KeyArrowDown, gocui.ModNone, cursorDownHandler); err != nil {
		return err
	}
	if err := g.SetKeybinding("main", gocui.KeyArrowUp, gocui.ModNone, cursorUpHandler); err != nil {
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

func nextViewHandler(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() == "side" {
		return g.SetCurrentView("main")
		// } else if v.Name() == "main" {
		// 	return g.SetCurrentView("input")
		// } else if v.Name() == "input" {
		// 	return g.SetCurrentView("side")
	}
	return g.SetCurrentView("side")
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

		mainView, err := g.View("main")
		if err != nil {
			return err
		}
		mainView.Clear()
		// mainView.Editable = true
		mainView.Highlight = true
		mainView.Title = l
		mainView.SetCursor(0, 0)
		mainView.SetOrigin(0, 0)
		// fmt.Fprintf(mainView, "----> %s\n", l)
		switch l {
		case "Namespaces":
			printK8SNamespaces(mainView, client)
		case "Nodes":
			printK8SNodes(mainView, client)
		case "Persitent Volumes":
			printK8SPersistentVolumes(mainView, client)
		case "Deployments":
			printK8SDeployments(mainView, client)
		case "Replica Sets":
			printK8SReplicaSets(mainView, client)
		case "Replication Controllers":
			printK8SReplicationControllers(mainView, client)
		case "Daemon Sets":
			printK8SDaemonSets(mainView, client)
		case "Jobs":
			printK8SJobs(mainView, client)
		case "Pods":
			printK8SPods(mainView, client)
		case "Services":
			printK8SServices(mainView, client)
		case "Ingress":
			printK8SIngresses(mainView, client)
		case "Persistent Volume Claims":
			printK8SPersistentVolumeClaims(mainView, client)
		case "Secrets":
			printK8SSecrets(mainView, client)
		case "Config Maps":
			printK8SConfigMaps(mainView, client)
		}
		return g.SetCurrentView("main")
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
	ov, _ := g.View("main")

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
