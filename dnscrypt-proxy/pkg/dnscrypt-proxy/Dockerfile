FROM linuxkit/alpine:f3cd219615428b2bd943411723eb28875275fae7 as build

RUN apk add --no-cache go git musl-dev

ENV GOPATH=/go PATH=$PATH:/go/bin
ENV GITREPO=github.com/jedisct1/dnscrypt-proxy
ENV COMMIT=2.0.10

RUN go get -d ${GITREPO}/dnscrypt-proxy && \
    cd /go/src/${GITREPO} && \
    git checkout ${COMMIT} && \
    go build -o /dnscrypt-proxy ${GITREPO}/dnscrypt-proxy

# Generating the blacklist requites python2 which we don't have in linuxkit/alpine
FROM alpine:3.7 as blacklist
ENV GITREPO=github.com/jedisct1/dnscrypt-proxy
COPY --from=build /go/src/${GITREPO}/ /dnscrypt-proxy/
RUN apk update && apk add python
RUN cd /dnscrypt-proxy/utils/generate-domains-blacklists && \
    python generate-domains-blacklist.py > /dnscrypt-blacklist.txt

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
COPY --from=blacklist /dnscrypt-blacklist.txt /etc/dnscrypt-blacklist.txt
COPY /resolv.conf /etc/resolv.conf
COPY /dnscrypt-proxy.toml /etc/dnscrypt-proxy.toml
COPY /dnscrypt-whitelist.txt /etc/dnscrypt-whitelist.txt

ENTRYPOINT [ "/usr/bin/dnscrypt-proxy", "-config", "/etc/dnscrypt-proxy.toml" ]
