/*
Copyright 2021 The Kubernetes Authors.

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

package genericprovider

import (
	operatorv1 "sigs.k8s.io/cluster-api-operator/api/v1alpha1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ControlPlaneProviderKind     = "ControlPlaneProvider"
	ControlPlaneProviderListKind = "ControlPlaneProviderList"
)

type ControlPlaneProviderWrapper struct {
	//nolint: dupl
	*operatorv1.ControlPlaneProvider
}

func (c *ControlPlaneProviderWrapper) GetConditions() clusterv1.Conditions {
	//nolint: dupl
	return c.Status.Conditions
}

func (c *ControlPlaneProviderWrapper) SetConditions(conditions clusterv1.Conditions) {
	//nolint: dupl
	c.Status.Conditions = conditions
}

func (c *ControlPlaneProviderWrapper) GetSpec() operatorv1.ProviderSpec {
	//nolint: dupl
	return c.Spec.ProviderSpec
}

func (c *ControlPlaneProviderWrapper) SetSpec(in operatorv1.ProviderSpec) {
	//nolint: dupl
	c.Spec.ProviderSpec = in
}

func (c *ControlPlaneProviderWrapper) GetStatus() operatorv1.ProviderStatus {
	//nolint: dupl
	return c.Status.ProviderStatus
}

func (c *ControlPlaneProviderWrapper) SetStatus(in operatorv1.ProviderStatus) {
	//nolint: dupl
	c.Status.ProviderStatus = in
}

func (c *ControlPlaneProviderWrapper) GetObject() client.Object {
	//nolint: dupl
	return c.ControlPlaneProvider
}

type ControlPlaneProviderListWrapper struct {
	//nolint: dupl
	*operatorv1.ControlPlaneProviderList
}

func (c *ControlPlaneProviderListWrapper) GetItems() []GenericProvider {
	//nolint: dupl
	providers := []GenericProvider{}
	for _, provider := range c.Items {
		providers = append(providers, &ControlPlaneProviderWrapper{&provider})
	}

	return providers
}

func (c *ControlPlaneProviderListWrapper) GetObject() client.ObjectList {
	//nolint: dupl
	return c.ControlPlaneProviderList
}
