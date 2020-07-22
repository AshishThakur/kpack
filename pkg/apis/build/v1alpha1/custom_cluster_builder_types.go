package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
)

const CustomClusterBuilderKind = "CustomClusterBuilder"

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object,k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMetaAccessor

// +k8s:openapi-gen=true
type CustomClusterBuilder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CustomClusterBuilderSpec `json:"spec"`
	Status CustomBuilderStatus      `json:"status"`
}

// +k8s:openapi-gen=true
type CustomClusterBuilderSpec struct {
	CustomBuilderSpec `json:",inline"`
	ServiceAccountRef corev1.ObjectReference `json:"serviceAccountRef,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +k8s:openapi-gen=true
type CustomClusterBuilderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	// +listType
	Items []CustomClusterBuilder `json:"items"`
}

func (*CustomClusterBuilder) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind(CustomClusterBuilderKind)
}

func (c *CustomClusterBuilder) NamespacedName() types.NamespacedName {
	return types.NamespacedName{Namespace: c.Namespace, Name: c.Name}
}
