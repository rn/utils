.phony: build
build: Dockerfile
	mkdir -p ./bin
	docker build -t lspci-build .
	docker run --rm -v $(CURDIR)/bin:/out lspci-build
