/*
Copyright 2021 The Pixiu Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/caoyingjunz/pixiu/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AdvancedDeployments returns a AdvancedDeploymentInformer.
	AdvancedDeployments() AdvancedDeploymentInformer
	// AdvancedImages returns a AdvancedImageInformer.
	AdvancedImages() AdvancedImageInformer
	// ImageSets returns a ImageSetInformer.
	ImageSets() ImageSetInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AdvancedDeployments returns a AdvancedDeploymentInformer.
func (v *version) AdvancedDeployments() AdvancedDeploymentInformer {
	return &advancedDeploymentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AdvancedImages returns a AdvancedImageInformer.
func (v *version) AdvancedImages() AdvancedImageInformer {
	return &advancedImageInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ImageSets returns a ImageSetInformer.
func (v *version) ImageSets() ImageSetInformer {
	return &imageSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
