// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by kubetype-gen. DO NOT EDIT.

package v1alpha1

import (
	extensionsv1alpha1 "istio.io/api/extensions/v1alpha1"
	metav1alpha1 "istio.io/api/meta/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WasmPlugin provides a mechanism to extend the functionality provided by
// the Istio proxy through WebAssembly filters.
//
// <!-- crd generation tags
// +cue-gen:WasmPlugin:groupName:extensions.istio.io
// +cue-gen:WasmPlugin:versions:v1alpha1
// +cue-gen:WasmPlugin:storageVersion
// +cue-gen:WasmPlugin:annotations:helm.sh/resource-policy=keep
// +cue-gen:WasmPlugin:labels:app=istio-pilot,chart=istio,heritage=Tiller,release=istio
// +cue-gen:WasmPlugin:subresource:status
// +cue-gen:WasmPlugin:spec:required
// +cue-gen:WasmPlugin:scope:Namespaced
// +cue-gen:WasmPlugin:releaseChannel:extended
// +cue-gen:WasmPlugin:resource:categories=istio-io,extensions-istio-io
// +cue-gen:WasmPlugin:preserveUnknownFields:pluginConfig
// +cue-gen:WasmPlugin:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=extensions.istio.io/v1alpha1
// +genclient
// +k8s:deepcopy-gen=true
// -->
// +kubebuilder:validation:XValidation:message="only one of targetRefs or selector can be set",rule="(has(self.selector)?1:0)+(has(self.targetRef)?1:0)+(has(self.targetRefs)?1:0)<=1"
type WasmPlugin struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec extensionsv1alpha1.WasmPlugin `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status metav1alpha1.IstioStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WasmPluginList is a collection of WasmPlugins.
type WasmPluginList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []*WasmPlugin `json:"items" protobuf:"bytes,2,rep,name=items"`
}
