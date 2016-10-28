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
	"github.com/jroimartin/gocui"
)

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func clearView(v *gocui.View) {
	v.Clear()
	v.SetOrigin(0, 0)
	v.SetCursor(0, 0)
}

func closeView(g *gocui.Gui, v *gocui.View) error {
	g.DeleteView(v.Name())
	if _, err := setCurrentViewOnTop(g, mainViewName); err != nil {
		return err
	}
	return nil
}
