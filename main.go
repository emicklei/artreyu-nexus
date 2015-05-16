package main

import (
	"github.com/emicklei/artreyu/model"
	"github.com/spf13/cobra"
)

type ArtifactCommand struct {
	*cobra.Command
	artifact model.Artifact
}

func NewArtifactCommand() *ArtifactCommand {
	cmd := new(cobra.Command)
	artifact := model.Artifact{}
	cmd.PersistentFlags().StringVarP(&artifact.Name,
		"artifact",
		"a",
		"",
		"name of the artifact")
	cmd.PersistentFlags().StringVarP(&artifact.Group,
		"group",
		"g",
		"",
		"name of the group")
	cmd.PersistentFlags().StringVarP(&artifact.Version,
		"version",
		"s",
		"",
		"version of the artifact")
	cmd.PersistentFlags().StringVarP(&artifact.Type,
		"type",
		"t",
		"",
		"type (extension) of the artifact")
	return &ArtifactCommand{cmd, artifact}
}

func main() {
	cmd := NewArtifactCommand()
	cmd.Command.Use = "artreyu-nexus"
	cmd.Command.Short = "archives and fetches from a Sonatype Nexus Repository"

	//	cmd.AddCommand(NewArchiveCommand())
	//	cmd.AddCommand(NewFetchCommand())
}
