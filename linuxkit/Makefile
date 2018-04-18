ARCH := $(shell uname -m)

ifeq ($(ARCH),x86_64)
DNSCRYPT_FORMAT = kernel+initrd
endif
ifeq ($(ARCH),aarch64)
DNSCRYPT_FORMAT = rpi3
endif

.PHONY: dnscrypt-proxy
dnscrypt-proxy:
	linuxkit pkg build -dev -org dnscrypt pkg/dnscrypt-proxy
	linuxkit build -format $(DNSCRYPT_FORMAT) -name dnscrypt-proxy dnscrypt-proxy-$(ARCH).yml
