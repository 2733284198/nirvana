// +build !ignore_autogenerated

/*
Copyright 2018 Caicloud Authors

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

// This file was autogenerated by openapi-gen. Do not edit it manually!

package api

import (
	common "github.com/caicloud/nirvana/utils/openapi/common"
	"github.com/go-openapi/spec"
)

// GetOpenAPIDefinitions defines function to get OpenAPI definition
func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/caicloud/nirvana/examples/openapi/pkg/api/v1.Application": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "Application defines application api model",
					Properties: map[string]spec.Schema{
						"name": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"namespace": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"target": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"target2": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"integer"},
								Format: "int",
							},
						},
						"target1": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"boolean"},
								Format: "",
							},
						},
					},
					Required: []string{
						"name",
						"namespace",
						"target",
						"target2",
						"target1",
					},
				},
			},
		},
	}
}
