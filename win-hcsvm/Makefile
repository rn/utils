GO_COMPILE=linuxkit/go-compile:bb0c6ae2f12a1b55df24ebce2067784a931285df

.PHONY: hcsvm.exe
hcsvm.exe: hcsvm.go
	docker run -it --rm \
		-v $(CURDIR):/go/src/github.com/rn/utils/hcsvm \
		-w /go/src/github.com/rn/utils/hcsvm \
		-e GOOS=windows \
		--entrypoint go $(GO_COMPILE) build hcsvm.go

.PHONY: vendor
vendor:
	docker run -it --rm \
		-v $(CURDIR):/go/src/github.com/rn/utils/hcsvm \
		-w /go/src/github.com/rn/utils/hcsvm \
		--entrypoint /go/bin/vndr $(GO_COMPILE)

.PHONY: clean
clean:
	rm hcsvm.exe
