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

// +kubebuilder:rbac:groups=apps,resources=statefulsets,verbs=get;list;watch;create;update;patch;delete

const StatefulSetCollectionNameParentName = "parent.Name"

// CreateStatefulSetCollectionNameParentName creates the parent.Name StatefulSet resource.
func CreateStatefulSetCollectionNameParentName(
	parent *nodesv1alpha1.PocketValidator,
	collection *nodesv1alpha1.PocketSet,
) ([]client.Object, error) {

	resourceObjs := []client.Object{}
	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "StatefulSet",
			"metadata": map[string]interface{}{
				"name":      parent.Name,     //  controlled by field:
				"namespace": collection.Name, //  controlled by collection field:
			},
			"spec": map[string]interface{}{
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": parent.Name, //  controlled by field:
					},
				},
				"serviceName": parent.Name,                   //  controlled by field:
				"replicas":    parent.Spec.ValidatorReplicas, //  controlled by field: validatorReplicas
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app":        parent.Name, //  controlled by field:
							"v1-purpose": "validator",
						},
					},
					"spec": map[string]interface{}{
						"containers": []interface{}{
							map[string]interface{}{
								"name":  "pocket",
								"image": parent.Spec.PocketImage, //  controlled by field: pocketImage
								"args": []interface{}{
									"pocket",
									"-config=/configs/config.json",
									"-genesis=/genesis.json",
								},
								"ports": []interface{}{
									map[string]interface{}{
										"containerPort": 8221,
										"name":          "pre2p",
									},
									map[string]interface{}{
										"containerPort": 8222,
										"name":          "p2p",
									},
								},
								"volumeMounts": []interface{}{
									map[string]interface{}{
										"name":      "config-volume",
										"mountPath": "/configs",
									},
									map[string]interface{}{
										"name":      "genesis-volume",
										"mountPath": "/genesis.json",
										"subPath":   "genesis.json",
									},
									map[string]interface{}{
										"name":      "" + parent.Name + "-blockstore", //  controlled by field:
										"mountPath": "/blockstore",
									},
								},
							},
						},
						"volumes": []interface{}{
							map[string]interface{}{
								"name": "config-volume",
								"configMap": map[string]interface{}{
									"name": "" + parent.Name + "-config", //  controlled by field:
								},
							},
							map[string]interface{}{
								"name": "genesis-volume",
								"configMap": map[string]interface{}{
									"name": "" + collection.Name + "-genesis", //  controlled by collection field:
								},
							},
						},
					},
				},
				"volumeClaimTemplates": []interface{}{
					map[string]interface{}{
						"metadata": map[string]interface{}{
							"name": "" + parent.Name + "-blockstore", //  controlled by field:
						},
						"spec": map[string]interface{}{
							"accessModes": []interface{}{
								"ReadWriteOnce",
							},
							"resources": map[string]interface{}{
								"requests": map[string]interface{}{
									"storage": "1Gi",
								},
							},
						},
					},
				},
			},
		},
	}

	resourceObjs = append(resourceObjs, resourceObj)

	return resourceObjs, nil
}

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

const ServiceCollectionNameParentName = "parent.Name"

// CreateServiceCollectionNameParentName creates the parent.Name Service resource.
func CreateServiceCollectionNameParentName(
	parent *nodesv1alpha1.PocketValidator,
	collection *nodesv1alpha1.PocketSet,
) ([]client.Object, error) {

	resourceObjs := []client.Object{}
	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      parent.Name,     //  controlled by field:
				"namespace": collection.Name, //  controlled by collection field:
				"labels": map[string]interface{}{
					"app": parent.Name, //  controlled by field:
				},
			},
			"spec": map[string]interface{}{
				"ports": []interface{}{
					map[string]interface{}{
						"port": 8221,
						"name": "pre2p",
					},
					map[string]interface{}{
						"port": 8222,
						"name": "p2p",
					},
				},
				"selector": map[string]interface{}{
					"app": parent.Name, //  controlled by field:
				},
			},
		},
	}

	resourceObjs = append(resourceObjs, resourceObj)

	return resourceObjs, nil
}

// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;patch;delete

// CreateConfigMapCollectionNameParentNameConfig creates the !!start parent.Name !!end-config ConfigMap resource.
func CreateConfigMapCollectionNameParentNameConfig(
	parent *nodesv1alpha1.PocketValidator,
	collection *nodesv1alpha1.PocketSet,
) ([]client.Object, error) {

	resourceObjs := []client.Object{}
	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]interface{}{
				"name":      "" + parent.Name + "-config", //  controlled by field:
				"namespace": collection.Name,              //  controlled by collection field:
			},
			"data": map[string]interface{}{
				// controlled by field: privateKey
				// controlled by field:
				"config.json": `{
  "base": {
    "root_directory": "/go/src/github.com/pocket-network",
    "private_key": "` + parent.Spec.PrivateKey + `"
  },
  "consensus": {
    "max_mempool_bytes": 500000000,
    "pacemaker_config": {
      "timeout_msec": 5000,
      "manual": true,
      "debug_time_between_steps_msec": 1000
    }
  },
  "utility": {},
  "persistence": {
    "postgres_url": "postgres://validator:postgres@` + parent.Name + `-database:5432/validatordb",
    "node_schema": "validator",
    "block_store_path": "/blockstore"
  },
  "p2p": {
    "consensus_port": 8080,
    "use_rain_tree": true,
    "connection_type": 1
  },
  "telemetry": {
    "enabled": true,
    "address": "0.0.0.0:9000",
    "endpoint": "/metrics"
  }
}
`,
			},
		},
	}

	resourceObjs = append(resourceObjs, resourceObj)

	return resourceObjs, nil
}
