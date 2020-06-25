package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NovaMigrationTargetSpec defines the desired state of NovaMigrationTarget
// +k8s:openapi-gen=true
type NovaMigrationTargetSpec struct {
	// name of configmap which holds general information on the OSP env
	CommonConfigMap string `json:"commonConfigMap"`
	// container image to run for the daemon
	NovaComputeImage string `json:"novaComputeImage"`
	// SSHD port
	SshdPort int32 `json:"sshdPort"`
	// service account used to create pods
	ServiceAccount string `json:"serviceAccount"`
	// Name of the worker role created for OSP computes
	RoleName string `json:"roleName"`
}

// NovaMigrationTargetStatus defines the observed state of NovaMigrationTarget
// +k8s:openapi-gen=true
type NovaMigrationTargetStatus struct {
	// Count is the number of nodes the daemon is deployed to
	Count int32 `json:"count"`
	// Daemonset hash used to detect changes
	DaemonsetHash string `json:"daemonsetHash"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NovaMigrationTarget is the Schema for the novamigrationtargets API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type NovaMigrationTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NovaMigrationTargetSpec   `json:"spec,omitempty"`
	Status NovaMigrationTargetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NovaMigrationTargetList contains a list of NovaMigrationTarget
type NovaMigrationTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NovaMigrationTarget `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NovaMigrationTarget{}, &NovaMigrationTargetList{})
}
