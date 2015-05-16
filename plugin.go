package main

import (
	"github.com/emicklei/artreyu/model"
	"github.com/spf13/cobra"
)

func SetupArtifactFlags(cmd *cobra.Command, a *model.Artifact) {
	cmd.PersistentFlags().StringVarP(&a.Name,
		"artifact",
		"a",
		"",
		"name of the artifact")
	cmd.PersistentFlags().StringVarP(&a.Group,
		"group",
		"g",
		"",
		"name of the group")
	cmd.PersistentFlags().StringVarP(&a.Version,
		"version",
		"s",
		"",
		"version of the artifact")
	cmd.PersistentFlags().StringVarP(&a.Type,
		"type",
		"t",
		"",
		"type (extension) of the artifact")
}
