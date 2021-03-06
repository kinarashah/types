package client

import (
	"github.com/rancher/norman/types"
)

const (
	ClusterRoleTemplateType                        = "clusterRoleTemplate"
	ClusterRoleTemplateFieldAnnotations            = "annotations"
	ClusterRoleTemplateFieldBuiltin                = "builtin"
	ClusterRoleTemplateFieldClusterRoleTemplateIds = "clusterRoleTemplateIds"
	ClusterRoleTemplateFieldCreated                = "created"
	ClusterRoleTemplateFieldLabels                 = "labels"
	ClusterRoleTemplateFieldName                   = "name"
	ClusterRoleTemplateFieldOwnerReferences        = "ownerReferences"
	ClusterRoleTemplateFieldRemoved                = "removed"
	ClusterRoleTemplateFieldResourcePath           = "resourcePath"
	ClusterRoleTemplateFieldRules                  = "rules"
	ClusterRoleTemplateFieldState                  = "state"
	ClusterRoleTemplateFieldTransitioning          = "transitioning"
	ClusterRoleTemplateFieldTransitioningMessage   = "transitioningMessage"
	ClusterRoleTemplateFieldUuid                   = "uuid"
)

type ClusterRoleTemplate struct {
	types.Resource
	Annotations            map[string]string `json:"annotations,omitempty"`
	Builtin                *bool             `json:"builtin,omitempty"`
	ClusterRoleTemplateIds []string          `json:"clusterRoleTemplateIds,omitempty"`
	Created                string            `json:"created,omitempty"`
	Labels                 map[string]string `json:"labels,omitempty"`
	Name                   string            `json:"name,omitempty"`
	OwnerReferences        []OwnerReference  `json:"ownerReferences,omitempty"`
	Removed                string            `json:"removed,omitempty"`
	ResourcePath           string            `json:"resourcePath,omitempty"`
	Rules                  []PolicyRule      `json:"rules,omitempty"`
	State                  string            `json:"state,omitempty"`
	Transitioning          string            `json:"transitioning,omitempty"`
	TransitioningMessage   string            `json:"transitioningMessage,omitempty"`
	Uuid                   string            `json:"uuid,omitempty"`
}
type ClusterRoleTemplateCollection struct {
	types.Collection
	Data   []ClusterRoleTemplate `json:"data,omitempty"`
	client *ClusterRoleTemplateClient
}

type ClusterRoleTemplateClient struct {
	apiClient *Client
}

type ClusterRoleTemplateOperations interface {
	List(opts *types.ListOpts) (*ClusterRoleTemplateCollection, error)
	Create(opts *ClusterRoleTemplate) (*ClusterRoleTemplate, error)
	Update(existing *ClusterRoleTemplate, updates interface{}) (*ClusterRoleTemplate, error)
	ByID(id string) (*ClusterRoleTemplate, error)
	Delete(container *ClusterRoleTemplate) error
}

func newClusterRoleTemplateClient(apiClient *Client) *ClusterRoleTemplateClient {
	return &ClusterRoleTemplateClient{
		apiClient: apiClient,
	}
}

func (c *ClusterRoleTemplateClient) Create(container *ClusterRoleTemplate) (*ClusterRoleTemplate, error) {
	resp := &ClusterRoleTemplate{}
	err := c.apiClient.Ops.DoCreate(ClusterRoleTemplateType, container, resp)
	return resp, err
}

func (c *ClusterRoleTemplateClient) Update(existing *ClusterRoleTemplate, updates interface{}) (*ClusterRoleTemplate, error) {
	resp := &ClusterRoleTemplate{}
	err := c.apiClient.Ops.DoUpdate(ClusterRoleTemplateType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ClusterRoleTemplateClient) List(opts *types.ListOpts) (*ClusterRoleTemplateCollection, error) {
	resp := &ClusterRoleTemplateCollection{}
	err := c.apiClient.Ops.DoList(ClusterRoleTemplateType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ClusterRoleTemplateCollection) Next() (*ClusterRoleTemplateCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ClusterRoleTemplateCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ClusterRoleTemplateClient) ByID(id string) (*ClusterRoleTemplate, error) {
	resp := &ClusterRoleTemplate{}
	err := c.apiClient.Ops.DoByID(ClusterRoleTemplateType, id, resp)
	return resp, err
}

func (c *ClusterRoleTemplateClient) Delete(container *ClusterRoleTemplate) error {
	return c.apiClient.Ops.DoResourceDelete(ClusterRoleTemplateType, &container.Resource)
}
