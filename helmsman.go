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

package main

import (
	"flag"
	"fmt"
	// "log"
	"os"

	"github.com/golang/glog"

	"github.com/nlamirault/helmsman/k8s"
	"github.com/nlamirault/helmsman/ui"
	vers "github.com/nlamirault/helmsman/version"
)

const (
	// BANNER is what is printed for help/info output.
	BANNER = "Hellsman - v%s\n"
)

var (
	debug      bool
	version    bool
	kubeconfig string
)

func usageAndExit(message string, exitCode int) {
	if message != "" {
		fmt.Fprintf(os.Stderr, message)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(exitCode)
}

func init() {
	flag.BoolVar(&version, "version", false, "print version and exit")
	flag.BoolVar(&debug, "d", false, "run in debug mode")
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Absolute path to the kubeconfig file")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(BANNER, vers.Version))
		flag.PrintDefaults()
	}
	flag.Parse()

	if version {
		fmt.Printf("Helsman. The Kubernetes Text UI. v%s\n", vers.Version)
		os.Exit(0)
	}

	if kubeconfig == "" {
		usageAndExit("kubeconfig filename cannot be empty.", 1)
	}
}

func main() {
	k8sclient, err := k8s.NewKubernetesClient(kubeconfig)
	if err != nil {
		glog.Errorf("[ERROR] Kubernetes client failed: %s", err.Error())
		os.Exit(1)
	}

	glog.Infof("Helmsman using :%s", k8sclient)
	tui := ui.TUI{}
	tui.Setup(k8sclient)
}
