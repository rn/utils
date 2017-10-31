GO_COMPILE=linuxkit/go-compile:fb53f01a669de5e91ec855b4f67a57b514b4f6ed

.PHONY: hcsvm.exe
hcsvm.exe: hcsvm.go
	docker run -it --rm \
		-v $(CURDIR):/go/src/github.com/rn/play/hcsvm \
		-w /go/src/github.com/rn/play/hcsvm \
		-e GOOS=windows \
		--entrypoint go $(GO_COMPILE) build hcsvm.go

.PHONY: vendor
vendor:
	docker run -it --rm \
		-v $(CURDIR):/go/src/github.com/rn/play/hcsvm \
		-w /go/src/github.com/rn/play/hcsvm \
		--entrypoint /go/bin/vndr $(GO_COMPILE)

.PHONY: clean
clean:
	rm hcsvm.exe
