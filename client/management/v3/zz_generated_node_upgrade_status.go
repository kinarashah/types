package client

import (
	"github.com/rancher/norman/types"
)

const (
	NodeUpgradeStatusType                  = "nodeUpgradeStatus"
	NodeUpgradeStatusFieldAnnotations      = "annotations"
	NodeUpgradeStatusFieldCreated          = "created"
	NodeUpgradeStatusFieldCreatorID        = "creatorId"
	NodeUpgradeStatusFieldCurrentToken     = "currentToken"
	NodeUpgradeStatusFieldLabels           = "labels"
	NodeUpgradeStatusFieldLastAppliedToken = "lastAppliedToken"
	NodeUpgradeStatusFieldName             = "name"
	NodeUpgradeStatusFieldNamespaceId      = "namespaceId"
	NodeUpgradeStatusFieldNodes            = "nodes"
	NodeUpgradeStatusFieldOwnerReferences  = "ownerReferences"
	NodeUpgradeStatusFieldRemoved          = "removed"
	NodeUpgradeStatusFieldUUID             = "uuid"
)

type NodeUpgradeStatus struct {
	types.Resource
	Annotations      map[string]string            `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created          string                       `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID        string                       `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	CurrentToken     string                       `json:"currentToken,omitempty" yaml:"currentToken,omitempty"`
	Labels           map[string]string            `json:"labels,omitempty" yaml:"labels,omitempty"`
	LastAppliedToken string                       `json:"lastAppliedToken,omitempty" yaml:"lastAppliedToken,omitempty"`
	Name             string                       `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId      string                       `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"`
	Nodes            map[string]map[string]string `json:"nodes,omitempty" yaml:"nodes,omitempty"`
	OwnerReferences  []OwnerReference             `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed          string                       `json:"removed,omitempty" yaml:"removed,omitempty"`
	UUID             string                       `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type NodeUpgradeStatusCollection struct {
	types.Collection
	Data   []NodeUpgradeStatus `json:"data,omitempty"`
	client *NodeUpgradeStatusClient
}

type NodeUpgradeStatusClient struct {
	apiClient *Client
}

type NodeUpgradeStatusOperations interface {
	List(opts *types.ListOpts) (*NodeUpgradeStatusCollection, error)
	Create(opts *NodeUpgradeStatus) (*NodeUpgradeStatus, error)
	Update(existing *NodeUpgradeStatus, updates interface{}) (*NodeUpgradeStatus, error)
	Replace(existing *NodeUpgradeStatus) (*NodeUpgradeStatus, error)
	ByID(id string) (*NodeUpgradeStatus, error)
	Delete(container *NodeUpgradeStatus) error
}

func newNodeUpgradeStatusClient(apiClient *Client) *NodeUpgradeStatusClient {
	return &NodeUpgradeStatusClient{
		apiClient: apiClient,
	}
}

func (c *NodeUpgradeStatusClient) Create(container *NodeUpgradeStatus) (*NodeUpgradeStatus, error) {
	resp := &NodeUpgradeStatus{}
	err := c.apiClient.Ops.DoCreate(NodeUpgradeStatusType, container, resp)
	return resp, err
}

func (c *NodeUpgradeStatusClient) Update(existing *NodeUpgradeStatus, updates interface{}) (*NodeUpgradeStatus, error) {
	resp := &NodeUpgradeStatus{}
	err := c.apiClient.Ops.DoUpdate(NodeUpgradeStatusType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *NodeUpgradeStatusClient) Replace(obj *NodeUpgradeStatus) (*NodeUpgradeStatus, error) {
	resp := &NodeUpgradeStatus{}
	err := c.apiClient.Ops.DoReplace(NodeUpgradeStatusType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *NodeUpgradeStatusClient) List(opts *types.ListOpts) (*NodeUpgradeStatusCollection, error) {
	resp := &NodeUpgradeStatusCollection{}
	err := c.apiClient.Ops.DoList(NodeUpgradeStatusType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *NodeUpgradeStatusCollection) Next() (*NodeUpgradeStatusCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &NodeUpgradeStatusCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *NodeUpgradeStatusClient) ByID(id string) (*NodeUpgradeStatus, error) {
	resp := &NodeUpgradeStatus{}
	err := c.apiClient.Ops.DoByID(NodeUpgradeStatusType, id, resp)
	return resp, err
}

func (c *NodeUpgradeStatusClient) Delete(container *NodeUpgradeStatus) error {
	return c.apiClient.Ops.DoResourceDelete(NodeUpgradeStatusType, &container.Resource)
}
