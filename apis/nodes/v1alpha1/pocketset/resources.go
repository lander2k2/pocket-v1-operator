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

package pocketset

import (
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	nodesv1alpha1 "github.com/lander2k2/pocket-v1-operator/apis/nodes/v1alpha1"
)

// samplePocketSet is a sample containing all fields
const samplePocketSet = `apiVersion: nodes.pokt.network/v1alpha1
kind: PocketSet
metadata:
  name: pocketset-sample
spec:
`

// samplePocketSetRequired is a sample containing only required fields
const samplePocketSetRequired = `apiVersion: nodes.pokt.network/v1alpha1
kind: PocketSet
metadata:
  name: pocketset-sample
spec:
`

// Sample returns the sample manifest for this custom resource.
func Sample(requiredOnly bool) string {
	if requiredOnly {
		return samplePocketSetRequired
	}

	return samplePocketSet
}

// Generate returns the child resources that are associated with this workload given
// appropriate structured inputs.
func Generate(collectionObj nodesv1alpha1.PocketSet) ([]client.Object, error) {
	resourceObjects := []client.Object{}

	for _, f := range CreateFuncs {
		resources, err := f(&collectionObj)

		if err != nil {
			return nil, err
		}

		resourceObjects = append(resourceObjects, resources...)
	}

	return resourceObjects, nil
}

// GenerateForCLI returns the child resources that are associated with this workload given
// appropriate YAML manifest files.
func GenerateForCLI(collectionFile []byte) ([]client.Object, error) {
	var collectionObj nodesv1alpha1.PocketSet
	if err := yaml.Unmarshal(collectionFile, &collectionObj); err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml into collection, %w", err)
	}

	if err := workload.Validate(&collectionObj); err != nil {
		return nil, fmt.Errorf("error validating collection yaml, %w", err)
	}

	return Generate(collectionObj)
}

// CreateFuncs is an array of functions that are called to create the child resources for the controller
// in memory during the reconciliation loop prior to persisting the changes or updates to the Kubernetes
// database.
var CreateFuncs = []func(
	*nodesv1alpha1.PocketSet,
) ([]client.Object, error){
	CreateNamespaceParentName,
	CreateServiceParentNameParentNameValidators,
	CreateConfigMapParentNameParentNameGenesis,
}

// InitFuncs is an array of functions that are called prior to starting the controller manager.  This is
// necessary in instances which the controller needs to "own" objects which depend on resources to
// pre-exist in the cluster. A common use case for this is the need to own a custom resource.
// If the controller needs to own a custom resource type, the CRD that defines it must
// first exist. In this case, the InitFunc will create the CRD so that the controller
// can own custom resources of that type.  Without the InitFunc the controller will
// crash loop because when it tries to own a non-existent resource type during manager
// setup, it will fail.
var InitFuncs = []func(
	*nodesv1alpha1.PocketSet,
) ([]client.Object, error){}

func ConvertWorkload(component workload.Workload) (*nodesv1alpha1.PocketSet, error) {
	p, ok := component.(*nodesv1alpha1.PocketSet)
	if !ok {
		return nil, nodesv1alpha1.ErrUnableToConvertPocketSet
	}

	return p, nil
}
