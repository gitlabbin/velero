package restore

import (
	"fmt"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
)

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

// GetGroupVersionResource (gk string), gk string group and kind as: "apps.Deployment"
func GetGroupVersionResource(gk string) (*schema.GroupVersionResource, error) {
	c := discovery.NewDiscoveryClientForConfigOrDie(&rest.Config{})

	groupResources, err := restmapper.GetAPIGroupResources(c)
	if err != nil {
		return nil, errors.Wrap(err, "getting Dynamic client")
	}
	mapper := restmapper.NewDiscoveryRESTMapper(groupResources)

	mapping, err := mapper.RESTMapping(schema.ParseGroupKind("apps.Deployment"))
	fmt.Println(mapping.Resource)

	return &mapping.Resource, nil
}
