/*
Copyright 2019 geovis.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package headlesskind

import (
	"context"
	fwzxv1beta1 "crdDemo/pkg/apis/fwzx/v1beta1"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"reflect"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new HeadlessKind Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileHeadlessKind{
		Client: mgr.GetClient(),
		scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("headlesskind-controller",
		mgr,
		controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to HeadlessKind
	err = c.Watch(
		&source.Kind{Type: &fwzxv1beta1.HeadlessKind{}},
		&handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// 监控deployment
	err = c.Watch(
		&source.Kind{Type: &appsv1.Deployment{}},
		&handler.EnqueueRequestForOwner{
			IsController: true,
			OwnerType:    &fwzxv1beta1.HeadlessKind{},
		})
	if err != nil {
		return err
	}

	// 监控deployment
	err = c.Watch(
		&source.Kind{Type: &appsv1.Deployment{}},
		&handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	err = c.Watch(
		&source.Kind{Type: &corev1.Pod{}},
		&handler.EnqueueRequestForOwner{
			IsController: true,
			OwnerType:    &appsv1.ReplicaSet{},
		})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileHeadlessKind{}

// ReconcileHeadlessKind reconciles a HeadlessKind object
type ReconcileHeadlessKind struct {
	client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a HeadlessKind object and makes changes based on the state read
// and what is in the HeadlessKind.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  The scaffolding writes
// a Deployment as an example
// Automatically generate RBAC rules to allow the Controller to read and write Deployments
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,resources=deployments/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=fwzx.geovis.ai,resources=headlesskinds,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=fwzx.geovis.ai,resources=headlesskinds/status,verbs=get;update;patch
func (r *ReconcileHeadlessKind) Reconcile_old(request reconcile.Request) (reconcile.Result, error) {
	// Fetch the HeadlessKind instance
	instance := &fwzxv1beta1.HeadlessKind{}
	err := r.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// TODO(user): Change this to be the object type created by your controller
	// Define the desired Deployment object
	i := int32(1)
	deploy := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.Name + "-deployment",
			Namespace: instance.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &i,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"deployment": instance.Name + "-deployment"},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"deployment": instance.Name + "-deployment"}},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "nginx",
							Image: "nginx",
						},
					},
				},
			},
		},
	}
	if err := controllerutil.SetControllerReference(instance, deploy, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// TODO(user): Change this for the object type created by your controller
	// Check if the Deployment already exists
	found := &appsv1.Deployment{}

	fmt.Println("instance.Spec.Hope", instance.Spec.Hope)
	fmt.Println("instance.Spec.Hope", instance.Spec.Hope)
	fmt.Println("instance.Spec.Hope", instance.Spec.Hope)
	fmt.Println("instance.Spec.Hope", instance.Spec.Hope)

	err = r.Get(context.TODO(),
		types.NamespacedName{Name: deploy.Name, Namespace: deploy.Namespace},
		found)
	if err != nil && errors.IsNotFound(err) {
		log.Info("Creating Deployment", "namespace", deploy.Namespace, "name", deploy.Name)
		err = r.Create(context.TODO(), deploy)
		return reconcile.Result{}, err
	} else if err != nil {
		return reconcile.Result{}, err
	}

	var podList corev1.PodList
	err = r.List(context.TODO(), &client.ListOptions{}, &podList)
	if err != nil {
		return reconcile.Result{}, err
	}

	for _, v := range podList.Items {
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println(v.Kind, v.Name, v.Namespace)
		for _, cs := range v.Status.ContainerStatuses {
			if cs.Ready == true {
				fmt.Println("没有ready的话删除pod。删除svc")
			}
		}
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
	}

	// TODO(user): Change this for the object type created by your controller
	// Update the found object and write the result back if there are any changes
	if !reflect.DeepEqual(deploy.Spec, found.Spec) {
		found.Spec = deploy.Spec
		log.Info("Updating Deployment", "namespace", deploy.Namespace, "name", deploy.Name)
		err = r.Update(context.TODO(), found)
		if err != nil {
			return reconcile.Result{}, err
		}
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileHeadlessKind) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	fmt.Println("working.......Reconcile(request reconcile.Request)")
	instance := &fwzxv1beta1.HeadlessKind{}
	err := r.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}

	i := int32(1)
	deploy := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.Name + "-deployment",
			Namespace: instance.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &i,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"deployment": instance.Name + "-deployment"},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"deployment": instance.Name + "-deployment",
						"x":          "y",
					}},

				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "nginx",
							Image: "nginx:alpine",

							ReadinessProbe: &corev1.Probe{
								Handler: corev1.Handler{

									HTTPGet: &corev1.HTTPGetAction{
										Path: "/",
										Port: intstr.IntOrString{IntVal: 80},
									},
								},
								InitialDelaySeconds: 15,
								TimeoutSeconds:      5,
								PeriodSeconds:       10,
								SuccessThreshold:    1,
								FailureThreshold:    3,
							},
						},
					},
				},
			},
		},
	}
	if err := controllerutil.SetControllerReference(instance, deploy, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// 找不到 deployment，就创建
	fmt.Println("寻找headless 的deploy 是否存在", instance.Name)
	found := appsv1.Deployment{}
	err = r.Get(context.TODO(),
		types.NamespacedName{Name: deploy.Name, Namespace: deploy.Namespace},
		&found)
	if err != nil && errors.IsNotFound(err) {
		err = r.Create(context.TODO(), deploy)
		return reconcile.Result{}, err
	}

	podList := corev1.PodList{}
	err = r.List(context.TODO(), &client.ListOptions{
		LabelSelector: labels.SelectorFromSet(map[string]string{"x": "y"}),
		Namespace:     "",
	}, &podList)
	if err != nil {
		return reconcile.Result{}, err
	}

	fmt.Println(len(podList.Items))

	for _, pod := range podList.Items {
		for _, cs := range pod.Status.ContainerStatuses {
			fmt.Println(pod.Name, cs.Name, cs.Ready)
		}
	}

	fmt.Println(found.Name)
	return reconcile.Result{}, nil
}
