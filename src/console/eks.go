package console

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var (
	eksConsoleAPIVersion = "v1"
	eksConsoleResource   = "eksconsoleshells"
	eksConsoleGroup      = "autoprovisioning.consoleshell.com"
	eksConsoleKind       = "EksConsoleShell"
)

func BuildEKSConsoleObj(name, label string) *unstructured.Unstructured {
	group := eksConsoleGroup + "/" + eksConsoleAPIVersion
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": group,
			"kind":       eksConsoleKind,
			"metadata": map[string]interface{}{
				"name": name,
				"labels": map[string]string{
					"consoleshell.com/tenat-name": label,
				},
			},
			"spec": map[string]int{
				"ttlSeconds": 600,
			},
		},
	}
}
