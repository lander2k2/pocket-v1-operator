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

package commands

import (
	"github.com/spf13/cobra"

	// common imports for subcommands
	cmdgenerate "github.com/lander2k2/pocket-v1-operator/cmd/pocketctl/commands/generate"
	cmdinit "github.com/lander2k2/pocket-v1-operator/cmd/pocketctl/commands/init"
	cmdversion "github.com/lander2k2/pocket-v1-operator/cmd/pocketctl/commands/version"

	// specific imports for workloads
	generatenodes "github.com/lander2k2/pocket-v1-operator/cmd/pocketctl/commands/generate/nodes"
	initnodes "github.com/lander2k2/pocket-v1-operator/cmd/pocketctl/commands/init/nodes"
	versionnodes "github.com/lander2k2/pocket-v1-operator/cmd/pocketctl/commands/version/nodes"
	//+kubebuilder:scaffold:operator-builder:subcommands:imports
)

// PocketctlCommand represents the base command when called without any subcommands.
type PocketctlCommand struct {
	*cobra.Command
}

// NewPocketctlCommand returns an instance of the PocketctlCommand.
func NewPocketctlCommand() *PocketctlCommand {
	c := &PocketctlCommand{
		Command: &cobra.Command{
			Use:   "pocketctl",
			Short: "Manage v1 pocket node deployments",
			Long:  "Manage v1 pocket node deployments",
		},
	}

	c.addSubCommands()

	return c
}

// Run represents the main entry point into the command
// This is called by main.main() to execute the root command.
func (c *PocketctlCommand) Run() {
	cobra.CheckErr(c.Execute())
}

func (c *PocketctlCommand) newInitSubCommand() {
	parentCommand := cmdinit.GetParent(cmdinit.NewBaseInitSubCommand(c.Command))
	_ = parentCommand

	// add the init subcommands
	initnodes.NewPocketSetSubCommand(parentCommand)
	initnodes.NewPocketValidatorSubCommand(parentCommand)
	//+kubebuilder:scaffold:operator-builder:subcommands:init
}

func (c *PocketctlCommand) newGenerateSubCommand() {
	parentCommand := cmdgenerate.GetParent(cmdgenerate.NewBaseGenerateSubCommand(c.Command))
	_ = parentCommand

	// add the generate subcommands
	generatenodes.NewPocketSetSubCommand(parentCommand)
	generatenodes.NewPocketValidatorSubCommand(parentCommand)
	//+kubebuilder:scaffold:operator-builder:subcommands:generate
}

func (c *PocketctlCommand) newVersionSubCommand() {
	parentCommand := cmdversion.GetParent(cmdversion.NewBaseVersionSubCommand(c.Command))
	_ = parentCommand

	// add the version subcommands
	versionnodes.NewPocketSetSubCommand(parentCommand)
	versionnodes.NewPocketValidatorSubCommand(parentCommand)
	//+kubebuilder:scaffold:operator-builder:subcommands:version
}

// addSubCommands adds any additional subCommands to the root command.
func (c *PocketctlCommand) addSubCommands() {
	c.newInitSubCommand()
	c.newGenerateSubCommand()
	c.newVersionSubCommand()
}
