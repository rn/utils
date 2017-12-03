The Chrome OS Virtual Machine Monitor
[`crosvm`](https://chromium.googlesource.com/chromiumos/platform/crosvm/)
is a lightweight VMM written in Rust. It runs on top of KVM.

The `Makefile` and `Dockerfile` compile `crosvm` and a suitable
version of `libminijail`. To build:

```
make
```

You should end up with a `crosvm` and `libminijail.so` binary. Copy
`libminijail.so` to `/usr/lib` or wherever `ldd` picks it up. You may
also need `libcap` (on Ubuntu or Debian `apt-get install -y
libcap-dev`).
