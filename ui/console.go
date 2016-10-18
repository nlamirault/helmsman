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
	//"fmt"
	"log"
	// "time"

	"github.com/jroimartin/gocui"

	"github.com/nlamirault/helmsman/k8s"
)

// type KubernetesCommand struct {
// 	Command string
// }

// type TUI struct {
// 	Gocui          *gocui.Gui
// 	KubernetesChan chan<- KubernetesCommand
// }

// func NewTUI() *TUI {
// 	return &TUI{
// 		KubernetesChan: make(chan KubernetesCommand),
// 	}
// }

type TUI struct {
	Gocui            *gocui.Gui
	KubernetesClient *k8s.Client
}

func (tui *TUI) Setup(k8sclient *k8s.Client) {

	var err error

	g := gocui.NewGui()
	// g.ShowCursor = true

	if err = g.Init(); err != nil {
		log.Panicln(err)
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

	registerKeybindings(tui.Gocui, tui.KubernetesClient)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
