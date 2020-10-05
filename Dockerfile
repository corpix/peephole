FROM golang:1.15 as builder

WORKDIR /go/src/github.com/corpix/peephole
COPY    . .
RUN     make build

FROM fedora:latest

RUN  mkdir          /etc/peephole
COPY --from=builder /go/src/github.com/corpix/peephole/config.json     /etc/peephole/config.json
COPY --from=builder /go/src/github.com/corpix/peephole/build/peephole  /usr/bin/peephole

EXPOSE 1338/tcp

CMD [                           \
    "/usr/bin/peephole",        \
    "--config",                 \
    "/etc/peephole/config.json" \
]
