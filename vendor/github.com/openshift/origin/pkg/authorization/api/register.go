package api

import (
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/runtime"
)

const GroupName = ""
const FutureGroupName = "authorization.openshift.io"

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = unversioned.GroupVersion{Group: GroupName, Version: runtime.APIVersionInternal}

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) unversioned.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns back a Group qualified GroupResource
func Resource(resource string) unversioned.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func AddToScheme(scheme *runtime.Scheme) {
	// Add the API to Scheme.
	addKnownTypes(scheme)
}

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Role{},
		&RoleBinding{},
		&Policy{},
		&PolicyBinding{},
		&PolicyList{},
		&PolicyBindingList{},
		&RoleBindingList{},
		&RoleList{},

		&SelfSubjectRulesReview{},
		&ResourceAccessReview{},
		&SubjectAccessReview{},
		&LocalResourceAccessReview{},
		&LocalSubjectAccessReview{},
		&ResourceAccessReviewResponse{},
		&SubjectAccessReviewResponse{},
		&IsPersonalSubjectAccessReview{},

		&ClusterRole{},
		&ClusterRoleBinding{},
		&ClusterPolicy{},
		&ClusterPolicyBinding{},
		&ClusterPolicyList{},
		&ClusterPolicyBindingList{},
		&ClusterRoleBindingList{},
		&ClusterRoleList{},
	)
}
