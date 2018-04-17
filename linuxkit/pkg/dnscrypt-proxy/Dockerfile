FROM linuxkit/alpine:f3cd219615428b2bd943411723eb28875275fae7 as build

RUN apk add --no-cache go git musl-dev python3

ENV GOPATH=/go PATH=$PATH:/go/bin
ENV GITREPO=github.com/jedisct1/dnscrypt-proxy
ENV COMMIT=2.0.10

RUN go get -d ${GITREPO}/dnscrypt-proxy && \
    cd /go/src/${GITREPO} && \
    git checkout ${COMMIT} && \
    go build -o /dnscrypt-proxy ${GITREPO}/dnscrypt-proxy

# This is for the root fs. Not really needed, but handy for debug
FROM linuxkit/alpine:f3cd219615428b2bd943411723eb28875275fae7 AS rootfs
RUN mkdir -p /out/etc/apk && cp -r /etc/apk/* /out/etc/apk/
RUN apk add --no-cache --initdb -p /out \
    alpine-baselayout \
    busybox

FROM scratch
ENTRYPOINT []
CMD []
WORKDIR /
COPY --from=rootfs /out /
COPY --from=build /dnscrypt-proxy /usr/bin/dnscrypt-proxy
COPY /resolv.conf /etc/resolv.conf
COPY /dnscrypt-proxy.toml /etc/dnscrypt-proxy.toml
COPY /dnscrypt-blacklist.txt /etc/dnscrypt-blacklist.txt

ENTRYPOINT [ "/usr/bin/dnscrypt-proxy", "-config", "/etc/dnscrypt-proxy.toml" ]
