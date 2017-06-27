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
	// "time"

	"github.com/golang/glog"
	"github.com/jroimartin/gocui"

	"github.com/nlamirault/helmsman/k8s"
)

type TUI struct {
	Gocui            *gocui.Gui
	KubernetesClient *k8s.Client
}

func (tui *TUI) Setup(k8sclient *k8s.Client) {
	glog.V(2).Info("Create UI")
	var err error

	g := gocui.NewGui()
	// g.ShowCursor = true

	if err = g.Init(); err != nil {
		glog.Fatalf("UI Failed: %s", err)
	}

	defer g.Close()

	g.SetLayout(layout)
	g.Cursor = true
	g.FgColor = gocui.ColorWhite
	g.BgColor = gocui.ColorBlack
	g.SelFgColor = gocui.ColorYellow
	g.SelBgColor = gocui.ColorBlack
	tui.Gocui = g
	tui.KubernetesClient = k8sclient

	err = tui.registerKeybindings()
	if err != nil {
		glog.Fatalf("Initialize keybindings failed: %s", err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		glog.Fatalf("UI failed: %s", err)
	}
}

// writeView writes string to view
func (tui *TUI) writeView(name string, text string) {
	v, _ := tui.Gocui.View(name)
	v.Clear()
	fmt.Fprint(v, text)
	v.SetCursor(len(text), 0)
}

// createview creates a new view
func (tui *TUI) createView(viewName string, x1, y1, x2, y2 int) error {
	if _, err := tui.Gocui.SetView(viewName, x1, y1, x2, y2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// v.Title = p.title
		//v.Editor = p.editor
		//v.Editable = p.editable
		tui.writeView(viewName, "hello world")
	}
	return nil
}

// clearView clears a view
// func (tui *TUI) clearView(name string) {
// 	v, _ := tui.Gocui.View(name)
// 	v.Clear()
// 	v.SetOrigin(0, 0)
// 	v.SetCursor(0, 0)
// }

// func (tui *TUI) closeView(name string) error {
// 	g.DeleteView(v.Name())
// 	if _, err := setCurrentViewOnTop(g, mainViewName); err != nil {
// 		return err
// 	}
// 	return nil
// }

// setView set the focus to the view specified by name
func (tui *TUI) setView(name string) error {
	if _, err := tui.Gocui.SetCurrentView(name); err != nil {
		return err
	}
	return nil
}

// showModal shows a modal dialog on top of other views
func (tui *TUI) showModal(name, text string, width float64, height float64) {
	fmt.Printf("Shpw modal")
	// p := vp[name]
	//p.text = text
	//vp[name] = p

	maxX, maxY := tui.Gocui.Size()

	modalWidth := int(float64(maxX) * width)
	modalHeight := int(float64(maxY) * height)

	x1 := (maxX - modalWidth) / 2
	x2 := x1 + modalWidth
	y1 := (maxY - modalHeight) / 2
	y2 := y1 + modalHeight

	tui.createView(name, x1, y1, x2, y2)
	tui.setView(name)
}
