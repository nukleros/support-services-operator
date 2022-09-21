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
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/nukleros/support-services-operator/apis/platform"

	v1alpha1ingresscomponent "github.com/nukleros/support-services-operator/apis/platform/v1alpha1/ingresscomponent"
	cmdinit "github.com/nukleros/support-services-operator/cmd/ssctl/commands/init"
	//+kubebuilder:scaffold:operator-builder:imports
)

// getIngressComponentManifest returns the sample IngressComponent manifest
// based upon API Version input.
func getIngressComponentManifest(i *cmdinit.InitSubCommand) (string, error) {
	apiVersion := i.APIVersion
	if apiVersion == "" || apiVersion == "latest" {
		return platform.IngressComponentLatestSample, nil
	}

	// generate a map of all versions to samples for each api version created
	manifestMap := map[string]string{
		"v1alpha1": v1alpha1ingresscomponent.Sample(i.RequiredOnly),
		//+kubebuilder:scaffold:operator-builder:versionmap
	}

	// return the manifest if it is not blank
	manifest := manifestMap[apiVersion]
	if manifest != "" {
		return manifest, nil
	}

	// return an error if we did not find a manifest for an api version
	return "", fmt.Errorf("unsupported API Version: " + apiVersion)
}

// NewIngressComponentSubCommand creates a new command and adds it to its
// parent command.
func NewIngressComponentSubCommand(parentCommand *cobra.Command) {
	initCmd := &cmdinit.InitSubCommand{
		Name:         "ingress",
		Description:  "Manage the ingress support services",
		InitFunc:     InitIngressComponent,
		SubCommandOf: parentCommand,
	}

	initCmd.Setup()
}

func InitIngressComponent(i *cmdinit.InitSubCommand) error {
	manifest, err := getIngressComponentManifest(i)
	if err != nil {
		return fmt.Errorf("unable to get manifest for IngressComponent; %w", err)
	}

	outputStream := os.Stdout

	if _, err := outputStream.WriteString(manifest); err != nil {
		return fmt.Errorf("failed to write to stdout, %w", err)
	}

	return nil
}
