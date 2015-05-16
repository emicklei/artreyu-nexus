clean:
	rm -rf target
	
build: clean
	mkdir -p target
	go build -ldflags "-X main.VERSION '$(VERSION)' -X main.BUILDDATE `date -u +%Y:%m:%d.%H:%M:%S`" -o /target/artreyu-nexus *.go