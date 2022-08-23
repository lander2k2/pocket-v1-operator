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
	"github.com/spf13/cobra"

	cmdversion "github.com/lander2k2/pocket-v1-operator/cmd/pocketctl/commands/version"

	"github.com/lander2k2/pocket-v1-operator/apis/nodes"
)

// NewPocketSetSubCommand creates a new command and adds it to its
// parent command.
func NewPocketSetSubCommand(parentCommand *cobra.Command) {
	versionCmd := &cmdversion.VersionSubCommand{
		Name:         "set",
		Description:  "Manage sets of pocket nodes",
		VersionFunc:  VersionPocketSet,
		SubCommandOf: parentCommand,
	}

	versionCmd.Setup()
}

func VersionPocketSet(v *cmdversion.VersionSubCommand) error {
	apiVersions := make([]string, len(nodes.PocketSetGroupVersions()))

	for i, groupVersion := range nodes.PocketSetGroupVersions() {
		apiVersions[i] = groupVersion.Version
	}

	versionInfo := cmdversion.VersionInfo{
		CLIVersion:  cmdversion.CLIVersion,
		APIVersions: apiVersions,
	}

	return versionInfo.Display()
}
