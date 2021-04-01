// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "github.com/openshift/cluster-kube-descheduler-operator/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// KubeDeschedulers returns a KubeDeschedulerInformer.
	KubeDeschedulers() KubeDeschedulerInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// KubeDeschedulers returns a KubeDeschedulerInformer.
func (v *version) KubeDeschedulers() KubeDeschedulerInformer {
	return &kubeDeschedulerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
