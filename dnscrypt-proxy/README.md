# dnscrypt-proxy

This directory contains a minimal LinuxKit image to build a
[`dnscrypt-proxy`](https://github.com/jedisct1/dnscrypt-proxy) based
DNS proxy. It does three things:

- Use DNS over HTTP (or DNSSEC) for DNS queries, preventing, for
  example, your ISP to rewrite/log/block certain DNS queries.
- Use non-logging upstream DNS servers
- Block DNS queries for known trackers and other nasty sites

It works both on x86_64 and arm64, in particular a RPi3.


## Building

Just type:
```sh
make
```

This:
- build a local LinuxKit package called `dnscrypt/dnscrypt-proxy:dev`
  - Compiles dnscrypt-proxy
  - Creates the latest balcklist of domains to block
  - build a minimal LinuxKit package
- Build a LinuxKit image

On a x86_64 machine you will end up with a `kernel+initrd` and on a
arm64 machine you'll end up with a tarball with contents you can copy
over to a FAT32 partition on a RPi3 SD card. See the RPi3
documentation in the LinuxKit repo.


## Testing

The simplest way to test the image is with `hyperkit` and Docker for
Mac (at least that's how I test it):

- Build the LinuxKit image on your Mac (`make`)
- Run it `linuxkit run dnscrypt-proxy`
  - this will boot a VM and its network connection will be on the same
    network as the Docker for Mac VM.
- Run a alpine container: `docker run --rm -it alpine`, then
  - Install bind-tools: `apk add --no-cache bind-tools`
  - Get the IP address of the `dnscrypt-proxy` VM, say `192.168.65.9`
  - Perform DNS queries: `dig @192.168.65.9 www.google.com`
  - In the `dnscrypt-proxy` VM, check out the logs under:
    - `/var/log/dnscrypt-proxy.err.log`
    - `/var/log/dnscrypt/*.log`
