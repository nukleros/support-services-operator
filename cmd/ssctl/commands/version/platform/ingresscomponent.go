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

package platform

import (
	"github.com/spf13/cobra"

	cmdversion "github.com/nukleros/support-services-operator/cmd/ssctl/commands/version"

	"github.com/nukleros/support-services-operator/apis/platform"
)

// NewIngressComponentSubCommand creates a new command and adds it to its
// parent command.
func NewIngressComponentSubCommand(parentCommand *cobra.Command) {
	versionCmd := &cmdversion.VersionSubCommand{
		Name:         "ingress",
		Description:  "Manage the ingress support services",
		VersionFunc:  VersionIngressComponent,
		SubCommandOf: parentCommand,
	}

	versionCmd.Setup()
}

func VersionIngressComponent(v *cmdversion.VersionSubCommand) error {
	apiVersions := make([]string, len(platform.IngressComponentGroupVersions()))

	for i, groupVersion := range platform.IngressComponentGroupVersions() {
		apiVersions[i] = groupVersion.Version
	}

	versionInfo := cmdversion.VersionInfo{
		CLIVersion:  cmdversion.CLIVersion,
		APIVersions: apiVersions,
	}

	return versionInfo.Display()
}
