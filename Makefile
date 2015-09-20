local:
	go build -ldflags "-X main.VERSION '$(VERSION)' -X main.BUILDDATE `date -u +%Y:%m:%d.%H:%M:%S`" -o $(GOPATH)/bin/artreyu-nexus
	
build:
	mkdir -p /target
	go build -ldflags "-X main.VERSION '$(VERSION)' -X main.BUILDDATE `date -u +%Y:%m:%d.%H:%M:%S`" -o /target/artreyu-nexus *.go	
	
dockerbuild:
	docker build --no-cache=true --tag=artreyu-nexus-builder .
	docker run --rm -e VERSION=$(GIT_COMMIT) -v $(TARGET):/target -t artreyu-nexus-builder