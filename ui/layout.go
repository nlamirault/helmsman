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
	// "log"

	"github.com/jroimartin/gocui"
)

const (
	sideViewName   = "side"
	mainViewName   = "main"
	detailViewName = "detail"
	inputViewName  = "input"
	menuWidth      = 30
	inputHeight    = 1
)

func layout(g *gocui.Gui) error {
	if err := setSideLayout(g); err != nil {
		return err
	}
	if err := setSummaryLayout(g); err != nil {
		return err
	}
	if err := setDetailLayout(g); err != nil {
		return err
	}
	if err := setDetailLayout(g); err != nil {
		return err
	}
	if _, err := g.SetViewOnTop(mainViewName); err != nil {
		return err
	}
	return nil
}

func setSideLayout(g *gocui.Gui) error {
	_, maxY := g.Size()
	if v, err := g.SetView(sideViewName, 0, 0, menuWidth-1, maxY-inputHeight); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		// v.BgColor = gocui.AttrBold
		v.Editable = false
		// v.Highlight = true
		v.Title = "Kubernetes"
		for _, key := range k8smenu {
			fmt.Fprintf(v, "\n\033[34;01m%s\033[0m\n", key)
			fmt.Fprintf(v, "%s\n", "===========================")
			for _, entry := range k8sdispatcher[key] {
				fmt.Fprintf(v, "\033[32;01m%s\033[0m\n", entry)
			}
		}
	}
	return nil
}

func setSummaryLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(mainViewName, menuWidth, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		v.Wrap = true
		printHelp(v)
		if _, err := g.SetCurrentView(mainViewName); err != nil {
			return err
		}
	}
	return nil
}

func setDetailLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(detailViewName, menuWidth, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		v.Wrap = true
		printHelp(v)
		if _, err := g.SetCurrentView(detailViewName); err != nil {
			return err
		}
	}
	return nil
}

func setInputLatout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(inputViewName, -1, maxY-2, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		v.Wrap = true
	}
	return nil
}
