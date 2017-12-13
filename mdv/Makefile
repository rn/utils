GO_COMPILE=linuxkit/go-compile:bb0c6ae2f12a1b55df24ebce2067784a931285df

GOOS?=darwin

build: mdv.go
	docker run -it --rm \
		-v $(CURDIR):/go/src/github.com/rn/utils/mdv \
		-w /go/src/github.com/rn/utils/mdv \
		-e GOOS=$(GOOS) \
		--entrypoint go $(GO_COMPILE) build mdv.go

.PHONY: vendor
vendor:
	docker run -it --rm \
		-v $(CURDIR):/go/src/github.com/rn/utils/mdv \
		-w /go/src/github.com/rn/utils/mdv \
		--entrypoint /go/bin/vndr $(GO_COMPILE)

.PHONY: clean
clean:
	rm -f mdv mdv.exe
