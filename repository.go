package main

import (
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/emicklei/artreyu/model"
	"github.com/emicklei/artreyu/transport"
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

	destinationURL, err := url.Parse(destination)
	if err != nil {
		return fmt.Errorf("invalid http request:%v", err)
	}
	destinationURL.User = url.UserPassword(r.config.User, r.config.Password)
	return transport.HttpPostFile(source, destinationURL.String())
}

func (r Repository) Fetch(a model.Artifact, destination string) error {
	repo := "releases"
	if a.IsSnapshot() {
		repo = "snapshots"
	}
	source := r.config.URL + filepath.Join(r.config.Path, repo, a.StorageLocation(r.settings.OS, a.AnyOS))
	model.Printf("downloading %s to %s\n", source, destination)
	sourceURL, err := url.Parse(source)
	if err != nil {
		return fmt.Errorf("invalid http request:%v", err)
	}
	sourceURL.User = url.UserPassword(r.config.User, r.config.Password)
	return transport.HttpGetFile(http.DefaultClient, sourceURL.String(), destination)
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
