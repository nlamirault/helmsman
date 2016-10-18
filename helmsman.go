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

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nlamirault/helmsman/k8s"
	"github.com/nlamirault/helmsman/ui"
	"github.com/nlamirault/helmsman/version"
)

func main() {
	var (
		showVersion = flag.Bool("version", false, "Print version information.")
		kubeconfig  = flag.String("kubeconfig", "./config", "Absolute path to the kubeconfig file")
	)
	flag.Parse()

	if *showVersion {
		fmt.Printf("Helsman. The Kubernetes Text UI. v%s\n", version.Version)
		os.Exit(0)
	}

	k8sclient, err := k8s.NewKubernetesClient(*kubeconfig)
	if err != nil {
		log.Printf("[ERROR] Kubernetes client failed: %s", err.Error())
		os.Exit(1)
	}

	log.Printf("[INFO] Helmsman using :%s", k8sclient)
	tui := ui.TUI{}
	// gui.Setup(func(gui *console.GUI) {
	// 	gui.Gocui.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone,
	// 		func(_gocui *gocui.Gui, _v *gocui.View) error {
	// 			return inputhandler.Handle(stdHandler, travianHandler, trav, gui, _v)
	// 		})

	// 	gui.Println("Welcome to Hellsman")
	// })
	tui.Setup(k8sclient)

}
