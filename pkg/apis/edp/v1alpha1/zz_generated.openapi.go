// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	"github.com/go-openapi/spec"
	"k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"codebase-operator/pkg/apis/edp/v1alpha1.ApplicationBranch":       schema_pkg_apis_edp_v1alpha1_ApplicationBranch(ref),
		"codebase-operator/pkg/apis/edp/v1alpha1.ApplicationBranchSpec":   schema_pkg_apis_edp_v1alpha1_ApplicationBranchSpec(ref),
		"codebase-operator/pkg/apis/edp/v1alpha1.ApplicationBranchStatus": schema_pkg_apis_edp_v1alpha1_ApplicationBranchStatus(ref),
		"codebase-operator/pkg/apis/edp/v1alpha1.Codebase":                schema_pkg_apis_edp_v1alpha1_Codebase(ref),
		"codebase-operator/pkg/apis/edp/v1alpha1.CodebaseStatus":          schema_pkg_apis_edp_v1alpha1_CodebaseStatus(ref),
	}
}

func schema_pkg_apis_edp_v1alpha1_ApplicationBranch(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ApplicationBranch is the Schema for the applicationbranches API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("codebase-operator/pkg/apis/edp/v1alpha1.ApplicationBranchSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("codebase-operator/pkg/apis/edp/v1alpha1.ApplicationBranchStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"codebase-operator/pkg/apis/edp/v1alpha1.ApplicationBranchSpec", "codebase-operator/pkg/apis/edp/v1alpha1.ApplicationBranchStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_edp_v1alpha1_ApplicationBranchSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ApplicationBranchSpec defines the desired state of ApplicationBranch",
				Properties: map[string]spec.Schema{
					"appName": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"branchName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"fromCommit": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"appName", "branchName", "fromCommit"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_edp_v1alpha1_ApplicationBranchStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ApplicationBranchStatus defines the observed state of ApplicationBranch",
				Properties: map[string]spec.Schema{
					"lastTimeUpdated": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "date-time",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"lastTimeUpdated", "status"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_edp_v1alpha1_Codebase(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Codebase is the Schema for the codebases API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("codebase-operator/pkg/apis/edp/v1alpha1.CodebaseSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("codebase-operator/pkg/apis/edp/v1alpha1.CodebaseStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"codebase-operator/pkg/apis/edp/v1alpha1.CodebaseSpec", "codebase-operator/pkg/apis/edp/v1alpha1.CodebaseStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_edp_v1alpha1_CodebaseStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CodebaseStatus defines the observed state of Codebase",
				Properties: map[string]spec.Schema{
					"available": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"lastTimeUpdated": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "date-time",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"available", "lastTimeUpdated", "status"},
			},
		},
		Dependencies: []string{},
	}
}
