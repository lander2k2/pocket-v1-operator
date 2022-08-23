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

package nodes

import (
	v1alpha1nodes "github.com/lander2k2/pocket-v1-operator/apis/nodes/v1alpha1"
	//+kubebuilder:scaffold:operator-builder:imports

	"k8s.io/apimachinery/pkg/runtime/schema"
)

// PocketValidatorGroupVersions returns all group version objects associated with this kind.
func PocketValidatorGroupVersions() []schema.GroupVersion {
	return []schema.GroupVersion{
		v1alpha1nodes.GroupVersion,
		//+kubebuilder:scaffold:operator-builder:groupversions
	}
}
