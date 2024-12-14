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

package v1beta1

import (
	v1alpha1 "istio.io/api/meta/v1alpha1"
	securityv1beta1 "istio.io/api/security/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AuthorizationPolicy enables access control on workloads.
//
// <!-- crd generation tags
// +cue-gen:AuthorizationPolicy:groupName:security.istio.io
// +cue-gen:AuthorizationPolicy:versions:v1beta1,v1
// +cue-gen:AuthorizationPolicy:storageVersion
// +cue-gen:AuthorizationPolicy:annotations:helm.sh/resource-policy=keep
// +cue-gen:AuthorizationPolicy:labels:app=istio-pilot,chart=istio,istio=security,heritage=Tiller,release=istio
// +cue-gen:AuthorizationPolicy:subresource:status
// +cue-gen:AuthorizationPolicy:scope:Namespaced
// +cue-gen:AuthorizationPolicy:resource:categories=istio-io,security-istio-io,shortNames=ap,plural=authorizationpolicies
// +cue-gen:AuthorizationPolicy:preserveUnknownFields:false
// +cue-gen:AuthorizationPolicy:printerColumn:name=Action,type=string,JSONPath=.spec.action,description="The operation to take."
// +cue-gen:AuthorizationPolicy:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=security.istio.io/v1beta1
// +genclient
// +k8s:deepcopy-gen=true
// -->
// +kubebuilder:validation:XValidation:message="only one of targetRefs or selector can be set",rule="(has(self.selector)?1:0)+(has(self.targetRef)?1:0)+(has(self.targetRefs)?1:0)<=1"
type AuthorizationPolicy struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec securityv1beta1.AuthorizationPolicy `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status v1alpha1.IstioStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AuthorizationPolicyList is a collection of AuthorizationPolicies.
type AuthorizationPolicyList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []*AuthorizationPolicy `json:"items" protobuf:"bytes,2,rep,name=items"`
}

//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// <!-- crd generation tags
// +cue-gen:PeerAuthentication:groupName:security.istio.io
// +cue-gen:PeerAuthentication:versions:v1beta1,v1
// +cue-gen:PeerAuthentication:storageVersion
// +cue-gen:PeerAuthentication:annotations:helm.sh/resource-policy=keep
// +cue-gen:PeerAuthentication:labels:app=istio-pilot,chart=istio,istio=security,heritage=Tiller,release=istio
// +cue-gen:PeerAuthentication:subresource:status
// +cue-gen:PeerAuthentication:scope:Namespaced
// +cue-gen:PeerAuthentication:resource:categories=istio-io,security-istio-io,shortNames=pa
// +cue-gen:PeerAuthentication:preserveUnknownFields:false
// +cue-gen:PeerAuthentication:printerColumn:name=Mode,type=string,JSONPath=.spec.mtls.mode,description="Defines the mTLS mode used for peer authentication."
// +cue-gen:PeerAuthentication:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=security.istio.io/v1beta1
// +genclient
// +k8s:deepcopy-gen=true
// -->
// +kubebuilder:validation:XValidation:message="portLevelMtls requires selector",rule="(has(self.selector) && has(self.selector.matchLabels) && self.selector.matchLabels.size() > 0) || !has(self.portLevelMtls)"
type PeerAuthentication struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec securityv1beta1.PeerAuthentication `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status v1alpha1.IstioStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PeerAuthenticationList is a collection of PeerAuthentications.
type PeerAuthenticationList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []*PeerAuthentication `json:"items" protobuf:"bytes,2,rep,name=items"`
}

//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// <!-- crd generation tags
// +cue-gen:RequestAuthentication:groupName:security.istio.io
// +cue-gen:RequestAuthentication:versions:v1beta1,v1
// +cue-gen:RequestAuthentication:storageVersion
// +cue-gen:RequestAuthentication:annotations:helm.sh/resource-policy=keep
// +cue-gen:RequestAuthentication:labels:app=istio-pilot,chart=istio,istio=security,heritage=Tiller,release=istio
// +cue-gen:RequestAuthentication:subresource:status
// +cue-gen:RequestAuthentication:scope:Namespaced
// +cue-gen:RequestAuthentication:resource:categories=istio-io,security-istio-io,shortNames=ra
// +cue-gen:RequestAuthentication:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=security.istio.io/v1beta1
// +genclient
// +k8s:deepcopy-gen=true
// -->
// +kubebuilder:validation:XValidation:message="only one of targetRefs or selector can be set",rule="(has(self.selector)?1:0)+(has(self.targetRef)?1:0)+(has(self.targetRefs)?1:0)<=1"
type RequestAuthentication struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec securityv1beta1.RequestAuthentication `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status v1alpha1.IstioStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RequestAuthenticationList is a collection of RequestAuthentications.
type RequestAuthenticationList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []*RequestAuthentication `json:"items" protobuf:"bytes,2,rep,name=items"`
}
