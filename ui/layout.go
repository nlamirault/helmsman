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

	"github.com/nlamirault/helmsman/version"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	const menuWidth = 30
	const inputHeight = 1
	if v, err := g.SetView("side", 0, 0, menuWidth-1, maxY-inputHeight); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		// v.BgColor = gocui.AttrBold
		v.Editable = false
		// v.Highlight = true
		v.Title = "Kubernetes"
		// fmt.Fprintf(v, "%s\n", "Loading...")
		fmt.Fprintf(v, "\033[34;01m%s\033[0m\n", "Admin")
		fmt.Fprintf(v, "%s\n", "===========================")
		fmt.Fprintf(v, "\033[32;01m%s\033[0m\n", "Namespaces")
		fmt.Fprintf(v, "\033[32;01m%s\033[0m\n", "Nodes")
		fmt.Fprintf(v, "\033[32;01m%s\033[0m\n", "Persistent Volumes")
		fmt.Fprintf(v, "\n\033[34;01m%s\033[0m\n", "Namespace")
		fmt.Fprintf(v, "%s\n", "===========================")
		fmt.Fprintf(v, "\n\033[34;01m%s\033[0m\n", "Workloads")
		fmt.Fprintf(v, "%s\n", "===========================")
		fmt.Fprintf(v, "\033[32;01m%s\033[0m\n", "Deployments")
		fmt.Fprintf(v, "\033[32;01m%s\033[0m\n", "Replica Sets")
		fmt.Fprintf(v, "\033[32;01m%s\033[0m\n", "Replication Controllers")
		fmt.Fprintf(v, "\033[32;01m%s\033[0m\n", "Daemon Sets")
		fmt.Fprintf(v, "\033[32;01m%s\033[0m\n", "Pet Sets")
		fmt.Fprintf(v, "\033[32;01m%s\033[0m\n", "Jobs")
		fmt.Fprintf(v, "\033[32;01m%s\033[0m\n", "Pods")
		fmt.Fprintf(v, "\n%s\n", "Services and Discovery")
		fmt.Fprintf(v, "%s\n", "===========================")
		fmt.Fprintf(v, "\033[32;01m%s\033[0m\n", "Services")
		fmt.Fprintf(v, "\033[32;01m%s\033[0m\n", "Ingress")
		fmt.Fprintf(v, "\n%s\n", "Storage")
		fmt.Fprintf(v, "%s\n", "===========================")
		fmt.Fprintf(v, "\033[32;01m%s\033[0m\n", "Persistent Volume Claims")
		fmt.Fprintf(v, "\n\033[34;01m%s\033[0m\n", "Config")
		fmt.Fprintf(v, "%s\n", "===========================")
		fmt.Fprintf(v, "\033[32;01m%s\033[0m\n", "Secrets")
		fmt.Fprintf(v, "\033[32;01m%s\033[0m\n", "Config Maps")
	}
	if v, err := g.SetView("main", menuWidth, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		v.Wrap = true
		fmt.Fprintf(v, "\n\nWelcome to Helmsman v%s\n", version.Version)
		fmt.Fprintf(v, "Help:\n\n")
		fmt.Fprintf(v, " <TAB>  : Move between panes\n")
		fmt.Fprintf(v, " <keys> : Move cursor\n")
		fmt.Fprintf(v, " <C-q>  : Quit\n")
		if err := g.SetCurrentView("main"); err != nil {
			return err
		}
	}

	if v, err := g.SetView("input", -1, maxY-2, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		v.Wrap = true
	}
	// if err := g.SetCurrentView("side"); err != nil {
	// 	return err
	// }
	return nil
}
