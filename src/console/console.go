package console

import (
	"context"

	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
)

var EksConsoleGVR = schema.GroupVersionResource{
	Group:    eksConsoleGroup,
	Version:  eksConsoleAPIVersion,
	Resource: eksConsoleResource,
}

type Console struct {
	*dynamic.DynamicClient
	*logrus.Logger
}

func (c *Console) Create(ctx context.Context, namespace string, obj *unstructured.Unstructured) error {
	_, err := c.Resource(EksConsoleGVR).Namespace(namespace).
		Create(ctx, obj, metav1.CreateOptions{})
	if err != nil {
		c.Logger.Errorf("error creating the EKSConsole object.")
		return err
	}
	return nil
}

func (c *Console) Delete(ctx context.Context, namespacedName types.NamespacedName) error {
	return c.Resource(EksConsoleGVR).Namespace(namespacedName.Namespace).
		Delete(ctx, namespacedName.Name, metav1.DeleteOptions{})
}

func (c *Console) Get(ctx context.Context, namespacedName types.NamespacedName) (*unstructured.Unstructured, error) {
	obj, err := c.Resource(EksConsoleGVR).Namespace(namespacedName.Namespace).
		Get(ctx, namespacedName.Name, metav1.GetOptions{})
	if err != nil {
		c.Logger.Error("error fetching the Resource")

		return nil, err
	}

	return obj, nil
}
