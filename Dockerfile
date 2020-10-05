FROM golang:1.15 as builder

WORKDIR /go/src/github.com/corpix/peephole
COPY    . .
ENV     CGO_ENABLED 0
RUN     make build

FROM alpine:latest

RUN  mkdir          /etc/peephole
COPY --from=builder /go/src/github.com/corpix/peephole/config.yaml /etc/peephole/config.yaml
COPY --from=builder /go/src/github.com/corpix/peephole/main        /usr/bin/peephole

EXPOSE 1338/tcp

CMD [                           \
    "/usr/bin/peephole",        \
    "--config",                 \
    "/etc/peephole/config.yaml" \
]
