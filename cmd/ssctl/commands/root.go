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
	cmdgenerate "github.com/nukleros/support-services-operator/cmd/ssctl/commands/generate"
	cmdinit "github.com/nukleros/support-services-operator/cmd/ssctl/commands/init"
	cmdversion "github.com/nukleros/support-services-operator/cmd/ssctl/commands/version"

	// specific imports for workloads
	generateapplication "github.com/nukleros/support-services-operator/cmd/ssctl/commands/generate/application"
	generateservices "github.com/nukleros/support-services-operator/cmd/ssctl/commands/generate/services"
	initapplication "github.com/nukleros/support-services-operator/cmd/ssctl/commands/init/application"
	initservices "github.com/nukleros/support-services-operator/cmd/ssctl/commands/init/services"
	versionapplication "github.com/nukleros/support-services-operator/cmd/ssctl/commands/version/application"
	versionservices "github.com/nukleros/support-services-operator/cmd/ssctl/commands/version/services"
	//+kubebuilder:scaffold:operator-builder:subcommands:imports
)

// SsctlCommand represents the base command when called without any subcommands.
type SsctlCommand struct {
	*cobra.Command
}

// NewSsctlCommand returns an instance of the SsctlCommand.
func NewSsctlCommand() *SsctlCommand {
	c := &SsctlCommand{
		Command: &cobra.Command{
			Use:   "ssctl",
			Short: "Manage a cluster's support service installations",
			Long:  "Manage a cluster's support service installations",
		},
	}

	c.addSubCommands()

	return c
}

// Run represents the main entry point into the command
// This is called by main.main() to execute the root command.
func (c *SsctlCommand) Run() {
	cobra.CheckErr(c.Execute())
}

func (c *SsctlCommand) newInitSubCommand() {
	parentCommand := cmdinit.GetParent(cmdinit.NewBaseInitSubCommand(c.Command))
	_ = parentCommand

	// add the init subcommands
	initservices.NewSupportServicesSubCommand(parentCommand)
	initapplication.NewDatabaseComponentSubCommand(parentCommand)
	//+kubebuilder:scaffold:operator-builder:subcommands:init
}

func (c *SsctlCommand) newGenerateSubCommand() {
	parentCommand := cmdgenerate.GetParent(cmdgenerate.NewBaseGenerateSubCommand(c.Command))
	_ = parentCommand

	// add the generate subcommands
	generateservices.NewSupportServicesSubCommand(parentCommand)
	generateapplication.NewDatabaseComponentSubCommand(parentCommand)
	//+kubebuilder:scaffold:operator-builder:subcommands:generate
}

func (c *SsctlCommand) newVersionSubCommand() {
	parentCommand := cmdversion.GetParent(cmdversion.NewBaseVersionSubCommand(c.Command))
	_ = parentCommand

	// add the version subcommands
	versionservices.NewSupportServicesSubCommand(parentCommand)
	versionapplication.NewDatabaseComponentSubCommand(parentCommand)
	//+kubebuilder:scaffold:operator-builder:subcommands:version
}

// addSubCommands adds any additional subCommands to the root command.
func (c *SsctlCommand) addSubCommands() {
	c.newInitSubCommand()
	c.newGenerateSubCommand()
	c.newVersionSubCommand()
}
