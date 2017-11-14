FROM golang:1.9.2 as builder

WORKDIR /go/src/github.com/corpix/go-boilerplate
COPY    . .
RUN     make

FROM fedora:latest

RUN  mkdir          /etc/go-boilerplate
COPY                /go/src/github.com/corpix/go-boilerplate/config.json           /etc/go-boilerplate/config.json
COPY --from=builder /go/src/github.com/corpix/go-boilerplate/build/go-boilerplate  /usr/bin/go-boilerplate

CMD [                                 \
    "/usr/bin/go-boilerplate",        \
    "--config",                       \
    "/etc/go-boilerplate/config.json" \
]
