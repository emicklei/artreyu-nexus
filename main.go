package main

import (
	"github.com/emicklei/artreyu/command"
	"github.com/emicklei/artreyu/model"
)

var VERSION string = "dev"
var BUILDDATE string = "now"

func main() {
	model.Printf("artreyu-nexus - artreyu Sonatype Nexus plugin (build:%s, commit:%s)\n", BUILDDATE, VERSION)

	cmd, settings, artifact := command.NewPluginCommand()
	cmd.Use = "artreyu-nexus"
	cmd.Short = "archives and fetches from a Sonatype Nexus Repository"

	// Need closures because only after cmd.Execute() the model data is populated.
	getArtifact := func() model.Artifact {
		return *artifact
	}
	getRepo := func() model.Repository {
		return NewRepository(model.RepositoryConfigNamed(settings, "nexus"), settings)
	}

	cmd.AddCommand(command.NewArchiveCommand(getArtifact, getRepo))
	cmd.AddCommand(command.NewFetchCommand(getArtifact, getRepo))
	cmd.Execute()
}
