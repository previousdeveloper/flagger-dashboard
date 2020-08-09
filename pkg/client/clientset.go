package client

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	initialContext  = ""
	flaggerGroup    = "flagger.app"
	flaggerVersion  = "v1beta1"
	flaggerResource = "canaries"
)

type K8sClient struct {
	k8sRest dynamic.Interface
}

func NewK8sClient() K8sOperation {
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{CurrentContext: initialContext, Timeout: "1s"}).ClientConfig()
	if err != nil {
		panic(err.Error())
	}
	dynamicClient, err := dynamic.NewForConfig(config)

	if err != nil {
		panic(err.Error())
	}
	return &K8sClient{k8sRest: dynamicClient}
}

type K8sOperation interface {
	GetResourceByNamespace(namespace string) ([]unstructured.Unstructured, error)
}

func (k8sApi *K8sClient) GetResourceByNamespace(namespace string) ([]unstructured.Unstructured, error) {
	virtualServiceGVR := schema.GroupVersionResource{
		Group:    flaggerGroup,
		Version:  flaggerVersion,
		Resource: flaggerResource,
	}

	virtualServices, err := k8sApi.k8sRest.Resource(virtualServiceGVR).Namespace(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return virtualServices.Items, nil
}
