package kubernetes

import (
	"testing"

	"github.com/kubernetes/kompose/pkg/kobject"
	autoscalingv2 "k8s.io/api/autoscaling/v2"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestHPAUsesRegisteredStableAPI(t *testing.T) {
	scheme := runtime.NewScheme()
	if err := autoscalingv2.AddToScheme(scheme); err != nil {
		t.Fatal(err)
	}
	service := &kobject.ServiceConfig{}
	object := createHPAResources("web", service)
	kinds, _, err := scheme.ObjectKinds(&object)
	if err != nil {
		t.Fatalf("HPA must use the stable registered Go type: %v", err)
	}
	if len(kinds) != 1 || kinds[0].GroupVersion() != autoscalingv2.SchemeGroupVersion {
		t.Fatalf("unexpected HPA kinds: %v", kinds)
	}
	if object.Spec.ScaleTargetRef.APIVersion != "apps/v1" || object.Spec.ScaleTargetRef.Kind != "Deployment" {
		t.Fatalf("invalid scale target: %+v", object.Spec.ScaleTargetRef)
	}
	if len(object.Spec.Metrics) == 0 {
		t.Fatal("resource metrics must be retained")
	}
}
