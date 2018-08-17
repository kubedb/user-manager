package v1alpha1

import (
	"reflect"

	"github.com/appscode/go/log"
	crdutils "github.com/appscode/kutil/apiextensions/v1beta1"
	meta_util "github.com/appscode/kutil/meta"
	"github.com/golang/glog"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

func (r MysqlRole) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crdutils.NewCustomResourceDefinition(crdutils.Config{
		Group:         SchemeGroupVersion.Group,
		Plural:        ResourceMysqlRoles,
		Singular:      ResourceMysqlRole,
		Kind:          ResourceKindMysqlRole,
		Categories:    []string{"datastore", "kubedb", "appscode", "all"},
		ResourceScope: string(apiextensions.NamespaceScoped),
		Versions: []apiextensions.CustomResourceDefinitionVersion{
			{
				Name:    SchemeGroupVersion.Version,
				Served:  true,
				Storage: true,
			},
		},
		Labels: crdutils.Labels{
			LabelsMap: map[string]string{"app": "kubedb"},
		},
		SpecDefinitionName:      "github.com/kubedb/user-manager/apis/authorization/v1alpha1.MysqlRole",
		EnableValidation:        true,
		GetOpenAPIDefinitions:   GetOpenAPIDefinitions,
		EnableStatusSubresource: EnableStatusSubresource,
	})
}

func (r MysqlRole) IsValid() error {
	return nil
}

func (r *MysqlRole) AlreadyObserved(other *MysqlRole) bool {
	if r == nil {
		return other == nil
	}
	if other == nil { // && d != nil
		return false
	}
	if r == other {
		return true
	}

	var match bool

	if EnableStatusSubresource {
		match = r.Status.ObservedGeneration >= r.Generation
	} else {
		match = meta_util.Equal(r.Spec, other.Spec)
	}
	if match {
		match = reflect.DeepEqual(r.Labels, other.Labels)
	}
	if match {
		match = meta_util.EqualAnnotation(r.Annotations, other.Annotations)
	}

	if !match && bool(glog.V(log.LevelDebug)) {
		diff := meta_util.Diff(other, r)
		glog.V(log.LevelDebug).Infof("%s %s/%s has changed. Diff: %s", meta_util.GetKind(r), r.Namespace, r.Name, diff)
	}
	return match
}

func (b MysqlRoleBinding) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crdutils.NewCustomResourceDefinition(crdutils.Config{
		Group:         SchemeGroupVersion.Group,
		Plural:        ResourceMysqlRoleBindings,
		Singular:      ResourceMysqlRoleBinding,
		Kind:          ResourceKindMysqlRoleBinding,
		Categories:    []string{"datastore", "kubedb", "appscode", "all"},
		ResourceScope: string(apiextensions.NamespaceScoped),
		Versions: []apiextensions.CustomResourceDefinitionVersion{
			{
				Name:    SchemeGroupVersion.Version,
				Served:  true,
				Storage: true,
			},
		},
		Labels: crdutils.Labels{
			LabelsMap: map[string]string{"app": "kubedb"},
		},
		SpecDefinitionName:      "github.com/kubedb/user-manager/apis/authorization/v1alpha1.MysqlRoleBinding",
		EnableValidation:        true,
		GetOpenAPIDefinitions:   GetOpenAPIDefinitions,
		EnableStatusSubresource: EnableStatusSubresource,
	})
}

func (b MysqlRoleBinding) IsValid() error {
	return nil
}

func (b *MysqlRoleBinding) AlreadyObserved(other *MysqlRoleBinding) bool {
	if b == nil {
		return other == nil
	}
	if other == nil { // && d != nil
		return false
	}
	if b == other {
		return true
	}

	var match bool

	if EnableStatusSubresource {
		match = b.Status.ObservedGeneration >= b.Generation
	} else {
		match = meta_util.Equal(b.Spec, other.Spec)
	}
	if match {
		match = reflect.DeepEqual(b.Labels, other.Labels)
	}
	if match {
		match = meta_util.EqualAnnotation(b.Annotations, other.Annotations)
	}

	if !match && bool(glog.V(log.LevelDebug)) {
		diff := meta_util.Diff(other, b)
		glog.V(log.LevelDebug).Infof("%s %s/%s has changed. Diff: %s", meta_util.GetKind(b), b.Namespace, b.Name, diff)
	}
	return match
}
