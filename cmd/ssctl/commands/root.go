/*
Copyright 2023.

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
	generatecertificates "github.com/nukleros/support-services-operator/cmd/ssctl/commands/generate/certificates"
	generateingress "github.com/nukleros/support-services-operator/cmd/ssctl/commands/generate/ingress"
	generateorchestration "github.com/nukleros/support-services-operator/cmd/ssctl/commands/generate/orchestration"
	generatesecrets "github.com/nukleros/support-services-operator/cmd/ssctl/commands/generate/secrets"
	initcertificates "github.com/nukleros/support-services-operator/cmd/ssctl/commands/init/certificates"
	initingress "github.com/nukleros/support-services-operator/cmd/ssctl/commands/init/ingress"
	initorchestration "github.com/nukleros/support-services-operator/cmd/ssctl/commands/init/orchestration"
	initsecrets "github.com/nukleros/support-services-operator/cmd/ssctl/commands/init/secrets"
	versioncertificates "github.com/nukleros/support-services-operator/cmd/ssctl/commands/version/certificates"
	versioningress "github.com/nukleros/support-services-operator/cmd/ssctl/commands/version/ingress"
	versionorchestration "github.com/nukleros/support-services-operator/cmd/ssctl/commands/version/orchestration"
	versionsecrets "github.com/nukleros/support-services-operator/cmd/ssctl/commands/version/secrets"
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
			Short: "Manage Kubernetes cluster support service installations",
			Long:  "Manage Kubernetes cluster support service installations",
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
	initorchestration.NewSupportServicesSubCommand(parentCommand)
	initcertificates.NewCertManagerSubCommand(parentCommand)
	initingress.NewExternalDNSSubCommand(parentCommand)
	initsecrets.NewExternalSecretsSubCommand(parentCommand)
	initsecrets.NewReloaderSubCommand(parentCommand)
	//+kubebuilder:scaffold:operator-builder:subcommands:init
}

func (c *SsctlCommand) newGenerateSubCommand() {
	parentCommand := cmdgenerate.GetParent(cmdgenerate.NewBaseGenerateSubCommand(c.Command))
	_ = parentCommand

	// add the generate subcommands
	generateorchestration.NewSupportServicesSubCommand(parentCommand)
	generatecertificates.NewCertManagerSubCommand(parentCommand)
	generateingress.NewExternalDNSSubCommand(parentCommand)
	generatesecrets.NewExternalSecretsSubCommand(parentCommand)
	generatesecrets.NewReloaderSubCommand(parentCommand)
	//+kubebuilder:scaffold:operator-builder:subcommands:generate
}

func (c *SsctlCommand) newVersionSubCommand() {
	parentCommand := cmdversion.GetParent(cmdversion.NewBaseVersionSubCommand(c.Command))
	_ = parentCommand

	// add the version subcommands
	versionorchestration.NewSupportServicesSubCommand(parentCommand)
	versioncertificates.NewCertManagerSubCommand(parentCommand)
	versioningress.NewExternalDNSSubCommand(parentCommand)
	versionsecrets.NewExternalSecretsSubCommand(parentCommand)
	versionsecrets.NewReloaderSubCommand(parentCommand)
	//+kubebuilder:scaffold:operator-builder:subcommands:version
}

// addSubCommands adds any additional subCommands to the root command.
func (c *SsctlCommand) addSubCommands() {
	c.newInitSubCommand()
	c.newGenerateSubCommand()
	c.newVersionSubCommand()
}
