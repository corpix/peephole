FROM fedora:latest

RUN mkdir                  /etc/go-boilerplate
ADD ./build/go-boilerplate /usr/bin/go-boilerplate

CMD [                                      \
    "/usr/bin/go-boilerplate",        \
    "--config",                            \
    "/etc/go-boilerplate/config.json" \
]
