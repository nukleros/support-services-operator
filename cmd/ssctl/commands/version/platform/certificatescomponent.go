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

// NewCertificatesComponentSubCommand creates a new command and adds it to its
// parent command.
func NewCertificatesComponentSubCommand(parentCommand *cobra.Command) {
	versionCmd := &cmdversion.VersionSubCommand{
		Name:         "certificates",
		Description:  "Manage the certificate support services",
		VersionFunc:  VersionCertificatesComponent,
		SubCommandOf: parentCommand,
	}

	versionCmd.Setup()
}

func VersionCertificatesComponent(v *cmdversion.VersionSubCommand) error {
	apiVersions := make([]string, len(platform.CertificatesComponentGroupVersions()))

	for i, groupVersion := range platform.CertificatesComponentGroupVersions() {
		apiVersions[i] = groupVersion.Version
	}

	versionInfo := cmdversion.VersionInfo{
		CLIVersion:  cmdversion.CLIVersion,
		APIVersions: apiVersions,
	}

	return versionInfo.Display()
}
