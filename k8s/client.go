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

package k8s

import (
	"log"

	"k8s.io/client-go/1.4/kubernetes"
	"k8s.io/client-go/1.4/pkg/api"
	"k8s.io/client-go/1.4/pkg/api/v1"
	"k8s.io/client-go/1.4/pkg/apis/extensions/v1beta1"
	"k8s.io/client-go/1.4/tools/clientcmd"
)

// Client define a Kubernetes client.
type Client struct {
	Clientset *kubernetes.Clientset
}

// NewKubernetesClient create new Kubernetes client using configuration from kubectl.
func NewKubernetesClient(kubeconfigPath string) (*Client, error) {
	// uses the current context in kubeconfig
	log.Printf("[DEBUG] Load Kubernetes configuration from %s", kubeconfigPath)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] Creates the Kubernetes clientset")
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return &Client{
		Clientset: clientset,
	}, nil
}

func (client *Client) GetNamespaces() (*v1.NamespaceList, error) {
	return client.Clientset.Core().Namespaces().List(api.ListOptions{})
}

func (client *Client) GetNodes() (*v1.NodeList, error) {
	return client.Clientset.Core().Nodes().List(api.ListOptions{})
}

func (client *Client) GetNode(name string) (*v1.Node, error) {
	return client.Clientset.Core().Nodes().Get(name)
}

func (client *Client) GetPersistentVolumes() (*v1.PersistentVolumeList, error) {
	return client.Clientset.Core().PersistentVolumes().List(api.ListOptions{})
}

func (client *Client) GetDeployments() (*v1beta1.DeploymentList, error) {
	return client.Clientset.Extensions().Deployments("").List(api.ListOptions{})
}

func (client *Client) GetReplicaSets() (*v1beta1.ReplicaSetList, error) {
	return client.Clientset.Extensions().ReplicaSets("").List(api.ListOptions{})
}

func (client *Client) GetReplicationControllers() (*v1.ReplicationControllerList, error) {
	return client.Clientset.Core().ReplicationControllers("").List(api.ListOptions{})
}

func (client *Client) GetDaemonSets() (*v1beta1.DaemonSetList, error) {
	return client.Clientset.Extensions().DaemonSets("").List(api.ListOptions{})
}

func (client *Client) GetJobs() (*v1beta1.JobList, error) {
	return client.Clientset.Extensions().Jobs("").List(api.ListOptions{})
}

func (client *Client) GetPods() (*v1.PodList, error) {
	return client.Clientset.Core().Pods("").List(api.ListOptions{})
}

func (client *Client) GetServices() (*v1.ServiceList, error) {
	return client.Clientset.Core().Services("").List(api.ListOptions{})
}

func (client *Client) GetIngresses() (*v1beta1.IngressList, error) {
	return client.Clientset.Extensions().Ingresses("").List(api.ListOptions{})
}

func (client *Client) GetPersistentVolumeClaims() (*v1.PersistentVolumeClaimList, error) {
	return client.Clientset.Core().PersistentVolumeClaims("").List(api.ListOptions{})
}

func (client *Client) GetSecrets() (*v1.SecretList, error) {
	return client.Clientset.Core().Secrets("").List(api.ListOptions{})
}

func (client *Client) GetConfigMaps() (*v1.ConfigMapList, error) {
	return client.Clientset.Core().ConfigMaps("").List(api.ListOptions{})
}
