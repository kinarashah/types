package client

import (
	"github.com/rancher/norman/types"
)

const (
	CloudCredentialType                 = "cloudCredential"
	CloudCredentialFieldAnnotations     = "annotations"
	CloudCredentialFieldCreated         = "created"
	CloudCredentialFieldCreatorID       = "creatorId"
	CloudCredentialFieldLabels          = "labels"
	CloudCredentialFieldName            = "name"
	CloudCredentialFieldNamespaceId     = "namespaceId"
	CloudCredentialFieldOwnerReferences = "ownerReferences"
	CloudCredentialFieldRemoved         = "removed"
	CloudCredentialFieldUUID            = "uuid"
)

type CloudCredential struct {
	types.Resource
	Annotations     map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created         string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID       string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Labels          map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name            string            `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId     string            `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"`
	OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed         string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	UUID            string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type CloudCredentialCollection struct {
	types.Collection
	Data   []CloudCredential `json:"data,omitempty"`
	client *CloudCredentialClient
}

type CloudCredentialClient struct {
	apiClient *Client
}

type CloudCredentialOperations interface {
	List(opts *types.ListOpts) (*CloudCredentialCollection, error)
	Create(opts *CloudCredential) (*CloudCredential, error)
	Update(existing *CloudCredential, updates interface{}) (*CloudCredential, error)
	Replace(existing *CloudCredential) (*CloudCredential, error)
	ByID(id string) (*CloudCredential, error)
	Delete(container *CloudCredential) error
}

func newCloudCredentialClient(apiClient *Client) *CloudCredentialClient {
	return &CloudCredentialClient{
		apiClient: apiClient,
	}
}

func (c *CloudCredentialClient) Create(container *CloudCredential) (*CloudCredential, error) {
	resp := &CloudCredential{}
	err := c.apiClient.Ops.DoCreate(CloudCredentialType, container, resp)
	return resp, err
}

func (c *CloudCredentialClient) Update(existing *CloudCredential, updates interface{}) (*CloudCredential, error) {
	resp := &CloudCredential{}
	err := c.apiClient.Ops.DoUpdate(CloudCredentialType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *CloudCredentialClient) Replace(obj *CloudCredential) (*CloudCredential, error) {
	resp := &CloudCredential{}
	err := c.apiClient.Ops.DoReplace(CloudCredentialType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *CloudCredentialClient) List(opts *types.ListOpts) (*CloudCredentialCollection, error) {
	resp := &CloudCredentialCollection{}
	err := c.apiClient.Ops.DoList(CloudCredentialType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *CloudCredentialCollection) Next() (*CloudCredentialCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &CloudCredentialCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *CloudCredentialClient) ByID(id string) (*CloudCredential, error) {
	resp := &CloudCredential{}
	err := c.apiClient.Ops.DoByID(CloudCredentialType, id, resp)
	return resp, err
}

func (c *CloudCredentialClient) Delete(container *CloudCredential) error {
	return c.apiClient.Ops.DoResourceDelete(CloudCredentialType, &container.Resource)
}
