FROM alpine:3.10

ENV PCIUTILS_COMMIT=caca31a0eea41c7b051705704c1158fddc02fbd2

RUN apk update && \
    apk add \
        build-base \
        git \
        linux-headers


RUN git clone https://github.com/pciutils/pciutils.git && \
    cd /pciutils && \
    git checkout ${PCIUTILS_COMMIT}

WORKDIR /pciutils

RUN mkdir -p /res && \
    make LDFLAGS=-static SHARED=no && \
    cp lspci setpci /res

ENTRYPOINT cp -r /res/* /out

