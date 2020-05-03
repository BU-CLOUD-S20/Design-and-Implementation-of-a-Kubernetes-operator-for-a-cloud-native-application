package e2e

import(
    "context"
    "time"
    // "math/rand"
    // "reflect"
    // "strconv"
    "testing"

    cachev1alpha1 "demo1/opdemo/pkg/apis/app/v1" //edit this


    // corev1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    appsv1 "k8s.io/api/apps/v1"

    // framework "github.com/operator-framework/operator-sdk/pkg/test"
    // "github.com/operator-framework/operator-sdk/pkg/test/e2eutil"

    "k8s.io/apimachinery/pkg/runtime"
    "k8s.io/apimachinery/pkg/types"
    "k8s.io/client-go/kubernetes/scheme"
    // "sigs.k8s.io/controller-runtime/pkg/client"
    "sigs.k8s.io/controller-runtime/pkg/client/fake"
    "sigs.k8s.io/controller-runtime/pkg/reconcile"

    logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var(
	retryInterval        = time.Second * 5
	timeout              = time.Second * 60
	cleanupRetryInterval = time.Second * 1
	cleanupTimeout       = time.Second * 5

)

func testAppserviceCreate(t *testing.T) {
	logf.SetLogger(logf.ZapLogger(true))
 var (
        name            = "opdemo"
        namespace       = "appservice"
        replicas  int32 = 1
    )
	appService := &cachev1alpha1.AppService{
        ObjectMeta: metav1.ObjectMeta{
            Name:      name,
            Namespace: namespace,
        },
        Spec: cachev1alpha1.AppServiceSpec{
            Size: replicas, // Set desired number of replicas.
        },
    }

    // Objects to track in the fake client.
    objs := []runtime.Object{ appService }

    // Register operator types with the runtime scheme.
    s := scheme.Scheme
    s.AddKnownTypes(cachev1alpha1.SchemeGroupVersion, appService)

    // Create a fake client to mock API calls.
    cl := fake.NewFakeClient(objs...)

    // Create a ReconcileAppservice object with the scheme and fake client.
    r := &ReconcileAppservice{client: cl, scheme: s}

    // Mock request to simulate Reconcile() being called on an event for a
    // watched resource .
    req := reconcile.Request{
        NamespacedName: types.NamespacedName{
            Name:      name,
            Namespace: namespace,
        },
    }
    res, err := r.Reconcile(req)
    if err != nil {
       t.Fatalf("reconcile: (%v)", err)
    }
    // Check the result of reconciliation to make sure it has the desired state.
    if !res.Requeue {
        t.Error("reconcile did not requeue request as expected")
    }
    // Check if deployment has been created and has the correct size.
    dep := &appsv1.Deployment{}
    err = r.client.Get(context.TODO(), req.NamespacedName, dep)
    if err != nil {
        t.Fatalf("get deployment: (%v)", err)
    }
    // Check if the quantity of Replicas for this deployment is equals the specification
    dsize := *dep.Spec.Replicas
    if dsize != replicas {
        t.Errorf("dep size (%d) is not the expected size (%d)", dsize, replicas)
    }

}
