package v3

import (
	"github.com/rancher/norman/condition"
	"github.com/rancher/norman/types"
	"github.com/rancher/types/apis/project.cattle.io/v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	MultiClusterAppConditionInstalled condition.Cond = "Installed"
	MultiClusterAppConditionDeployed  condition.Cond = "Deployed"
)

type MultiClusterApp struct {
	types.Namespaced
	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status

	Spec   MultiClusterAppSpec   `json:"spec"`
	Status MultiClusterAppStatus `json:"status"`
}

type MultiClusterAppSpec struct {
	TemplateVersionName string   `json:"templateVersionName,omitempty" norman:"type=reference[templateVersion],required"`
	Answers             []Answer `json:"answers,omitempty"`
	Targets             []Target `json:"targets,omitempty" norman:"required"`
}

type MultiClusterAppStatus struct {
	Conditions []v3.AppCondition `json:"conditions,omitempty"`
}

type Target struct {
	ProjectName string `json:"projectName,omitempty" norman:"type=reference[project],required"`
	AppName     string `json:"appName,omitempty" norman:"type=reference[v3/projects/schemas/app]"`
	Healthstate string `json:"healthState,omitempty"`
}

type Answer struct {
	ProjectName string            `json:"projectName,omitempty" norman:"type=reference[project]"`
	ClusterName string            `json:"clusterName,omitempty" norman:"type=reference[cluster]"`
	Values      map[string]string `json:"values,omitempty" norman:"required"`
}

type MultiClusterAppRevision struct {
	types.Namespaced
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	TemplateVersionName string            `json:"templateVersionName,omitempty" norman:"type=reference[templateVersion],required"`
	Answers             map[string]string `json:"answers,omitempty"`
}

type MultiClusterAppUpgradeInput struct {
	TemplateVersionName string            `json:"templateVersionName,omitempty"`
	Answers             map[string]string `json:"answers,omitempty"`
	BatchSize           int               `json:"batchSize,omitempty"`
}

type MultiClusterAppRollbackInput struct {
	RevisionName string `json:"revisionName,omitempty" norman:"type=reference[/v3/schemas/multiClusterAppRevision]"`
	BatchSize           int               `json:"batchSize,omitempty"`
}
