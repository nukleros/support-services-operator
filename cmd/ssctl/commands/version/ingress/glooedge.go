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

package ingress

import (
	"github.com/spf13/cobra"

	cmdversion "github.com/nukleros/support-services-operator/cmd/ssctl/commands/version"

	"github.com/nukleros/support-services-operator/apis/ingress"
)

// NewGlooEdgeSubCommand creates a new command and adds it to its
// parent command.
func NewGlooEdgeSubCommand(parentCommand *cobra.Command) {
	versionCmd := &cmdversion.VersionSubCommand{
		Name:         "gloo-edge",
		Description:  "Manage gloo-edge installation for ingress support services",
		VersionFunc:  VersionGlooEdge,
		SubCommandOf: parentCommand,
	}

	versionCmd.Setup()
}

func VersionGlooEdge(v *cmdversion.VersionSubCommand) error {
	apiVersions := make([]string, len(ingress.GlooEdgeGroupVersions()))

	for i, groupVersion := range ingress.GlooEdgeGroupVersions() {
		apiVersions[i] = groupVersion.Version
	}

	versionInfo := cmdversion.VersionInfo{
		CLIVersion:  cmdversion.CLIVersion,
		APIVersions: apiVersions,
	}

	return versionInfo.Display()
}
