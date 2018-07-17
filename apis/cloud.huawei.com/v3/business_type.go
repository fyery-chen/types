package v3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Business struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec BusinessQuotaSpec `json:"spec"`
	// Most recent observed status of the alert. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status BusinessStatus `json:"status"`
}

type BusinessStatus struct {
	BusinessState string `json:"businessState,omitempty" norman:"options=transferring|transferError|active,default=active"`
}

type BusinessQuotaSpec struct {
	DisplayName string `json:"displayName,omitempty" norman:"required"`
	Description string `json:"description,omitempty"`
	NodeCount   int    `json:"nodeCount,omitempty"`
}

type BusinessQuotaCheck struct {
	BusinessName string `json:"businessName,omitempty" norman:"type=string,required"`
	NodeCount    int    `json:"nodeCount,omitempty" norman:"required"`
}

type BusinessQuotaCheckOutput struct {
	Message string `json:"message,omitempty"`
}
