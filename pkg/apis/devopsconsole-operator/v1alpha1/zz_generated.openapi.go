// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/redhat-developer/devopsconsole-operator/pkg/apis/devopsconsole-operator/v1alpha1.EnvDeployment": schema_pkg_apis_devopsconsole_operator_v1alpha1_EnvDeployment(ref),
		"github.com/redhat-developer/devopsconsole-operator/pkg/apis/devopsconsole-operator/v1alpha1.Environment":   schema_pkg_apis_devopsconsole_operator_v1alpha1_Environment(ref),
	}
}

func schema_pkg_apis_devopsconsole_operator_v1alpha1_EnvDeployment(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EnvDeployment is the Schema for the envdeployments API",
				Type:        []string{"object"},
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
							Ref: ref("github.com/redhat-developer/devopsconsole-operator/pkg/apis/devopsconsole-operator/v1alpha1.EnvDeploymentSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/redhat-developer/devopsconsole-operator/pkg/apis/devopsconsole-operator/v1alpha1.EnvDeploymentStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/redhat-developer/devopsconsole-operator/pkg/apis/devopsconsole-operator/v1alpha1.EnvDeploymentSpec", "github.com/redhat-developer/devopsconsole-operator/pkg/apis/devopsconsole-operator/v1alpha1.EnvDeploymentStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_devopsconsole_operator_v1alpha1_Environment(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Environment is the Schema for the environments API",
				Type:        []string{"object"},
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
							Ref: ref("github.com/redhat-developer/devopsconsole-operator/pkg/apis/devopsconsole-operator/v1alpha1.EnvironmentSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/redhat-developer/devopsconsole-operator/pkg/apis/devopsconsole-operator/v1alpha1.EnvironmentStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/redhat-developer/devopsconsole-operator/pkg/apis/devopsconsole-operator/v1alpha1.EnvironmentSpec", "github.com/redhat-developer/devopsconsole-operator/pkg/apis/devopsconsole-operator/v1alpha1.EnvironmentStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}
