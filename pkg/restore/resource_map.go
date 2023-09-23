package restore

import "k8s.io/apimachinery/pkg/runtime/schema"

var kindResourceMap = map[string]schema.GroupVersionResource{
	"Deployment": {
		Group:    "apps",
		Version:  "v1",
		Resource: "deployments",
	},
	"ConfigMap": {
		Group:    "apps",
		Version:  "v1",
		Resource: "configmaps",
	},
	"Job": {
		Group:    "batch",
		Version:  "v1",
		Resource: "jobs",
	},
	"Secret": {
		Group:    "api",
		Version:  "v1",
		Resource: "secrets",
	},
	"Service": {
		Group:    "api",
		Version:  "v1",
		Resource: "services",
	},
	"StatefulSet": {
		Group:    "apps",
		Version:  "v1",
		Resource: "statefulsets",
	},
	"PersistentVolume": {
		Group:    "api",
		Version:  "v1",
		Resource: "persistentvolumes",
	},
	"CustomResourceDefinition": {
		Group:    "apiextensions.k8s.io",
		Version:  "v1beta1",
		Resource: "customresourcedefinitions",
	},
}
