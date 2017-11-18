#! /bin/sh
# Create a bridge and add eth0 as master

ip link add br0 type bridge
ip link set dev eth0 master br0

ip link set dev eth0 up
ip link set dev br0 up
