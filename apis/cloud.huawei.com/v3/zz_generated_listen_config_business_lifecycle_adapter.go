package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type ListenConfigBusinessLifecycle interface {
	Create(obj *ListenConfigBusiness) (*ListenConfigBusiness, error)
	Remove(obj *ListenConfigBusiness) (*ListenConfigBusiness, error)
	Updated(obj *ListenConfigBusiness) (*ListenConfigBusiness, error)
}

type listenConfigBusinessLifecycleAdapter struct {
	lifecycle ListenConfigBusinessLifecycle
}

func (w *listenConfigBusinessLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*ListenConfigBusiness))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *listenConfigBusinessLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*ListenConfigBusiness))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *listenConfigBusinessLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*ListenConfigBusiness))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewListenConfigBusinessLifecycleAdapter(name string, clusterScoped bool, client ListenConfigBusinessInterface, l ListenConfigBusinessLifecycle) ListenConfigBusinessHandlerFunc {
	adapter := &listenConfigBusinessLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *ListenConfigBusiness) error {
		if obj == nil {
			return syncFn(key, nil)
		}
		return syncFn(key, obj)
	}
}
