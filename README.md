This is a repository where I dump miscellaneous utilities and other
code. Some of it is work in progress.

- [`crossvm`](./crossvm): Dockerfile/Makefile to compile
  [`crosvm`](https://chromium.googlesource.com/chromiumos/platform/crosvm/).
- [`photo`](./photo): A collection of scripts to manage meta-data and
  other things for scanned films/photos.
- [`rpncalc`](./rpncalc): A text based calculator (written in Python),
  which uses Reverse Polish Notation (like HP pocket
  calculators). It's also handy for low level work as it prints
  numbers in decimal, hex, and binary.
- [`hcsvm`](./win-hcsvm): A WIP go based utility to start Linux VMs on
  Hyper-V using the Host Compute Service (HCS). HCS is used to spin up
  Utility/Service VMs for containers on Windows 10 Pro and Windows
  Server 2016.
- [`npterm`](./win-npterm): A simple terminal utility (written in Go)
  which connects to a Named Pipe. It optionally logs all output to a
  file. I use this for serial console access to Hyper-V VMs.
  
