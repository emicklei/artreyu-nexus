package main

import "github.com/emicklei/artreyu/model"

type Archive struct {
	artifact   model.Artifact
	repository model.Repository
	source     string
}

func (a Archive) Perform() {
	if len(a.source) == 0 {
		model.Fatalf("missing source")
	}
	err := a.repository.Store(a.artifact, a.source)
	if err != nil {
		model.Fatalf("unable to upload artifact:%v", err)
	}
}
