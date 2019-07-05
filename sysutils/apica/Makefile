.phony: build
build: Dockerfile
	mkdir -p ./bin
	docker build -t acpica-build .
	docker run --rm -v $(CURDIR)/bin:/out acpica-build
