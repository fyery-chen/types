//go:generate go run generator/cleanup/main.go
//go:generate go run main.go

package main

import (
	businessSchema "github.com/rancher/types/apis/cloud.huawei.com/v3/schema"
	managementSchema "github.com/rancher/types/apis/management.cattle.io/v3/schema"
	managementPublicSchema "github.com/rancher/types/apis/management.cattle.io/v3public/schema"
	"github.com/rancher/types/generator"
	"k8s.io/api/apps/v1beta2"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"

	"k8s.io/api/core/v1"
	extv1beta1 "k8s.io/api/extensions/v1beta1"
	knetworkingv1 "k8s.io/api/networking/v1"
	rbacv1 "k8s.io/api/rbac/v1"
)

func main() {
	generator.Generate(managementSchema.Schemas)
	generator.Generate(businessSchema.Schemas)
	generator.Generate(managementPublicSchema.PublicSchemas)
	generator.GenerateNativeTypes(v1.SchemeGroupVersion, []interface{}{
		v1.Endpoints{},
		v1.Pod{},
		v1.Service{},
		v1.Secret{},
		v1.ConfigMap{},
		v1.ServiceAccount{},
		v1.ReplicationController{},
	}, []interface{}{
		v1.Node{},
		v1.ComponentStatus{},
		v1.Namespace{},
		v1.Event{},
	})
	generator.GenerateNativeTypes(v1beta2.SchemeGroupVersion, []interface{}{
		v1beta2.Deployment{},
		v1beta2.DaemonSet{},
		v1beta2.StatefulSet{},
		v1beta2.ReplicaSet{},
	}, nil)
	generator.GenerateNativeTypes(rbacv1.SchemeGroupVersion, []interface{}{
		rbacv1.RoleBinding{},
		rbacv1.Role{},
	}, []interface{}{
		rbacv1.ClusterRoleBinding{},
		rbacv1.ClusterRole{},
	})
	generator.GenerateNativeTypes(knetworkingv1.SchemeGroupVersion, []interface{}{
		knetworkingv1.NetworkPolicy{},
	}, nil)
	generator.GenerateNativeTypes(batchv1.SchemeGroupVersion, []interface{}{
		batchv1.Job{},
	}, nil)
	generator.GenerateNativeTypes(batchv1beta1.SchemeGroupVersion, []interface{}{
		batchv1beta1.CronJob{},
	}, nil)
	generator.GenerateNativeTypes(extv1beta1.SchemeGroupVersion,
		[]interface{}{
			extv1beta1.Ingress{},
		},
		[]interface{}{
			extv1beta1.PodSecurityPolicy{},
		},
	)
}
