local:
	go build -ldflags "-X main.VERSION '$(VERSION)' -X main.BUILDDATE `date -u +%Y:%m:%d.%H:%M:%S`" -o $(GOPATH)/bin/artreyu-nexus
	
build:
	mkdir -p /target
	go build -ldflags "-X main.VERSION '$(VERSION)' -X main.BUILDDATE `date -u +%Y:%m:%d.%H:%M:%S`" -o /target/artreyu-nexus *.go	
	
# this task exists for Jenkins	
dockerbuild:
	docker build --no-cache=true --tag=artreyu-nexus-builder .
	docker run --rm -e VERSION=$(GIT_COMMIT) -v $(TARGET):/target -t artreyu-nexus-builder
	
# this task exists for local docker
docker:
	docker build --no-cache=true --tag=artreyu-nexus-builder .
	docker run --rm -v target:/target -t artreyu-nexus-builder	