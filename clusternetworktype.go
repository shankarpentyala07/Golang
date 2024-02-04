package main

import (
	"context"
	"fmt"
	"os"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	"github.com/openshift/client-go/config/clientset/versioned"
)

func main() {
	kubeconfig, err := getKubeConfig()
	if err != nil {
		fmt.Printf("Error getting kubeconfig: %v\n", err)
		os.Exit(1)
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Printf("Error building kubeconfig: %v\n", err)
		os.Exit(1)
	}

	ocClient, err := versioned.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error creating OpenShift client: %v\n", err)
		os.Exit(1)
	}

	networkType, err := getNetworkType(ocClient)
	if err != nil {
		fmt.Printf("Error getting network type: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Network Type: %s\n", networkType)
}

func getKubeConfig() (string, error) {
	var kubeconfig string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = home + "/.kube/config"
	} else {
		kubeconfig = os.Getenv("KUBECONFIG")
		if kubeconfig == "" {
			kubeconfig = "/etc/kubernetes/kubeconfig"
		}
	}
	return kubeconfig, nil
}

func getNetworkType(client *versioned.Clientset) (string, error) {
	ctx := context.TODO()
	clusterConfig, err := client.ConfigV1().Networks().Get(ctx, "cluster", metav1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("Error getting network cluster: %v", err)
	}
	return clusterConfig.Spec.NetworkType, nil
}
