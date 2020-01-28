package v3

import (
	"github.com/rancher/norman/lifecycle"
	"github.com/rancher/norman/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

type NodeUpgradeStatusLifecycle interface {
	Create(obj *NodeUpgradeStatus) (runtime.Object, error)
	Remove(obj *NodeUpgradeStatus) (runtime.Object, error)
	Updated(obj *NodeUpgradeStatus) (runtime.Object, error)
}

type nodeUpgradeStatusLifecycleAdapter struct {
	lifecycle NodeUpgradeStatusLifecycle
}

func (w *nodeUpgradeStatusLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *nodeUpgradeStatusLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *nodeUpgradeStatusLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*NodeUpgradeStatus))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *nodeUpgradeStatusLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*NodeUpgradeStatus))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *nodeUpgradeStatusLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*NodeUpgradeStatus))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewNodeUpgradeStatusLifecycleAdapter(name string, clusterScoped bool, client NodeUpgradeStatusInterface, l NodeUpgradeStatusLifecycle) NodeUpgradeStatusHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(NodeUpgradeStatusGroupVersionResource)
	}
	adapter := &nodeUpgradeStatusLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *NodeUpgradeStatus) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
