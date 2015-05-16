package main

import (
	"path/filepath"

	"github.com/emicklei/artreyu/model"
	"github.com/spf13/cobra"
)

type Fetch struct {
	artifact    model.Artifact
	repository  model.Repository
	destination string
	autoExtract bool
}

func (f Fetch) Perform() {
	// check if destination is directory
	var regular string = f.destination
	if model.IsDirectory(f.destination) {
		regular = filepath.Join(f.destination, f.artifact.StorageBase())
	}

	err := f.repository.Fetch(f.artifact, regular)
	if err != nil {
		model.Fatalf("unable to download artifact:%v", err)
	}

	if f.autoExtract && f.artifact.Type == "tgz" {
		if err := model.Untargz(regular, filepath.Dir(regular)); err != nil {
			model.Fatalf("unable to extract artifact:%v", err)
			return
		}
		if err := model.FileRemove(regular); err != nil {
			model.Fatalf("unable to remove compressed artifact:%v", err)
			return
		}
	}
}

func NewFetchCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch [optional:destination]",
		Short: "download an artifact from the repository",
		Long: `destination can be a directory or regular file.
Parent directories will be created if absent.`,
		Run: nil,
	}
	cmd.Flags().BoolVarP(&autoExtract, "extract", "x", false, "extract the content of the compressed artifact")
	return cmd
}
