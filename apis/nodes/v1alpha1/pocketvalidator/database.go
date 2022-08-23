/*
Copyright 2022.

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

package pocketvalidator

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	nodesv1alpha1 "github.com/lander2k2/pocket-v1-operator/apis/nodes/v1alpha1"
)

// +kubebuilder:rbac:groups=acid.zalan.do,resources=postgresqls,verbs=get;list;watch;create;update;patch;delete

// CreatepostgresqlCollectionNameParentNameDatabase creates the !!start parent.Name !!end-database postgresql resource.
func CreatepostgresqlCollectionNameParentNameDatabase(
	parent *nodesv1alpha1.PocketValidator,
	collection *nodesv1alpha1.PocketSet,
) ([]client.Object, error) {

	resourceObjs := []client.Object{}
	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "acid.zalan.do/v1",
			"kind":       "postgresql",
			"metadata": map[string]interface{}{
				"name":      "" + parent.Name + "-database", //  controlled by field:
				"namespace": collection.Name,                //  controlled by collection field:
			},
			"spec": map[string]interface{}{
				"teamId": parent.Name, //  controlled by field:
				"volume": map[string]interface{}{
					"size": "1Gi",
				},
				"numberOfInstances": parent.Spec.DbReplicas, //  controlled by field: dbReplicas
				"users": map[string]interface{}{
					"validator": []interface{}{
						"superuser",
						"createdb",
						//foo_user: []  # role for application foo
					},
				},
				"databases": map[string]interface{}{
					"validatordb": "validator", //  dbname: owner
				},
				"preparedDatabases": map[string]interface{}{
					"bar": map[string]interface{}{},
				},
				"postgresql": map[string]interface{}{
					"version": "14",
				},
			},
		},
	}

	resourceObjs = append(resourceObjs, resourceObj)

	return resourceObjs, nil
}
