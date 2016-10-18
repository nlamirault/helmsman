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

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	const menuWidth = 30
	const inputHeight = 1
	if v, err := g.SetView("side", 0, 0, menuWidth-1, maxY-inputHeight); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		v.Highlight = true
		v.Title = "Kubernetes"
		// fmt.Fprintf(v, "%s\n", "Loading...")
		fmt.Fprintf(v, "%s\n", "Admin")
		fmt.Fprintf(v, "%s\n", "=================")
		fmt.Fprintf(v, "%s\n", "Namespaces")
		fmt.Fprintf(v, "%s\n", "Nodes")
		fmt.Fprintf(v, "%s\n", "Persistent Volumes")
		fmt.Fprintf(v, "\n%s\n", "Namespace")
		fmt.Fprintf(v, "%s\n", "=================")
		fmt.Fprintf(v, "\n%s\n", "Workloads")
		fmt.Fprintf(v, "%s\n", "=================")
		fmt.Fprintf(v, "%s\n", "Deployments")
		fmt.Fprintf(v, "%s\n", "Replica Sets")
		fmt.Fprintf(v, "%s\n", "Replication Controllers")
		fmt.Fprintf(v, "%s\n", "Daemon Sets")
		fmt.Fprintf(v, "%s\n", "Pet Sets")
		fmt.Fprintf(v, "%s\n", "Jobs")
		fmt.Fprintf(v, "%s\n", "Pods")
		fmt.Fprintf(v, "\n%s\n", "Services and Discovery")
		fmt.Fprintf(v, "%s\n", "=================")
		fmt.Fprintf(v, "%s\n", "Services")
		fmt.Fprintf(v, "%s\n", "Ingress")
		fmt.Fprintf(v, "\n%s\n", "Storage")
		fmt.Fprintf(v, "%s\n", "=================")
		fmt.Fprintf(v, "%s\n", "Persistent Volume Claims")
		fmt.Fprintf(v, "\n%s\n", "Config")
		fmt.Fprintf(v, "%s\n", "=================")
		fmt.Fprintf(v, "%s\n", "Secrets")
		fmt.Fprintf(v, "%s\n", "Config Maps")
	}
	if v, err := g.SetView("main", menuWidth, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		v.Wrap = true
		// if err := g.SetCurrentView("main"); err != nil {
		// 	return err
		// }
	}

	if v, err := g.SetView("input", -1, maxY-2, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		v.Wrap = true
	}
	if err := g.SetCurrentView("side"); err != nil {
		return err
	}
	return nil
}
