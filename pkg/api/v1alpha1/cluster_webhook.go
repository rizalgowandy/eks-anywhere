// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"fmt"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var clusterlog = logf.Log.WithName("cluster-resource")

func (r *Cluster) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// Change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-anywhere-eks-amazonaws-com-v1alpha1-cluster,mutating=false,failurePolicy=fail,sideEffects=None,groups=anywhere.eks.amazonaws.com,resources=clusters,verbs=create;update,versions=v1alpha1,name=validation.cluster.anywhere.amazonaws.com,admissionReviewVersions={v1,v1beta1}

var _ webhook.Validator = &Cluster{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Cluster) ValidateCreate() error {
	clusterlog.Info("validate create", "name", r.Name)
	if r.IsReconcilePaused() {
		clusterlog.Info("cluster is paused, so allowing create", "name", r.Name)
		return nil
	}
	return apierrors.NewBadRequest("Creating new cluster on existing cluster is not supported")
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Cluster) ValidateUpdate(old runtime.Object) error {
	clusterlog.Info("validate update", "name", r.Name)
	oldCluster, ok := old.(*Cluster)
	if !ok {
		return apierrors.NewBadRequest(fmt.Sprintf("expected a Cluster but got a %T", old))
	}

	var allErrs field.ErrorList

	allErrs = append(allErrs, validateImmutableFieldsCluster(r, oldCluster)...)

	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(GroupVersion.WithKind(ClusterKind).GroupKind(), r.Name, allErrs)
}

func validateImmutableFieldsCluster(new, old *Cluster) field.ErrorList {
	if old.IsReconcilePaused() {
		return nil
	}

	var allErrs field.ErrorList

	if old.Spec.KubernetesVersion != new.Spec.KubernetesVersion {
		allErrs = append(
			allErrs,
			field.Invalid(field.NewPath("spec", "kubernetesVersion"), new.Spec.KubernetesVersion, "field is immutable"),
		)
	}

	if old.Spec.ControlPlaneConfiguration.Count != new.Spec.ControlPlaneConfiguration.Count {
		allErrs = append(
			allErrs,
			field.Invalid(field.NewPath("spec", "ControlPlaneConfiguration.count"), new.Spec.ControlPlaneConfiguration.Count, "field is immutable"))
	}

	if !new.Spec.ControlPlaneConfiguration.Endpoint.Equal(old.Spec.ControlPlaneConfiguration.Endpoint) {
		allErrs = append(
			allErrs,
			field.Invalid(field.NewPath("spec", "ControlPlaneConfiguration.endpoint"), new.Spec.ControlPlaneConfiguration.Endpoint, "field is immutable"))
	}

	if !new.Spec.DatacenterRef.Equal(&old.Spec.DatacenterRef) {
		allErrs = append(
			allErrs,
			field.Invalid(field.NewPath("spec", "datacenterRef"), new.Spec.DatacenterRef, "field is immutable"))
	}

	if !new.Spec.ControlPlaneConfiguration.MachineGroupRef.Equal(old.Spec.ControlPlaneConfiguration.MachineGroupRef) {
		allErrs = append(
			allErrs,
			field.Invalid(field.NewPath("spec", "ControlPlaneConfiguration.machineGroupRef"), new.Spec.ControlPlaneConfiguration.MachineGroupRef, "field is immutable"))
	}

	if new.Spec.ExternalEtcdConfiguration != nil && old.Spec.ExternalEtcdConfiguration == nil {
		allErrs = append(
			allErrs,
			field.Invalid(field.NewPath("spec.externalEtcdConfiguration"), new.Spec.ExternalEtcdConfiguration, "cannot switch from local to external etcd topology"),
		)
	}
	if new.Spec.ExternalEtcdConfiguration != nil && old.Spec.ExternalEtcdConfiguration != nil {
		if old.Spec.ExternalEtcdConfiguration.Count != new.Spec.ExternalEtcdConfiguration.Count {
			allErrs = append(
				allErrs,
				field.Invalid(field.NewPath("spec.externalEtcdConfiguration.count"), new.Spec.ExternalEtcdConfiguration.Count, "field is immutable"),
			)
		}
	}

	if !new.Spec.ClusterNetwork.Equal(&old.Spec.ClusterNetwork) {
		allErrs = append(
			allErrs,
			field.Invalid(field.NewPath("spec", "ClusterNetwork"), new.Spec.ClusterNetwork, "field is immutable"))
	}

	if !new.Spec.ProxyConfiguration.Equal(old.Spec.ProxyConfiguration) {
		allErrs = append(
			allErrs,
			field.Invalid(field.NewPath("spec", "ProxyConfiguration"), new.Spec.ProxyConfiguration, "field is immutable"))
	}

	if !new.Spec.GitOpsRef.Equal(old.Spec.GitOpsRef) {
		allErrs = append(
			allErrs,
			field.Invalid(field.NewPath("spec", "GitOpsRef"), new.Spec.GitOpsRef, "field is immutable"))
	}

	if !RefSliceEqual(new.Spec.IdentityProviderRefs, old.Spec.IdentityProviderRefs) {
		allErrs = append(
			allErrs,
			field.Invalid(field.NewPath("spec", "IdentityProviderRefs"), new.Spec.IdentityProviderRefs, "field is immutable"))
	}

	return allErrs
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Cluster) ValidateDelete() error {
	clusterlog.Info("validate delete", "name", r.Name)

	return nil
}
