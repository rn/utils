FROM linuxkit/alpine:d307c8a386fa3f32cddda9409b9687e191cdd6f1 AS mirror
RUN mkdir -p /out/etc/apk && cp -r /etc/apk/* /out/etc/apk/
RUN apk add --no-cache --initdb -p /out \
    argp-standalone \
    bash \
    bc \
    binutils-dev \
    bison \
    build-base \
    busybox \
    curl \
    installkernel \
    kmod \
    libelf-dev \
    linux-headers \
    musl-dev \
    musl-utils \
    perl \
    tar \
    xz \
    xz-dev \
    zlib-dev
RUN rm -rf /out/etc/apk /out/lib/apk /out/var/cache

FROM scratch
ENTRYPOINT []
CMD []
WORKDIR /
COPY --from=mirror /out/ /
COPY compile.sh /
CMD ["/compile.sh"]
