FROM alpine:3.10

ENV HWLOC_COMMIT=46cd60cb6c59a464ea462d4f68937dcb6c38c00a

RUN apk update && \
    apk add \
        autoconf \
        automake \
        build-base \
        git \
        libtool \
        linux-headers


RUN git clone https://github.com/open-mpi/hwloc.git && \
    cd /hwloc && \
    git checkout ${HWLOC_COMMIT}

WORKDIR /hwloc

RUN mkdir -p /res && \
    ./autogen.sh && \
    ./configure --enable-static LDFLAGS=-static --enable-shared=no --disable-shared && \
    make LDFLAGS=-static && \
    cp utils/lstopo/lstopo-no-graphics /res/lstopo

ENTRYPOINT cp -r /res/* /out

