# This Makefile cross compiles using docker on macOS/Linux

GO_COMPILE=linuxkit/go-compile:bb0c6ae2f12a1b55df24ebce2067784a931285df

.PHONY: npterm.exe
npterm.exe: npterm.go
	docker run -it --rm \
		-v $(CURDIR):/go/src/github.com/rn/utils/npterm \
		-w /go/src/github.com/rn/utils/npterm \
		-e GOOS=windows \
		--entrypoint go $(GO_COMPILE) build npterm.go

.PHONY: vendor
vendor:
	docker run -it --rm \
		-v $(CURDIR):/go/src/github.com/rn/utils/npterm \
		-w /go/src/github.com/rn/utils/npterm \
		--entrypoint /go/bin/vndr $(GO_COMPILE)

.PHONY: clean
clean:
	rm npterm.exe
