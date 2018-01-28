FROM rust:1.22.1-stretch

ENV CROSVM_REPO=https://chromium.googlesource.com/chromiumos/platform/crosvm
ENV CROSVM_COMMIT=fc44d8059b2333b7f49c1a11012eb6ca7d2250a5
ENV MINIJAIL_REPO=https://android.googlesource.com/platform/external/minijail
ENV MINIJAIL_COMMIT=d45fc420bb8fd9d1fc9297174f3c344db8c20bbd

# Install deps
RUN apt-get update && apt-get install -y libcap-dev

# Get source code
RUN git clone ${MINIJAIL_REPO} && \
    cd /minijail && \
    git checkout ${MINIJAIL_COMMIT} && \
    cd / && \
    git clone ${CROSVM_REPO} && \
    cd crosvm && \
    git checkout ${CROSVM_COMMIT}

# Compile and install minijail
WORKDIR /minijail
RUN make && \
    cp libminijail.so /usr/lib/ && \
    cp libminijail.h /usr/include/

# Compile crosvm
WORKDIR /crosvm
RUN cargo build --release
    
RUN mkdir /out && \
    cp /minijail/libminijail.so /out && \
    cp /crosvm/target/release/crosvm /out/crosvm

WORKDIR /out
ENTRYPOINT ["tar", "cf", "-", "libminijail.so", "crosvm" ]
