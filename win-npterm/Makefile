# This Makefile cross compiles using docker on macOS/Linux

GO_COMPILE=linuxkit/go-compile:fb53f01a669de5e91ec855b4f67a57b514b4f6ed

.PHONY: npterm.exe
npterm.exe: npterm.go
	docker run -it --rm \
		-v $(CURDIR):/go/src/github.com/rn/play/npterm \
		-w /go/src/github.com/rn/play/npterm \
		-e GOOS=windows \
		--entrypoint go $(GO_COMPILE) build npterm.go

.PHONY: vendor
vendor:
	docker run -it --rm \
		-v $(CURDIR):/go/src/github.com/rn/play/npterm \
		-w /go/src/github.com/rn/play/npterm \
		--entrypoint /go/bin/vndr $(GO_COMPILE)

.PHONY: clean
clean:
	rm npterm.exe
