package v1

import (
	"context"

	"github.com/rancher/norman/clientbase"
	"github.com/rancher/norman/controller"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	SecretGroupVersionKind = schema.GroupVersionKind{
		Version: "v1",
		Group:   "",
		Kind:    "Secret",
	}
	SecretResource = metav1.APIResource{
		Name:         "secrets",
		SingularName: "secret",
		Namespaced:   false,
		Kind:         SecretGroupVersionKind.Kind,
	}
)

type SecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []v1.Secret
}

type SecretHandlerFunc func(key string, obj *v1.Secret) error

type SecretLister interface {
	List(namespace string, selector labels.Selector) (ret []*v1.Secret, err error)
	Get(namespace, name string) (*v1.Secret, error)
}

type SecretController interface {
	Informer() cache.SharedIndexInformer
	Lister() SecretLister
	AddHandler(handler SecretHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type SecretInterface interface {
	ObjectClient() *clientbase.ObjectClient
	Create(*v1.Secret) (*v1.Secret, error)
	Get(name string, opts metav1.GetOptions) (*v1.Secret, error)
	Update(*v1.Secret) (*v1.Secret, error)
	Delete(name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*SecretList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() SecretController
}

type secretLister struct {
	controller *secretController
}

func (l *secretLister) List(namespace string, selector labels.Selector) (ret []*v1.Secret, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*v1.Secret))
	})
	return
}

func (l *secretLister) Get(namespace, name string) (*v1.Secret, error) {
	obj, exists, err := l.controller.Informer().GetIndexer().GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(schema.GroupResource{
			Group:    SecretGroupVersionKind.Group,
			Resource: "secret",
		}, name)
	}
	return obj.(*v1.Secret), nil
}

type secretController struct {
	controller.GenericController
}

func (c *secretController) Lister() SecretLister {
	return &secretLister{
		controller: c,
	}
}

func (c *secretController) AddHandler(handler SecretHandlerFunc) {
	c.GenericController.AddHandler(func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*v1.Secret))
	})
}

type secretFactory struct {
}

func (c secretFactory) Object() runtime.Object {
	return &v1.Secret{}
}

func (c secretFactory) List() runtime.Object {
	return &SecretList{}
}

func (s *secretClient) Controller() SecretController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.secretControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(SecretGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &secretController{
		GenericController: genericController,
	}

	s.client.secretControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type secretClient struct {
	client       *Client
	ns           string
	objectClient *clientbase.ObjectClient
	controller   SecretController
}

func (s *secretClient) ObjectClient() *clientbase.ObjectClient {
	return s.objectClient
}

func (s *secretClient) Create(o *v1.Secret) (*v1.Secret, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*v1.Secret), err
}

func (s *secretClient) Get(name string, opts metav1.GetOptions) (*v1.Secret, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*v1.Secret), err
}

func (s *secretClient) Update(o *v1.Secret) (*v1.Secret, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*v1.Secret), err
}

func (s *secretClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *secretClient) List(opts metav1.ListOptions) (*SecretList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*SecretList), err
}

func (s *secretClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

func (s *secretClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}
