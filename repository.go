package main

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/emicklei/artreyu/model"
)

type Repository struct {
	settings *model.Settings
	config   model.RepositoryConfig
}

func NewRepository(config model.RepositoryConfig, settings *model.Settings) Repository {
	return Repository{settings, config}
}

func (r Repository) ID() string { return "nexus" }

func (r Repository) Store(a model.Artifact, source string) error {
	repo := "releases"
	if a.IsSnapshot() {
		repo = "snapshots"
	}
	destination := r.config.URL + filepath.Join(r.config.Path, repo, a.StorageLocation(r.settings.OS, a.AnyOS))
	model.Printf("uploading %s to %s\n", source, destination)
	cmd := exec.Command(
		"curl",
		"-u",
		fmt.Sprintf("%s:%s", r.config.User, r.config.Password),
		"--upload-file",
		source,
		destination)
	data, err := cmd.CombinedOutput()
	if err != nil {
		model.Printf("%s", string(data))
	}
	return err
}

func (r Repository) Fetch(a model.Artifact, destination string) error {
	repo := "releases"
	if a.IsSnapshot() {
		repo = "snapshots"
	}
	source := r.config.URL + filepath.Join(r.config.Path, repo, a.StorageLocation(r.settings.OS, a.AnyOS))
	model.Printf("downloading %s to %s\n", source, destination)
	cmd := exec.Command(
		"curl",
		"-u",
		fmt.Sprintf("%s:%s", r.config.User, r.config.Password),
		source,
		"-o",
		destination)
	data, err := cmd.CombinedOutput()
	if err != nil {
		model.Printf("%s", string(data))
	}
	return err
}

func (r Repository) Exists(a model.Artifact) bool {
	repo := "releases"
	if a.IsSnapshot() {
		repo = "snapshots"
	}
	source := r.config.URL + filepath.Join(r.config.Path, repo, a.StorageLocation(r.settings.OS, a.AnyOS))
	model.Printf("%s", source)
	// TODO
	return false
}
