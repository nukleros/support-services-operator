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

package platform

import (
	"github.com/spf13/cobra"

	cmdversion "github.com/nukleros/support-services-operator/cmd/ssctl/commands/version"

	"github.com/nukleros/support-services-operator/apis/platform"
)

// NewSecretsComponentSubCommand creates a new command and adds it to its
// parent command.
func NewSecretsComponentSubCommand(parentCommand *cobra.Command) {
	versionCmd := &cmdversion.VersionSubCommand{
		Name:         "secrets",
		Description:  "Manage the secrets management support services",
		VersionFunc:  VersionSecretsComponent,
		SubCommandOf: parentCommand,
	}

	versionCmd.Setup()
}

func VersionSecretsComponent(v *cmdversion.VersionSubCommand) error {
	apiVersions := make([]string, len(platform.SecretsComponentGroupVersions()))

	for i, groupVersion := range platform.SecretsComponentGroupVersions() {
		apiVersions[i] = groupVersion.Version
	}

	versionInfo := cmdversion.VersionInfo{
		CLIVersion:  cmdversion.CLIVersion,
		APIVersions: apiVersions,
	}

	return versionInfo.Display()
}
