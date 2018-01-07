#!/bin/sh

if [ "$#" -ne 1 ]; then
    ITER=3
else
    ITER=$1
fi

VER=4.14.1
curl -fsSLO https://www.kernel.org/pub/linux/kernel/v4.x/linux-${VER}.tar.xz
tar xf linux-${VER}.tar.xz
cd linux-${VER}
CPUS=$(getconf _NPROCESSORS_ONLN)

compile() {
    make mrproper > /dev/null
    make defconfig > /dev/null
    make oldconfig > /dev/null
    BEFORE=$(date)
    time sh -c "make -j $CPUS > /dev/null"
    echo $BEFORE
    date
}

for i in $(seq $ITER); do
    echo "=== Run $i"
    compile
done
