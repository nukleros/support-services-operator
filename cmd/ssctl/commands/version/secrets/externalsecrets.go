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

package secrets

import (
	"github.com/spf13/cobra"

	cmdversion "github.com/nukleros/support-services-operator/cmd/ssctl/commands/version"

	"github.com/nukleros/support-services-operator/apis/secrets"
)

// NewExternalSecretsSubCommand creates a new command and adds it to its
// parent command.
func NewExternalSecretsSubCommand(parentCommand *cobra.Command) {
	versionCmd := &cmdversion.VersionSubCommand{
		Name:         "external-secrets",
		Description:  "Manage exteranl-secrets installation for secrets support services",
		VersionFunc:  VersionExternalSecrets,
		SubCommandOf: parentCommand,
	}

	versionCmd.Setup()
}

func VersionExternalSecrets(v *cmdversion.VersionSubCommand) error {
	apiVersions := make([]string, len(secrets.ExternalSecretsGroupVersions()))

	for i, groupVersion := range secrets.ExternalSecretsGroupVersions() {
		apiVersions[i] = groupVersion.Version
	}

	versionInfo := cmdversion.VersionInfo{
		CLIVersion:  cmdversion.CLIVersion,
		APIVersions: apiVersions,
	}

	return versionInfo.Display()
}
