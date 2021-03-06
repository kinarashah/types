package client

import (
	"github.com/rancher/norman/types"
)

const (
	ProjectRoleTemplateBindingType                       = "projectRoleTemplateBinding"
	ProjectRoleTemplateBindingFieldAnnotations           = "annotations"
	ProjectRoleTemplateBindingFieldCreated               = "created"
	ProjectRoleTemplateBindingFieldLabels                = "labels"
	ProjectRoleTemplateBindingFieldName                  = "name"
	ProjectRoleTemplateBindingFieldOwnerReferences       = "ownerReferences"
	ProjectRoleTemplateBindingFieldProjectId             = "projectId"
	ProjectRoleTemplateBindingFieldProjectRoleTemplateId = "projectRoleTemplateId"
	ProjectRoleTemplateBindingFieldRemoved               = "removed"
	ProjectRoleTemplateBindingFieldResourcePath          = "resourcePath"
	ProjectRoleTemplateBindingFieldState                 = "state"
	ProjectRoleTemplateBindingFieldSubjectKind           = "subjectKind"
	ProjectRoleTemplateBindingFieldSubjectName           = "subjectName"
	ProjectRoleTemplateBindingFieldSubjectNamespace      = "subjectNamespace"
	ProjectRoleTemplateBindingFieldTransitioning         = "transitioning"
	ProjectRoleTemplateBindingFieldTransitioningMessage  = "transitioningMessage"
	ProjectRoleTemplateBindingFieldUuid                  = "uuid"
)

type ProjectRoleTemplateBinding struct {
	types.Resource
	Annotations           map[string]string `json:"annotations,omitempty"`
	Created               string            `json:"created,omitempty"`
	Labels                map[string]string `json:"labels,omitempty"`
	Name                  string            `json:"name,omitempty"`
	OwnerReferences       []OwnerReference  `json:"ownerReferences,omitempty"`
	ProjectId             string            `json:"projectId,omitempty"`
	ProjectRoleTemplateId string            `json:"projectRoleTemplateId,omitempty"`
	Removed               string            `json:"removed,omitempty"`
	ResourcePath          string            `json:"resourcePath,omitempty"`
	State                 string            `json:"state,omitempty"`
	SubjectKind           string            `json:"subjectKind,omitempty"`
	SubjectName           string            `json:"subjectName,omitempty"`
	SubjectNamespace      string            `json:"subjectNamespace,omitempty"`
	Transitioning         string            `json:"transitioning,omitempty"`
	TransitioningMessage  string            `json:"transitioningMessage,omitempty"`
	Uuid                  string            `json:"uuid,omitempty"`
}
type ProjectRoleTemplateBindingCollection struct {
	types.Collection
	Data   []ProjectRoleTemplateBinding `json:"data,omitempty"`
	client *ProjectRoleTemplateBindingClient
}

type ProjectRoleTemplateBindingClient struct {
	apiClient *Client
}

type ProjectRoleTemplateBindingOperations interface {
	List(opts *types.ListOpts) (*ProjectRoleTemplateBindingCollection, error)
	Create(opts *ProjectRoleTemplateBinding) (*ProjectRoleTemplateBinding, error)
	Update(existing *ProjectRoleTemplateBinding, updates interface{}) (*ProjectRoleTemplateBinding, error)
	ByID(id string) (*ProjectRoleTemplateBinding, error)
	Delete(container *ProjectRoleTemplateBinding) error
}

func newProjectRoleTemplateBindingClient(apiClient *Client) *ProjectRoleTemplateBindingClient {
	return &ProjectRoleTemplateBindingClient{
		apiClient: apiClient,
	}
}

func (c *ProjectRoleTemplateBindingClient) Create(container *ProjectRoleTemplateBinding) (*ProjectRoleTemplateBinding, error) {
	resp := &ProjectRoleTemplateBinding{}
	err := c.apiClient.Ops.DoCreate(ProjectRoleTemplateBindingType, container, resp)
	return resp, err
}

func (c *ProjectRoleTemplateBindingClient) Update(existing *ProjectRoleTemplateBinding, updates interface{}) (*ProjectRoleTemplateBinding, error) {
	resp := &ProjectRoleTemplateBinding{}
	err := c.apiClient.Ops.DoUpdate(ProjectRoleTemplateBindingType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ProjectRoleTemplateBindingClient) List(opts *types.ListOpts) (*ProjectRoleTemplateBindingCollection, error) {
	resp := &ProjectRoleTemplateBindingCollection{}
	err := c.apiClient.Ops.DoList(ProjectRoleTemplateBindingType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ProjectRoleTemplateBindingCollection) Next() (*ProjectRoleTemplateBindingCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ProjectRoleTemplateBindingCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ProjectRoleTemplateBindingClient) ByID(id string) (*ProjectRoleTemplateBinding, error) {
	resp := &ProjectRoleTemplateBinding{}
	err := c.apiClient.Ops.DoByID(ProjectRoleTemplateBindingType, id, resp)
	return resp, err
}

func (c *ProjectRoleTemplateBindingClient) Delete(container *ProjectRoleTemplateBinding) error {
	return c.apiClient.Ops.DoResourceDelete(ProjectRoleTemplateBindingType, &container.Resource)
}
