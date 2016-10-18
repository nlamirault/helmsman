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

	"github.com/jroimartin/gocui"

	"github.com/nlamirault/helmsman/k8s"
)

func printK8SNodes(v *gocui.View, client *k8s.Client) {
	nodes, err := client.GetNodes()
	if err != nil {
		fmt.Fprintf(v, "\033[31;01mKubernetes error:\n%s\033[0m", err.Error())
	} else {
		fmt.Fprintf(v, "Nodes:\n\n")
		for _, node := range nodes.Items {
			fmt.Fprintf(v, "> \033[32;01m%s\033[0m\n", node.Name)
			fmt.Fprintf(v, "  * Creation: %s\n", node.CreationTimestamp)
			fmt.Fprintf(v, "  * Labels:\n")
			for _, label := range node.Labels {
				fmt.Fprintf(v, "    - %s\n", label)
			}
		}
	}
}

func printK8SNamespaces(v *gocui.View, client *k8s.Client) {
	namespaces, err := client.GetNamespaces()
	if err != nil {
		fmt.Fprintf(v, "\033[31;01mKubernetes error:\n%s\033[0m", err.Error())
	} else {
		fmt.Fprintf(v, "Namespaces:\n\n")
		for _, namespace := range namespaces.Items {
			fmt.Fprintf(v, "> \033[32;01m%s\033[0m\n", namespace.Name)
			fmt.Fprintf(v, "  * Creation: %s\n", namespace.CreationTimestamp)
		}
	}
}

func printK8SServices(v *gocui.View, client *k8s.Client) {
	services, err := client.GetServices()
	if err != nil {
		fmt.Fprintf(v, "\033[31;01mKubernetes error:\n%s\033[0m", err.Error())
	} else {
		fmt.Fprintf(v, "Services:\n\n")
		for _, service := range services.Items {
			fmt.Fprintf(v, "> \033[32;01m%s\033[0m\n", service.Name)
			fmt.Fprintf(v, "  * Namespace: %s\n", service.Namespace)
			fmt.Fprintf(v, "  * Creation: %s\n", service.CreationTimestamp)
			fmt.Fprintf(v, "  * Ports:\n")
			for _, port := range service.Spec.Ports {
				fmt.Fprintf(v, "    - %s [%s] %s -> %s\n", port.Name, port.Protocol, port.Port, port.TargetPort)
			}
			fmt.Fprintf(v, "  * External IPs:\n")
			for _, ip := range service.Spec.ExternalIPs {
				fmt.Fprintf(v, "    - %s\n", ip)
			}
			fmt.Fprintf(v, "  * Labels:\n")
			for _, label := range service.Labels {
				fmt.Fprintf(v, "    - %s\n", label)
			}
			fmt.Fprintf(v, "  * ClusterIP: %s\n", service.Spec.ClusterIP)
		}
	}
}

func printK8SPersistentVolumes(v *gocui.View, client *k8s.Client) {
	volumes, err := client.GetPersistentVolumes()
	if err != nil {
		fmt.Fprintf(v, "\033[31;01mKubernetes error:\n%s\033[0m", err.Error())
	} else {
		fmt.Fprintf(v, "Persistent Volumes:\n\n")
		for _, volume := range volumes.Items {
			fmt.Fprintf(v, "> \033[32;01m%s\033[0m\n", volume.Name)
		}
	}
}

func printK8SDeployments(v *gocui.View, client *k8s.Client) {
	deployments, err := client.GetDeployments()
	if err != nil {
		fmt.Fprintf(v, "\033[31;01mKubernetes error:\n%s\033[0m", err.Error())
	} else {
		fmt.Fprintf(v, "Deployments:\n\n")
		for _, deploy := range deployments.Items {
			fmt.Fprintf(v, "> \033[32;01m%s\033[0m\n", deploy.Name)
		}
	}
}

func printK8SReplicaSets(v *gocui.View, client *k8s.Client) {
	replicasets, err := client.GetReplicaSets()
	if err != nil {
		fmt.Fprintf(v, "\033[31;01mKubernetes error:\n%s\033[0m", err.Error())
	} else {
		fmt.Fprintf(v, "Replica Sets:\n\n")
		for _, replica := range replicasets.Items {
			fmt.Fprintf(v, "> \033[32;01m%s\033[0m\n", replica.Name)
		}
	}
}

func printK8SReplicationControllers(v *gocui.View, client *k8s.Client) {
	rcs, err := client.GetReplicationControllers()
	if err != nil {
		fmt.Fprintf(v, "\033[31;01mKubernetes error:\n%s\033[0m", err.Error())
	} else {
		fmt.Fprintf(v, "Replication Controllers:\n\n")
		for _, rc := range rcs.Items {
			fmt.Fprintf(v, "> \033[32;01m%s\033[0m\n", rc.Name)
		}
	}
}

func printK8SDaemonSets(v *gocui.View, client *k8s.Client) {
	sets, err := client.GetDaemonSets()
	if err != nil {
		fmt.Fprintf(v, "\033[31;01mKubernetes error:\n%s\033[0m", err.Error())
	} else {
		fmt.Fprintf(v, "Daemon Sets:\n\n")
		for _, ds := range sets.Items {
			fmt.Fprintf(v, "> \033[32;01m%s\033[0m\n", ds.Name)
		}
	}
}

func printK8SJobs(v *gocui.View, client *k8s.Client) {
	jobs, err := client.GetJobs()
	if err != nil {
		fmt.Fprintf(v, "\033[31;01mKubernetes error:\n%s\033[0m", err.Error())
	} else {
		fmt.Fprintf(v, "Jobs:\n\n")
		for _, job := range jobs.Items {
			fmt.Fprintf(v, "> \033[32;01m%s\033[0m\n", job.Name)
		}
	}
}

func printK8SPods(v *gocui.View, client *k8s.Client) {
	pods, err := client.GetPods()
	if err != nil {
		fmt.Fprintf(v, "\033[31;01mKubernetes error:\n%s\033[0m", err.Error())
	} else {
		fmt.Fprintf(v, "Pods:\n\n")
		for _, pod := range pods.Items {
			fmt.Fprintf(v, "> \033[32;01m%s\033[0m\n", pod.Name)
		}
	}
}

func printK8SIngresses(v *gocui.View, client *k8s.Client) {
	ingresses, err := client.GetIngresses()
	if err != nil {
		fmt.Fprintf(v, "\033[31;01mKubernetes error:\n%s\033[0m", err.Error())
	} else {
		fmt.Fprintf(v, "Ingress:\n\n")
		for _, ingress := range ingresses.Items {
			fmt.Fprintf(v, "> \033[32;01m%s\033[0m\n", ingress.Name)
		}
	}
}

func printK8SPersistentVolumeClaims(v *gocui.View, client *k8s.Client) {
	pvcs, err := client.GetPersistentVolumeClaims()
	if err != nil {
		fmt.Fprintf(v, "\033[31;01mKubernetes error:\n%s\033[0m", err.Error())
	} else {
		fmt.Fprintf(v, "Persistent Volume Claims:\n\n")
		for _, pvc := range pvcs.Items {
			fmt.Fprintf(v, "> \033[32;01m%s\033[0m\n", pvc.Name)
		}
	}
}

func printK8SSecrets(v *gocui.View, client *k8s.Client) {
	secrets, err := client.GetSecrets()
	if err != nil {
		fmt.Fprintf(v, "\033[31;01mKubernetes error:\n%s\033[0m", err.Error())
	} else {
		fmt.Fprintf(v, "Secrets:\n\n")
		for _, secret := range secrets.Items {
			fmt.Fprintf(v, "> \033[32;01m%s\033[0m\n", secret.Name)
		}
	}
}

func printK8SConfigMaps(v *gocui.View, client *k8s.Client) {
	configmaps, err := client.GetConfigMaps()
	if err != nil {
		fmt.Fprintf(v, "\033[31;01mKubernetes error:\n%s\033[0m", err.Error())
	} else {
		fmt.Fprintf(v, "Config Maps:\n\n")
		for _, configmap := range configmaps.Items {
			fmt.Fprintf(v, "> \033[32;01m%s\033[0m\n", configmap.Name)
		}
	}
}
