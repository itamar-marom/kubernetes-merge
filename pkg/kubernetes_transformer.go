package pkg

import "reflect"

// KubernetesTransformer Define the merge strategy for Kubernetes objects.
type KubernetesTransformer struct{}

func (t KubernetesTransformer) Transformer(typ reflect.Type) func(dst, src reflect.Value) error {

	switch typ {
	}
	return nil
}
