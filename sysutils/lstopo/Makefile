.phony: build
build: Dockerfile
	mkdir -p ./bin
	docker build -t lstopo-build .
	docker run --rm -v $(CURDIR)/bin:/out lstopo-build
