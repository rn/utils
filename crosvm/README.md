The Chrome OS Virtual Machine Monitor
[`crosvm`](https://chromium.googlesource.com/chromiumos/platform/crosvm/)
is a lightweight VMM written in Rust. It runs on top of KVM.

## Build/Install

The `Makefile` and `Dockerfile` compile `crosvm` and a suitable
version of `libminijail`. To build:

```
make
```

You should end up with a `crosvm` and `libminijail.so` binaries as
well as the seccomp profiles in `./build`. Copy `libminijail.so` to
`/usr/lib` or wherever `ldd` picks it up. You may also need `libcap`
(on Ubuntu or Debian `apt-get install -y libcap-dev`).

You may also have to create an empty directory `/var/empty`

## Use with LinuxKit images

You can build a LinuxKit image suitable for `crosvm` with the
`kernel+squashfs` build format. For example, using this LinuxKit
YAML file (`minimal.yml`):

```
kernel:
  image: linuxkit/kernel:4.9.91
  cmdline: "console=tty0 console=ttyS0 console=ttyAMA0"
init:
  - linuxkit/init:v0.3
  - linuxkit/runc:v0.3
  - linuxkit/containerd:v0.3
services:
  - name: getty
    image: linuxkit/getty:v0.3
    env:
      - INSECURE=true
trust:
  org:
    - linuxkit
```

run:

```
linuxkit build -output kernel+squashfs minimal.yml
```

The kernel this produces (`minimal-kernel`) needs to be converted as
`crosvm` does not grok `bzImage`s. You can convert the LinuxKit kernel
image with
[extract-vmlinux](https://raw.githubusercontent.com/torvalds/linux/master/scripts/extract-vmlinux):

```
extract-vmlinux minimal-kernel > minimal-vmlinux
```

Then you can run `crosvm`:
```sh
./crosvm run --seccomp-policy-dir=./seccomp/x86_64 \
    --root ./minimal-squashfs.img \
    --mem 2048 \
    --multiprocess \
    --socket ./linuxkit-socket \
    minimal-vmlinux
```


Some known issues:
- With 4.14.x, a `BUG_ON()` is hit in `drivers/base/driver.c`. 4.9.x
  kernels seem to work.
- Networking does not yet work, so don't include a `onboot` `dhcpd` service.

