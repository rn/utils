FROM alpine:3.10

ENV ACPICA_COMMIT=450ffd8b9c100db561ecf23063620cb107d68c30

RUN apk update && \
    apk add \
        bison \
        build-base \
        git \
        flex \
        m4

RUN git clone https://github.com/acpica/acpica.git && \
    cd /acpica && \
    git checkout ${ACPICA_COMMIT}

WORKDIR /acpica

RUN mkdir -p /res && \
    make OPT_LDFLAGS=-static && \
    cp generate/unix/bin/* /res

ENTRYPOINT cp -r /res/* /out

