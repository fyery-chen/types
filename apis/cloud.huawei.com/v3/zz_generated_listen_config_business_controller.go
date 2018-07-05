package v3

import (
	"context"

	"github.com/rancher/norman/controller"
	"github.com/rancher/norman/objectclient"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	ListenConfigBusinessGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "ListenConfigBusiness",
	}
	ListenConfigBusinessResource = metav1.APIResource{
		Name:         "listenconfigbusinesses",
		SingularName: "listenconfigbusiness",
		Namespaced:   false,
		Kind:         ListenConfigBusinessGroupVersionKind.Kind,
	}
)

type ListenConfigBusinessList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ListenConfigBusiness
}

type ListenConfigBusinessHandlerFunc func(key string, obj *ListenConfigBusiness) error

type ListenConfigBusinessLister interface {
	List(namespace string, selector labels.Selector) (ret []*ListenConfigBusiness, err error)
	Get(namespace, name string) (*ListenConfigBusiness, error)
}

type ListenConfigBusinessController interface {
	Informer() cache.SharedIndexInformer
	Lister() ListenConfigBusinessLister
	AddHandler(name string, handler ListenConfigBusinessHandlerFunc)
	AddClusterScopedHandler(name, clusterName string, handler ListenConfigBusinessHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type ListenConfigBusinessInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*ListenConfigBusiness) (*ListenConfigBusiness, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*ListenConfigBusiness, error)
	Get(name string, opts metav1.GetOptions) (*ListenConfigBusiness, error)
	Update(*ListenConfigBusiness) (*ListenConfigBusiness, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*ListenConfigBusinessList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() ListenConfigBusinessController
	AddHandler(name string, sync ListenConfigBusinessHandlerFunc)
	AddLifecycle(name string, lifecycle ListenConfigBusinessLifecycle)
	AddClusterScopedHandler(name, clusterName string, sync ListenConfigBusinessHandlerFunc)
	AddClusterScopedLifecycle(name, clusterName string, lifecycle ListenConfigBusinessLifecycle)
}

type listenConfigBusinessLister struct {
	controller *listenConfigBusinessController
}

func (l *listenConfigBusinessLister) List(namespace string, selector labels.Selector) (ret []*ListenConfigBusiness, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*ListenConfigBusiness))
	})
	return
}

func (l *listenConfigBusinessLister) Get(namespace, name string) (*ListenConfigBusiness, error) {
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
			Group:    ListenConfigBusinessGroupVersionKind.Group,
			Resource: "listenConfigBusiness",
		}, name)
	}
	return obj.(*ListenConfigBusiness), nil
}

type listenConfigBusinessController struct {
	controller.GenericController
}

func (c *listenConfigBusinessController) Lister() ListenConfigBusinessLister {
	return &listenConfigBusinessLister{
		controller: c,
	}
}

func (c *listenConfigBusinessController) AddHandler(name string, handler ListenConfigBusinessHandlerFunc) {
	c.GenericController.AddHandler(name, func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*ListenConfigBusiness))
	})
}

func (c *listenConfigBusinessController) AddClusterScopedHandler(name, cluster string, handler ListenConfigBusinessHandlerFunc) {
	c.GenericController.AddHandler(name, func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}

		if !controller.ObjectInCluster(cluster, obj) {
			return nil
		}

		return handler(key, obj.(*ListenConfigBusiness))
	})
}

type listenConfigBusinessFactory struct {
}

func (c listenConfigBusinessFactory) Object() runtime.Object {
	return &ListenConfigBusiness{}
}

func (c listenConfigBusinessFactory) List() runtime.Object {
	return &ListenConfigBusinessList{}
}

func (s *listenConfigBusinessClient) Controller() ListenConfigBusinessController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.listenConfigBusinessControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(ListenConfigBusinessGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &listenConfigBusinessController{
		GenericController: genericController,
	}

	s.client.listenConfigBusinessControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type listenConfigBusinessClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   ListenConfigBusinessController
}

func (s *listenConfigBusinessClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *listenConfigBusinessClient) Create(o *ListenConfigBusiness) (*ListenConfigBusiness, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*ListenConfigBusiness), err
}

func (s *listenConfigBusinessClient) Get(name string, opts metav1.GetOptions) (*ListenConfigBusiness, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*ListenConfigBusiness), err
}

func (s *listenConfigBusinessClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*ListenConfigBusiness, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*ListenConfigBusiness), err
}

func (s *listenConfigBusinessClient) Update(o *ListenConfigBusiness) (*ListenConfigBusiness, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*ListenConfigBusiness), err
}

func (s *listenConfigBusinessClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *listenConfigBusinessClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *listenConfigBusinessClient) List(opts metav1.ListOptions) (*ListenConfigBusinessList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*ListenConfigBusinessList), err
}

func (s *listenConfigBusinessClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *listenConfigBusinessClient) Patch(o *ListenConfigBusiness, data []byte, subresources ...string) (*ListenConfigBusiness, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*ListenConfigBusiness), err
}

func (s *listenConfigBusinessClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *listenConfigBusinessClient) AddHandler(name string, sync ListenConfigBusinessHandlerFunc) {
	s.Controller().AddHandler(name, sync)
}

func (s *listenConfigBusinessClient) AddLifecycle(name string, lifecycle ListenConfigBusinessLifecycle) {
	sync := NewListenConfigBusinessLifecycleAdapter(name, false, s, lifecycle)
	s.AddHandler(name, sync)
}

func (s *listenConfigBusinessClient) AddClusterScopedHandler(name, clusterName string, sync ListenConfigBusinessHandlerFunc) {
	s.Controller().AddClusterScopedHandler(name, clusterName, sync)
}

func (s *listenConfigBusinessClient) AddClusterScopedLifecycle(name, clusterName string, lifecycle ListenConfigBusinessLifecycle) {
	sync := NewListenConfigBusinessLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.AddClusterScopedHandler(name, clusterName, sync)
}
