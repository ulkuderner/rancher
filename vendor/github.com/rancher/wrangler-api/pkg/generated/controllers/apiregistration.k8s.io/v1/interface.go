/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package v1

import (
	"github.com/rancher/wrangler/pkg/generic"
	v1 "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1"
	clientset "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/typed/apiregistration/v1"
	informers "k8s.io/kube-aggregator/pkg/client/informers/externalversions/apiregistration/v1"
)

type Interface interface {
	APIService() APIServiceController
}

func New(controllerManager *generic.ControllerManager, client clientset.ApiregistrationV1Interface,
	informers informers.Interface) Interface {
	return &version{
		controllerManager: controllerManager,
		client:            client,
		informers:         informers,
	}
}

type version struct {
	controllerManager *generic.ControllerManager
	informers         informers.Interface
	client            clientset.ApiregistrationV1Interface
}

func (c *version) APIService() APIServiceController {
	return NewAPIServiceController(v1.SchemeGroupVersion.WithKind("APIService"), c.controllerManager, c.client, c.informers.APIServices())
}