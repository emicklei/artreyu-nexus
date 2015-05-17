clean:
	rm -rf target
	
docker-build: clean
	mkdir target
	go build -ldflags "-X main.VERSION '$(VERSION)' -X main.BUILDDATE `date -u +%Y:%m:%d.%H:%M:%S`" -o /target/artreyu-nexus *.go	
	
build: clean
	mkdir target
	docker build -t --no-cache artreyu-builder .
	docker run --rm -e VERSION=$GIT_COMMIT -v `pwd`/target:/target -t artreyu-builder