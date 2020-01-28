package v3

import (
	"context"

	"github.com/rancher/norman/controller"
	"github.com/rancher/norman/objectclient"
	"github.com/rancher/norman/resource"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	NodeUpgradeStatusGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "NodeUpgradeStatus",
	}
	NodeUpgradeStatusResource = metav1.APIResource{
		Name:         "nodeupgradestatuses",
		SingularName: "nodeupgradestatus",
		Namespaced:   true,

		Kind: NodeUpgradeStatusGroupVersionKind.Kind,
	}

	NodeUpgradeStatusGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "nodeupgradestatuses",
	}
)

func init() {
	resource.Put(NodeUpgradeStatusGroupVersionResource)
}

func NewNodeUpgradeStatus(namespace, name string, obj NodeUpgradeStatus) *NodeUpgradeStatus {
	obj.APIVersion, obj.Kind = NodeUpgradeStatusGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type NodeUpgradeStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeUpgradeStatus `json:"items"`
}

type NodeUpgradeStatusHandlerFunc func(key string, obj *NodeUpgradeStatus) (runtime.Object, error)

type NodeUpgradeStatusChangeHandlerFunc func(obj *NodeUpgradeStatus) (runtime.Object, error)

type NodeUpgradeStatusLister interface {
	List(namespace string, selector labels.Selector) (ret []*NodeUpgradeStatus, err error)
	Get(namespace, name string) (*NodeUpgradeStatus, error)
}

type NodeUpgradeStatusController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() NodeUpgradeStatusLister
	AddHandler(ctx context.Context, name string, handler NodeUpgradeStatusHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync NodeUpgradeStatusHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler NodeUpgradeStatusHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, handler NodeUpgradeStatusHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type NodeUpgradeStatusInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*NodeUpgradeStatus) (*NodeUpgradeStatus, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*NodeUpgradeStatus, error)
	Get(name string, opts metav1.GetOptions) (*NodeUpgradeStatus, error)
	Update(*NodeUpgradeStatus) (*NodeUpgradeStatus, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*NodeUpgradeStatusList, error)
	ListNamespaced(namespace string, opts metav1.ListOptions) (*NodeUpgradeStatusList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() NodeUpgradeStatusController
	AddHandler(ctx context.Context, name string, sync NodeUpgradeStatusHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync NodeUpgradeStatusHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle NodeUpgradeStatusLifecycle)
	AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle NodeUpgradeStatusLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync NodeUpgradeStatusHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync NodeUpgradeStatusHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle NodeUpgradeStatusLifecycle)
	AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle NodeUpgradeStatusLifecycle)
}

type nodeUpgradeStatusLister struct {
	controller *nodeUpgradeStatusController
}

func (l *nodeUpgradeStatusLister) List(namespace string, selector labels.Selector) (ret []*NodeUpgradeStatus, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*NodeUpgradeStatus))
	})
	return
}

func (l *nodeUpgradeStatusLister) Get(namespace, name string) (*NodeUpgradeStatus, error) {
	var key string
	if namespace != "" {
		key = namespace + "/" + name
	} else {
		key = name
	}
	obj, exists, err := l.controller.Informer().GetIndexer().GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(schema.GroupResource{
			Group:    NodeUpgradeStatusGroupVersionKind.Group,
			Resource: "nodeUpgradeStatus",
		}, key)
	}
	return obj.(*NodeUpgradeStatus), nil
}

type nodeUpgradeStatusController struct {
	controller.GenericController
}

func (c *nodeUpgradeStatusController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *nodeUpgradeStatusController) Lister() NodeUpgradeStatusLister {
	return &nodeUpgradeStatusLister{
		controller: c,
	}
}

func (c *nodeUpgradeStatusController) AddHandler(ctx context.Context, name string, handler NodeUpgradeStatusHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*NodeUpgradeStatus); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *nodeUpgradeStatusController) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, handler NodeUpgradeStatusHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*NodeUpgradeStatus); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *nodeUpgradeStatusController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler NodeUpgradeStatusHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*NodeUpgradeStatus); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *nodeUpgradeStatusController) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, cluster string, handler NodeUpgradeStatusHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*NodeUpgradeStatus); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type nodeUpgradeStatusFactory struct {
}

func (c nodeUpgradeStatusFactory) Object() runtime.Object {
	return &NodeUpgradeStatus{}
}

func (c nodeUpgradeStatusFactory) List() runtime.Object {
	return &NodeUpgradeStatusList{}
}

func (s *nodeUpgradeStatusClient) Controller() NodeUpgradeStatusController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.nodeUpgradeStatusControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(NodeUpgradeStatusGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &nodeUpgradeStatusController{
		GenericController: genericController,
	}

	s.client.nodeUpgradeStatusControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type nodeUpgradeStatusClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   NodeUpgradeStatusController
}

func (s *nodeUpgradeStatusClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *nodeUpgradeStatusClient) Create(o *NodeUpgradeStatus) (*NodeUpgradeStatus, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*NodeUpgradeStatus), err
}

func (s *nodeUpgradeStatusClient) Get(name string, opts metav1.GetOptions) (*NodeUpgradeStatus, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*NodeUpgradeStatus), err
}

func (s *nodeUpgradeStatusClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*NodeUpgradeStatus, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*NodeUpgradeStatus), err
}

func (s *nodeUpgradeStatusClient) Update(o *NodeUpgradeStatus) (*NodeUpgradeStatus, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*NodeUpgradeStatus), err
}

func (s *nodeUpgradeStatusClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *nodeUpgradeStatusClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *nodeUpgradeStatusClient) List(opts metav1.ListOptions) (*NodeUpgradeStatusList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*NodeUpgradeStatusList), err
}

func (s *nodeUpgradeStatusClient) ListNamespaced(namespace string, opts metav1.ListOptions) (*NodeUpgradeStatusList, error) {
	obj, err := s.objectClient.ListNamespaced(namespace, opts)
	return obj.(*NodeUpgradeStatusList), err
}

func (s *nodeUpgradeStatusClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *nodeUpgradeStatusClient) Patch(o *NodeUpgradeStatus, patchType types.PatchType, data []byte, subresources ...string) (*NodeUpgradeStatus, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*NodeUpgradeStatus), err
}

func (s *nodeUpgradeStatusClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *nodeUpgradeStatusClient) AddHandler(ctx context.Context, name string, sync NodeUpgradeStatusHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *nodeUpgradeStatusClient) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync NodeUpgradeStatusHandlerFunc) {
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *nodeUpgradeStatusClient) AddLifecycle(ctx context.Context, name string, lifecycle NodeUpgradeStatusLifecycle) {
	sync := NewNodeUpgradeStatusLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *nodeUpgradeStatusClient) AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle NodeUpgradeStatusLifecycle) {
	sync := NewNodeUpgradeStatusLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *nodeUpgradeStatusClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync NodeUpgradeStatusHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *nodeUpgradeStatusClient) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync NodeUpgradeStatusHandlerFunc) {
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

func (s *nodeUpgradeStatusClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle NodeUpgradeStatusLifecycle) {
	sync := NewNodeUpgradeStatusLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *nodeUpgradeStatusClient) AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle NodeUpgradeStatusLifecycle) {
	sync := NewNodeUpgradeStatusLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

type NodeUpgradeStatusIndexer func(obj *NodeUpgradeStatus) ([]string, error)

type NodeUpgradeStatusClientCache interface {
	Get(namespace, name string) (*NodeUpgradeStatus, error)
	List(namespace string, selector labels.Selector) ([]*NodeUpgradeStatus, error)

	Index(name string, indexer NodeUpgradeStatusIndexer)
	GetIndexed(name, key string) ([]*NodeUpgradeStatus, error)
}

type NodeUpgradeStatusClient interface {
	Create(*NodeUpgradeStatus) (*NodeUpgradeStatus, error)
	Get(namespace, name string, opts metav1.GetOptions) (*NodeUpgradeStatus, error)
	Update(*NodeUpgradeStatus) (*NodeUpgradeStatus, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	List(namespace string, opts metav1.ListOptions) (*NodeUpgradeStatusList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)

	Cache() NodeUpgradeStatusClientCache

	OnCreate(ctx context.Context, name string, sync NodeUpgradeStatusChangeHandlerFunc)
	OnChange(ctx context.Context, name string, sync NodeUpgradeStatusChangeHandlerFunc)
	OnRemove(ctx context.Context, name string, sync NodeUpgradeStatusChangeHandlerFunc)
	Enqueue(namespace, name string)

	Generic() controller.GenericController
	ObjectClient() *objectclient.ObjectClient
	Interface() NodeUpgradeStatusInterface
}

type nodeUpgradeStatusClientCache struct {
	client *nodeUpgradeStatusClient2
}

type nodeUpgradeStatusClient2 struct {
	iface      NodeUpgradeStatusInterface
	controller NodeUpgradeStatusController
}

func (n *nodeUpgradeStatusClient2) Interface() NodeUpgradeStatusInterface {
	return n.iface
}

func (n *nodeUpgradeStatusClient2) Generic() controller.GenericController {
	return n.iface.Controller().Generic()
}

func (n *nodeUpgradeStatusClient2) ObjectClient() *objectclient.ObjectClient {
	return n.Interface().ObjectClient()
}

func (n *nodeUpgradeStatusClient2) Enqueue(namespace, name string) {
	n.iface.Controller().Enqueue(namespace, name)
}

func (n *nodeUpgradeStatusClient2) Create(obj *NodeUpgradeStatus) (*NodeUpgradeStatus, error) {
	return n.iface.Create(obj)
}

func (n *nodeUpgradeStatusClient2) Get(namespace, name string, opts metav1.GetOptions) (*NodeUpgradeStatus, error) {
	return n.iface.GetNamespaced(namespace, name, opts)
}

func (n *nodeUpgradeStatusClient2) Update(obj *NodeUpgradeStatus) (*NodeUpgradeStatus, error) {
	return n.iface.Update(obj)
}

func (n *nodeUpgradeStatusClient2) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return n.iface.DeleteNamespaced(namespace, name, options)
}

func (n *nodeUpgradeStatusClient2) List(namespace string, opts metav1.ListOptions) (*NodeUpgradeStatusList, error) {
	return n.iface.List(opts)
}

func (n *nodeUpgradeStatusClient2) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return n.iface.Watch(opts)
}

func (n *nodeUpgradeStatusClientCache) Get(namespace, name string) (*NodeUpgradeStatus, error) {
	return n.client.controller.Lister().Get(namespace, name)
}

func (n *nodeUpgradeStatusClientCache) List(namespace string, selector labels.Selector) ([]*NodeUpgradeStatus, error) {
	return n.client.controller.Lister().List(namespace, selector)
}

func (n *nodeUpgradeStatusClient2) Cache() NodeUpgradeStatusClientCache {
	n.loadController()
	return &nodeUpgradeStatusClientCache{
		client: n,
	}
}

func (n *nodeUpgradeStatusClient2) OnCreate(ctx context.Context, name string, sync NodeUpgradeStatusChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-create", &nodeUpgradeStatusLifecycleDelegate{create: sync})
}

func (n *nodeUpgradeStatusClient2) OnChange(ctx context.Context, name string, sync NodeUpgradeStatusChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-change", &nodeUpgradeStatusLifecycleDelegate{update: sync})
}

func (n *nodeUpgradeStatusClient2) OnRemove(ctx context.Context, name string, sync NodeUpgradeStatusChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name, &nodeUpgradeStatusLifecycleDelegate{remove: sync})
}

func (n *nodeUpgradeStatusClientCache) Index(name string, indexer NodeUpgradeStatusIndexer) {
	err := n.client.controller.Informer().GetIndexer().AddIndexers(map[string]cache.IndexFunc{
		name: func(obj interface{}) ([]string, error) {
			if v, ok := obj.(*NodeUpgradeStatus); ok {
				return indexer(v)
			}
			return nil, nil
		},
	})

	if err != nil {
		panic(err)
	}
}

func (n *nodeUpgradeStatusClientCache) GetIndexed(name, key string) ([]*NodeUpgradeStatus, error) {
	var result []*NodeUpgradeStatus
	objs, err := n.client.controller.Informer().GetIndexer().ByIndex(name, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		if v, ok := obj.(*NodeUpgradeStatus); ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (n *nodeUpgradeStatusClient2) loadController() {
	if n.controller == nil {
		n.controller = n.iface.Controller()
	}
}

type nodeUpgradeStatusLifecycleDelegate struct {
	create NodeUpgradeStatusChangeHandlerFunc
	update NodeUpgradeStatusChangeHandlerFunc
	remove NodeUpgradeStatusChangeHandlerFunc
}

func (n *nodeUpgradeStatusLifecycleDelegate) HasCreate() bool {
	return n.create != nil
}

func (n *nodeUpgradeStatusLifecycleDelegate) Create(obj *NodeUpgradeStatus) (runtime.Object, error) {
	if n.create == nil {
		return obj, nil
	}
	return n.create(obj)
}

func (n *nodeUpgradeStatusLifecycleDelegate) HasFinalize() bool {
	return n.remove != nil
}

func (n *nodeUpgradeStatusLifecycleDelegate) Remove(obj *NodeUpgradeStatus) (runtime.Object, error) {
	if n.remove == nil {
		return obj, nil
	}
	return n.remove(obj)
}

func (n *nodeUpgradeStatusLifecycleDelegate) Updated(obj *NodeUpgradeStatus) (runtime.Object, error) {
	if n.update == nil {
		return obj, nil
	}
	return n.update(obj)
}
